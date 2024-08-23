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

	InsnFence  FuncName = "InsnFence"
	InsnFenceI FuncName = "InsnFenceI"

	InsnAddi  FuncName = "InsnAddi"
	InsnSlli  FuncName = "InsnSlli"
	InsnSlti  FuncName = "InsnSlti"
	InsnSltiu FuncName = "InsnSltiu"
	InsnXori  FuncName = "InsnXori"
	InsnSrli  FuncName = "InsnSrli"
	InsnSrai  FuncName = "InsnSrai"
	InsnOri   FuncName = "InsnOri"
	InsnAndi  FuncName = "InsnAndi"

	InsnAuipc FuncName = "InsnAuipc"

	InsnAddiw FuncName = "InsnAddiw"
	InsnSlliw FuncName = "InsnSlliw"
	InsnSrliw FuncName = "InsnSrliw"
	InsnSraiw FuncName = "InsnSraiw"

	InsnSb FuncName = "InsnSb"
	InsnSh FuncName = "InsnSh"
	InsnSw FuncName = "InsnSw"
	InsnSd FuncName = "InsnSd"

	InsnFsw FuncName = "InsnFsw"
	InsnFsd FuncName = "InsnFsd"

	InsnAdd  FuncName = "InsnAdd"
	InsnSll  FuncName = "InsnSll"
	InsnSlt  FuncName = "InsnSlt"
	InsnSltu FuncName = "InsnSltu"
	InsnXor  FuncName = "InsnXor"
	InsnSrl  FuncName = "InsnSrl"
	InsnOr   FuncName = "InsnOr"
	InsnAnd  FuncName = "InsnAnd"

	InsnMul    FuncName = "InsnMul"
	InsnMulh   FuncName = "InsnMulh"
	InsnMulhsu FuncName = "InsnMulhsu"
	InsnMulhu  FuncName = "InsnMulhu"
	InsnDiv    FuncName = "InsnDiv"
	InsnDivu   FuncName = "InsnDivu"
	InsnRem    FuncName = "InsnRem"
	InsnRemu   FuncName = "InsnRemu"
	InsnSub    FuncName = "InsnSub"
	InsnSra    FuncName = "InsnSra"

	InsnLui FuncName = "InsnLui"

	InsnAddw FuncName = "InsnAddw"
	InsnSllw FuncName = "InsnSllw"
	InsnSrlw FuncName = "InsnSrlw"

	InsnMulw  FuncName = "InsnMulw"
	InsnDivw  FuncName = "InsnDivw"
	InsnDivuw FuncName = "InsnDivuw"
	InsnRemw  FuncName = "InsnRemw"
	InsnRemuw FuncName = "InsnRemuw"

	InsnSubw FuncName = "InsnSubw"
	InsnSraw FuncName = "InsnSraw"
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
