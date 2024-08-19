package main

import (
	"fmt"
	"os"
)

type Machine struct {
	mmu *Mmu
}

func (machine *Machine) MachineLoadProgram(path string) {
	// 打开 ELF 文件
	file, err := os.Open(path)
	if err != nil {
		DPrintf("Machine load program fail, err: %s\n", err.Error())
		return
	}

	defer file.Close()

	if err = machine.mmu.MmuLoadElf(file); err != nil {
		DPrintf("Machine fail load elf, err: %v\n", err.Error())
		return
	}

	// 打印 ELF 头部信息

	fmt.Printf("Entry: %d\n", machine.mmu.EEntry)
}
