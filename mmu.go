package main

import (
	"encoding/binary"
	"errors"
	"os"
)

type Mmu struct {
	EEntry uint64
}

func (mmu *Mmu) LoadPhdr(phdr *Elf64PhdrT, ehdr *Elf64EhdrT, i uint16, file *os.File) error {
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
		DPrintf("MMU load elf ehdr fail, bad elf file")
		return errors.New("bad elf file")
	}

	if ehdr.EMachine != EmRiscV || ehdr.EIdent[EiClass] != ElfClass64 {
		DPrintf("MMU load elf ehdr fail, only risc-v 64 bit is support")
		return errors.New("bad elf file")
	}

	mmu.EEntry = ehdr.EEntry

	// 3. 读取每一个section
	for i := uint16(0); i < ehdr.EPhnum; i++ {
		phdr := Elf64PhdrT{}
		err := mmu.LoadPhdr(&phdr, &ehdr, i, file)
		if err != nil {
			DPrintf("MMU load elf phdr fail, err: %v", err.Error())
			return err
		}

		if phdr.PType == Elf64PhdrPTypeLoad {
			DPrintf("this is a load section, phdr: %v", phdr)
		}
	}

	return nil
}
