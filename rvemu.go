package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Go-Rv_Emu")

	if len(os.Args) != 2 {
		DPrintf("Usage: rvemu a.out\n")
		os.Exit(1)
	}

	machine := &Machine{
		mmu: &Mmu{},
	}

	machine.MachineLoadProgram(os.Args[1])
}
