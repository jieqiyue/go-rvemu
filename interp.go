package main

import "unsafe"

type InstrFunc func(state *State, instruction *Instruction)

func emptyFunc(state *State, instruction *Instruction) {
}

func funcAuipc(state *State, instruction *Instruction) {
	val := int64(state.pc) + int64(instruction.imm)
	state.gpRegs[instruction.rd] = uint64(val)
}

func funcLb(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*int8)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLh(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*int16)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLw(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*int32)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLd(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*int64)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLbu(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*uint8)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLhu(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*uint16)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcLwu(state *State, instruction *Instruction) {
	addr := int64(state.gpRegs[instruction.rs1]) + int64(instruction.imm)
	addr = int64(ToHost(uint64(addr)))

	p := (*uint32)(unsafe.Pointer(uintptr(addr)))
	state.gpRegs[instruction.rd] = uint64(*p)
}

func funcAddi(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(rs1) + imm)
}

func funcSlli(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = rs1 << (imm & 0x3f)
}

func funcSlti(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	if int64(rs1) < imm {
		state.gpRegs[instruction.rd] = 1
	} else {
		state.gpRegs[instruction.rd] = 0
	}
}

func funcSltiu(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	if uint64(rs1) < uint64(imm) {
		state.gpRegs[instruction.rd] = 1
	} else {
		state.gpRegs[instruction.rd] = 0
	}
}

func funcXori(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(rs1) ^ imm)
}

func funcSrli(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = rs1 >> (imm & 0x3f)
}

func funcSrai(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = int64(rs1) >> (imm & 0x3f)
}

func funcOri(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = rs1 | uint64(imm)
}

func funcAndi(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = rs1 & uint64(imm)
}

func funcAddiw(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(int32(int64(rs1) + imm)))
}

func funcSlliw(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(int32(rs1 << (imm & 0x1f))))
}

func funcSrliw(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(int32(uint32(rs1 >> (imm & 0x1f)))))
}

func funcSraiw(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	imm := int64(instruction.imm)

	state.gpRegs[instruction.rd] = uint64(int64(int32(rs1 >> (imm & 0x1f))))
}

func funcSb(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	p := (*uint8)(unsafe.Pointer(uintptr(ToHost(uint64(int64(rs1) + int64(instruction.imm))))))
	*p = uint8(rs2)
}

func funcSh(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	p := (*uint16)(unsafe.Pointer(uintptr(ToHost(uint64(int64(rs1) + int64(instruction.imm))))))
	*p = uint16(rs2)
}

func funcSw(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	p := (*uint32)(unsafe.Pointer(uintptr(ToHost(uint64(int64(rs1) + int64(instruction.imm))))))
	*p = uint32(rs2)
}

func funcSd(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	p := (*uint64)(unsafe.Pointer(uintptr(ToHost(uint64(int64(rs1) + int64(instruction.imm))))))
	*p = uint64(rs2)
}
func funcAdd(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 + rs2
}

func funcSll(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 << (rs2 & 0x3f)
}

func funcSlt(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = (int64(rs1) << int64(rs2))
}

func funcSltu(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = (uint64(rs1) << uint64(rs2))
}

func funcXor(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = (rs1 ^ rs2)
}

func funcSrl(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 >> (rs2 & 0x3f)
}

func funcOr(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 | rs2
}

func funcAnd(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 & rs2
}

func funcMul(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = rs1 * rs2
}

func funcMulh(state *State, instruction *Instruction) {
	rs1 := state.gpRegs[instruction.rs1]
	rs2 := state.gpRegs[instruction.rs2]
	state.gpRegs[instruction.rd] = (uint64(rs1) << uint64(rs2))
}

type FuncName string

