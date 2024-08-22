package main

type InstrFunc func(state *State, instruction *Instruction)

func funcLb(state *State, instruction *Instruction) {

}

type FuncName string

const (
	FuncLb FuncName = "FuncLb"

	InsnLb  FuncName = "InsnLb"
	InsnLh  FuncName = "InsnLh"
	InsnLw  FuncName = "InsnLw"
	InsnLd  FuncName = "InsnLd"
	InsnLbu FuncName = "InsnLbu"
	InsnLhu FuncName = "InsnLhu"
	InsnLwu FuncName = "InsnLwu"

	InsnFlw FuncName = "InsnFlw"
	InsnFld FuncName = "InsnFld"
)

var InstrFuncs = make(map[FuncName]InstrFunc)

func init() {
	InstrFuncs[FuncLb] = funcLb
}

func ExecBlockInterp(state *State) {
	insn := Instruction{}
	for {
		pcAddr := ToHost(state.pc)
		data := GetProcessMemory(pcAddr, uint64(Dword))
		// uint64转为uint32会截断高位的4字节数据，所以在这里没有影响
		InstructionDecode(&insn, uint32(data))

		// 执行相应的指令方法
		InstrFuncs[insn.iType](state, &insn)
		// zero寄存器置0
		state.gpRegs[zero] = 0

		if insn.cont {
			break
		}

		// 判断是否是压缩指令，压缩指令占两个字节
		if insn.rvc {
			state.pc = state.pc + 2
		} else {
			state.pc = state.pc + 4
		}
	}
}
