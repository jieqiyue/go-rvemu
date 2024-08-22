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
		Mmu:   &Mmu{},
		State: &State{},
	}

	machine.MachineLoadProgram(os.Args[1])

	for {
		exitReson := machine.MachineStep()
		assert(exitReson == ECall, "")
	}

	return
}