const (
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

	InsnFmaddS FuncName = "InsnFmaddS"
	InsnFmaddD FuncName = "InsnFmaddD"

	InsnFmsubS FuncName = "InsnFmsubS"
	InsnFmsubD FuncName = "InsnFmsubD"

	InsnFnmsubS FuncName = "InsnFnmsubS"
	InsnFnmsubD FuncName = "InsnFnmsubD"

	InsnFnmaddS FuncName = "InsnFnmaddS"
	InsnFnmaddD FuncName = "InsnFnmaddD"

	InsnFaddS FuncName = "InsnFaddS"
	InsnFaddD FuncName = "InsnFaddD"
	InsnFsubS FuncName = "InsnFsubS"
	InsnFsubD FuncName = "InsnFsubD"
	InsnFmulS FuncName = "InsnFmulS"
	InsnFmulD FuncName = "InsnFmulD"
	InsnFdivS FuncName = "InsnFdivS"
	InsnFdivD FuncName = "InsnFdivD"

	InsnFsgnjS  FuncName = "InsnFsgnjS"
	InsnFsgnjnS FuncName = "InsnFsgnjnS"
	InsnFsgnjxS FuncName = "InsnFsgnjxS"

	InsnFsgnjD  FuncName = "InsnFsgnjD"
	InsnFsgnjnD FuncName = "InsnFsgnjnD"
	InsnFsgnjxD FuncName = "InsnFsgnjxD"

	InsnFminS FuncName = "InsnFminS"
	InsnFmaxS FuncName = "InsnFmaxS"

	InsnFminD FuncName = "InsnFminD"
	InsnFmaxD FuncName = "InsnFmaxD"

	InsnFcvtSD FuncName = "InsnFcvtSD"
	InsnFcvtDS FuncName = "InsnFcvtDS"
	InsnFsqrtS FuncName = "InsnFsqrtS"
	InsnFsqrtD FuncName = "InsnFsqrtD"

	InsnFleS FuncName = "InsnFleS"
	InsnFltS FuncName = "InsnFltS"
	InsnFeqS FuncName = "InsnFeqS"

	InsnFleD FuncName = "InsnFleD"
	InsnFltD FuncName = "InsnFltD"
	InsnFeqD FuncName = "InsnFeqD"

	InsnFcvtWS  FuncName = "InsnFcvtWS"
	InsnFcvtWuS FuncName = "InsnFcvtWuS"
	InsnFcvtLS  FuncName = "InsnFcvtLS"
	InsnFcvtLuS FuncName = "InsnFcvtLuS"

	InsnFcvtWD  FuncName = "InsnFcvtWD"
	InsnFcvtWud FuncName = "InsnFcvtWud"
	InsnFcvtLD  FuncName = "InsnFcvtLD"
	InsnFcvtLuD FuncName = "InsnFcvtLuD"

	InsnFcvtSW  FuncName = "InsnFcvtSW"
	InsnFcvtSWu FuncName = "InsnFcvtSWu"
	InsnFcvtSL  FuncName = "InsnFcvtSL"
	InscFcvtSLu FuncName = "InscFcvtSLu"

	InsnFcvtDW  FuncName = "InsnFcvtDW"
	InsnFcvtDWu FuncName = "InsnFcvtDWu"
	InsnFcvtDL  FuncName = "InsnFcvtDL"
	InsnFcvtDLu FuncName = "InsnFcvtDLu"

	InsnFmvXW   FuncName = "InsnFmvXW"
	InsnFclassS FuncName = "InsnFclassS"

	InsnFmvXD FuncName = "InsnFmvXD"

	InsnFclassD FuncName = "InsnFclassD"

	InsnFmvWX FuncName = "InsnFmvWX"

	InsnFmvDX FuncName = "InsnFmvDX"

	InsnBeq  FuncName = "InsnBeq"
	InsnBne  FuncName = "InsnBne"
	InsnBlt  FuncName = "InsnBlt"
	InsnBge  FuncName = "InsnBge"
	InsnBltu FuncName = "InsnBltu"
	InsnBgeu FuncName = "InsnBgeu"

	InsnJalr FuncName = "InsnJalr"
	InsnJal  FuncName = "InsnJal"

	InsnEcall FuncName = "InsnEcall"

	InsnCsrrw  FuncName = "InsnCsrrw"
	InsnCsrrs  FuncName = "InsnCsrrs"
	InsnCsrrc  FuncName = "InsnCsrrc"
	InsnCsrrwi FuncName = "InsnCsrrwi"
	InsnCsrrsi FuncName = "InsnCsrrsi"
	InsnCsrrci FuncName = "InsnCsrrci"
)

var InstrFuncs = make(map[FuncName]InstrFunc)

func init() {
	InstrFuncs[InsnLb] = funcLb
	InstrFuncs[InsnLh] = funcLh
	InstrFuncs[InsnLw] = funcLw
	InstrFuncs[InsnLd] = funcLd
	InstrFuncs[InsnLbu] = funcLbu
	InstrFuncs[InsnLhu] = funcLhu
	InstrFuncs[InsnLwu] = funcLwu

	InstrFuncs[InsnFence] = emptyFunc
	InstrFuncs[InsnFenceI] = emptyFunc

	InstrFuncs[InsnAddi] = funcAddi
	InstrFuncs[InsnSlli] = funcSlli
	InstrFuncs[InsnSlti] = funcSlti
	InstrFuncs[InsnSltiu] = funcSltiu
	InstrFuncs[InsnXori] = funcXori
	InstrFuncs[InsnSrli] = funcSrli
	InstrFuncs[InsnSrai] = funcSrai
	InstrFuncs[InsnOri] = funcOri
	InstrFuncs[InsnAndi] = funcAndi
	InstrFuncs[InsnAuipc] = funcAuipc

	InstrFuncs[InsnAddiw] = funcAddiw
	InstrFuncs[InsnSlliw] = funcSlliw
	InstrFuncs[InsnSrliw] = funcSrliw
	InstrFuncs[InsnSraiw] = funcSraiw
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
