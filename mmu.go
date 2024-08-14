package main

import (
	"encoding/binary"
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
		DPrintf("Machine load elf header fail, err:%s\n", err.Error())
		return err
	}

	mmu.EEntry = header.EEntry

	return nil
}
