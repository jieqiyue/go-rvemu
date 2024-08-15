package main

import (
	"encoding/binary"
	"errors"
	"os"
)

type Mmu struct {
	EEntry uint64
}

func (mmu *Mmu) MmuLoadElf(file *os.File) error {
	// 读取 ELF 头部
	var header Elf64EhdrT
	err := binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		DPrintf("MMU load elf header fail, err:%s\n", err.Error())
		return err
	}

	// 判断ELF文件魔数是否符合
	if string(header.EIdent[:len(ElfMagic)]) != ElfMagic {
		DPrintf("MMU load elf header fail, bad elf file")
		return errors.New("bad elf file")
	}

	if header.EMachine != EmRiscV || header.EIdent[EiClass] != ElfClass64 {
		DPrintf("MMU load elf header fail, only risc-v 64 bit is support")
		return errors.New("bad elf file")
	}

	mmu.EEntry = header.EEntry

	return nil
}
