package main

import (
	"encoding/binary"
	"errors"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// HostMemOffset 将这个内存地址作为被加载的起始地址
const HostMemOffset = 0x088800000000

type Mmu struct {
	EEntry    uint64
	HostAlloc uint64
	Alloc     uint64
	Base      uint64
}

func elfFlagsToMmapProt(flags uint32) int {
	prot := 0
	if flags&PFR != 0 {
		prot = prot | syscall.PROT_READ
	}

	if flags&PFW != 0 {
		prot = prot | syscall.PROT_WRITE
	}

	if flags&PFX != 0 {
		prot = prot | syscall.PROT_EXEC
	}

	return prot
}

func PageDown(address uint64, pageSize uint64) uint64 {
	return address & (-pageSize)
}

func PageUp(address uint64, pageSize uint64) uint64 {
	return ((address) + (pageSize) - 1) & -(pageSize)
}

// ToHost 将内存地址转化为Host程序的内存地址
func (mmu *Mmu) ToHost(origin uint64) uint64 {
	return origin + HostMemOffset
}

func (mmu *Mmu) ToGuest(origin uint64) uint64 {
	return origin - HostMemOffset
}

func (mmu *Mmu) MmuLoadPhdr(phdr *Elf64PhdrT, ehdr *Elf64EhdrT, i uint16, file *os.File) error {
	// 1. 先把文件指针定位到这个段的开头为止
	_, err := file.Seek(int64(ehdr.EPhoff+uint64(i*ehdr.EPhentsize)), 0)
	if err != nil {
		DPrintf("read elf section seek fail, err: %v\n", err.Error())
		return err
	}

	// 2. 读取相应的节到内存中，节都是固定大小的
	// todo 如果这个文件没有这么大的话，binary.Read会报错吗？
	err = binary.Read(file, binary.LittleEndian, phdr)
	if err != nil {
		DPrintf("read elf section fail, err: %v\n", err.Error())
		return err
	}

	return nil
}

func (mmu *Mmu) MmuLoadSegment(phdr *Elf64PhdrT, file *os.File) error {
	pageSize := syscall.Getpagesize()
	// 该段所在文件的偏移量（不是从头开始偏移的？）
	offset := phdr.POffset
	vaAddr := mmu.ToHost(phdr.PVaddr)
	alignedVaAddr := PageDown(vaAddr, uint64(pageSize))
	fileSize := phdr.PFilesz + (vaAddr - alignedVaAddr) // 3576
	memSize := phdr.PMemsz + (vaAddr - alignedVaAddr)
	prot := elfFlagsToMmapProt(phdr.PFlags)

	// 满足这个的条件是offset的pageDown和vaAddr的pageDown是相等的
	// 关于MmapPtr函数：https://blog.csdn.net/qq_43009242/article/details/141318064
	memAddr, err := unix.MmapPtr(int(file.Fd()), int64(PageDown(offset, uint64(pageSize))), unsafe.Pointer(uintptr(alignedVaAddr)), uintptr(fileSize), prot, syscall.MAP_PRIVATE|syscall.MAP_FIXED)
	if err != nil {
		DPrintf("MMU mmap segment to memory fail, err: %v\n", err.Error())
		return err
	}

	DPrintf("mmap over, offset: %v\n", offset)
	DPrintf("mmap over, vaAddr: %v\n", vaAddr)
	DPrintf("mmap over, alignedVaAddr: %v\n", alignedVaAddr)
	DPrintf("mmap over, fileSize: %v\n", fileSize)
	DPrintf("mmap over, memSize: %v\n", memSize)
	DPrintf("mmap over, prot: %v\n", prot)

	assert(uintptr(memAddr) == uintptr(alignedVaAddr), "MMU mmap a segment to memory fail, alloc memory address not equal to desire address")

	remainingBss := PageUp(memSize, uint64(pageSize)) - PageUp(fileSize, uint64(pageSize))
	// 如果memSize大于fileSize，则表示超过一个页了，则需要再次分配bss段内存
	if remainingBss > 0 {
		bssAddr, err := unix.MmapPtr(-1, 0, unsafe.Pointer(uintptr(alignedVaAddr+PageUp(fileSize, uint64(pageSize)))), uintptr(remainingBss), prot, syscall.MAP_PRIVATE|syscall.MAP_FIXED|syscall.MAP_ANONYMOUS)
		if err != nil {
			DPrintf("MMU mmap bss to mem fail, err: %s\n", err.Error())
			return err
		}

		assert(uintptr(bssAddr) == uintptr(alignedVaAddr+PageUp(fileSize, uint64(pageSize))), "MMU mmap bss fail")
	}

	mmu.HostAlloc = Max(mmu.HostAlloc, alignedVaAddr+PageUp(memSize, uint64(pageSize)))
	mmu.Alloc = mmu.ToGuest(mmu.HostAlloc)
	mmu.Base = mmu.Alloc

	return nil
}

func (mmu *Mmu) MmuLoadElf(file *os.File) error {
	// 1. 读取 ELF 头部
	var ehdr Elf64EhdrT
	err := binary.Read(file, binary.LittleEndian, &ehdr)
	if err != nil {
		DPrintf("MMU load elf ehdr fail, err:%s\n", err.Error())
		return err
	}

	// 2. 判断ELF文件魔数是否符合
	if string(ehdr.EIdent[:len(ElfMagic)]) != ElfMagic {
		DPrintf("MMU load elf ehdr fail, bad elf file\n")
		return errors.New("bad elf file")
	}

	if ehdr.EMachine != EmRiscV || ehdr.EIdent[EiClass] != ElfClass64 {
		DPrintf("MMU load elf ehdr fail, only risc-v 64 bit is support\n")
		return errors.New("bad elf file")
	}

	mmu.EEntry = ehdr.EEntry

	// 3. 读取每一个section
	for i := uint16(0); i < ehdr.EPhnum; i++ {
		phdr := Elf64PhdrT{}
		err := mmu.MmuLoadPhdr(&phdr, &ehdr, i, file)
		if err != nil {
			DPrintf("MMU load elf phdr fail, err: %v\n", err.Error())
			return err
		}

		if phdr.PType == Elf64PhdrPTypeLoad {
			DPrintf("this is a load section, phdr: %v", phdr)
			err := mmu.MmuLoadSegment(&phdr, file)
			if err != nil {
				DPrintf("MMU load elf segment fail, err: %v\n", err.Error())
				return err
			}
		}
	}

	return nil
}
