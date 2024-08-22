package main

import (
	"fmt"
	"os"
)

type Machine struct {
	Mmu   *Mmu
	State *State
}

func (machine *Machine) MachineStep() ExitReason {
	for {
		ExecBlockInterp(machine.State)

		if machine.State.exitReason == InDirectBranch || machine.State.exitReason == DirectBranch {
			continue
		}
		break
	}

	assert(machine.State.exitReason == ECall, "not a reasonable exit reason")

	return ECall
}

func (machine *Machine) MachineLoadProgram(path string) {
	// 打开 ELF 文件
	file, err := os.Open(path)
	if err != nil {
		DPrintf("Machine load program fail, err: %s\n", err.Error())
		return
	}

	defer file.Close()

	if err = machine.Mmu.MmuLoadElf(file); err != nil {
		DPrintf("Machine fail load elf, err: %v\n", err.Error())
		return
	}

	// 打印 ELF 头部信息
	fmt.Printf("Entry: %d\n", machine.Mmu.EEntry)
	fmt.Printf("MMU:%v\n", machine.Mmu)
	machine.State.pc = machine.Mmu.EEntry
}
