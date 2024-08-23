package main

// 在go语言中逻辑右移还是算术右移是通过远算数的类型来决定的
func GetQuadrant(data uint32) uint32 {
	return ((data) >> 0) & 0x3
}

/*
	非压缩指令
    从risc-v的指令中提取出各个操作数（操作寄存器，目的寄存器，立即数等）
*/

func OpCode(data uint32) uint32 {
	return (data >> 2) & 0x1f
}

func Rd(data uint32) uint32 {
	return (data >> 7) & 0x1f
}

func Rs1(data uint32) uint32 {
	return (data >> 15) & 0x1f
}

func Rs2(data uint32) uint32 {
	return (data >> 20) & 0x1f
}

func Rs3(data uint32) uint32 {
	return (data >> 27) & 0x1f
}

func Funct2(data uint32) uint32 {
	return (data >> 25) & 0x3
}

func Funct3(data uint32) uint32 {
	return (data >> 12) & 0x7
}

func Funct7(data uint32) uint32 {
	return (data >> 25) & 0x7f
}

func Imm116(data uint32) uint32 {
	return (data >> 26) & 0x3f
}

func InstrBTypeRead(instruction *Instruction, data uint32) {
	imm12 := (data >> 31) & 0x1
	imm105 := (data >> 25) & 0x3f
	imm41 := (data >> 8) & 0xf
	imm11 := (data >> 7) & 0x1

	imm := int32((imm12 << 12) | (imm11 << 11) | (imm105 << 5) | (imm41 << 1))
	imm = (imm << 19) >> 19

	instruction.imm = imm
	instruction.rs1 = int8(Rs1(data))
	instruction.rs2 = int8(Rs2(data))

	return
}

func InstrITypeRead(instruction *Instruction, data uint32) {
	instruction.imm = int32(data >> 20)
	instruction.rs1 = int8(Rs1(data))
	instruction.rd = int8(Rd(data))

	return
}

func InstrUTypeRead(instruction *Instruction, data uint32) {
	instruction.imm = int32(data & 0xfffff000)
	instruction.rd = int8(Rd(data))

	return
}

func InstrSTypeRead(instruction *Instruction, data uint32) {
	imm115 := (data >> 25) & 0x7f
	imm40 := (data >> 7) & 0x1f

	imm := int32((imm115 << 5) | imm40)
	instruction.imm = imm
	instruction.rs1 = int8(Rs1(data))
	instruction.rs2 = int8(Rs2(data))

	return
}

func InstrRTypeRead(instruction *Instruction, data uint32) {
	instruction.rs1 = int8(Rs1(data))
	instruction.rs2 = int8(Rs2(data))
	instruction.rd = int8(Rd(data))

	return
}

func InstrFprTypeRead(instruction *Instruction, data uint32) {
	instruction.rs1 = int8(Rs1(data))
	instruction.rs2 = int8(Rs2(data))
	instruction.rs3 = int8(Rs3(data))
	instruction.rd = int8(Rd(data))

	return
}

func InstrJTypeRead(instruction *Instruction, data uint32) {
	imm20 := (data >> 31) & 0x1
	imm101 := (data >> 21) & 0x3ff
	imm11 := (data >> 20) & 0x1
	imm1912 := (data >> 12) & 0xff

	imm := int32((imm20 << 20) | (imm1912 << 12) | (imm11 << 11) | (imm101 << 1))
	imm = (imm << 11) >> 11

	instruction.imm = imm
	instruction.rd = int8(Rd(data))
	return
}

func InstrCsrTypeRead(instruction *Instruction, data uint32) {
	instruction.csr = int16(data >> 20)
	instruction.rs1 = int8(Rs1(data))
	instruction.rd = int8(Rd(data))

	return
}

func InstructionDecode(instruction *Instruction, data uint32) {
	quadrant := GetQuadrant(data)
	switch quadrant {
	case 0x0:
		Fatal("unimplemented")
	case 0x1:
		Fatal("unimplemented")
	case 0x2:
		Fatal("unimplemented")
	case 0x3:
		opCode := OpCode(data)
		switch opCode {
		case 0x0:
			funct3 := Funct3(data)
			InstrITypeRead(instruction, data)

			switch funct3 {
			case 0x0: /* LB */
				instruction.iType = InsnLb
				return
			case 0x1: /* LH */
				instruction.iType = InsnLh
				return
			case 0x2:
				instruction.iType = InsnLw
				return
			case 0x3:
				instruction.iType = InsnLd
				return
			case 0x4:
				instruction.iType = InsnLbu
				return
			case 0x5:
				instruction.iType = InsnLhu
				return
			case 0x6:
				instruction.iType = InsnLwu
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x1:
			funct3 := Funct3(data)

			InstrITypeRead(instruction, data)
			switch funct3 {
			case 0x2:
				instruction.iType = InsnFlw
				return
			case 0x3:
				instruction.iType = InsnFld
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x3:
			funct3 := Funct3(data)

			switch funct3 {
			case 0x0: /* FENCE */
				instruction.iType = InsnFence
				return
			case 0x1:
				instruction.iType = InsnFenceI
			default:
				UnReachable()
			}
			UnReachable()
		case 0x4:
			funct3 := Funct3(data)

			InstrITypeRead(instruction, data)
			switch funct3 {
			case 0x0:
				instruction.iType = InsnAddi
				return
			case 0x1:
				imm116 := Imm116(data)
				if imm116 == 0 {
					instruction.iType = InsnSlli
				} else {
					UnReachable()
				}
				return
			case 0x2:
				instruction.iType = InsnSlti
				return
			case 0x3:
				instruction.iType = InsnSltiu
				return
			case 0x4:
				instruction.iType = InsnXori
				return
			case 0x5:
				imm116 := Imm116(data)
				if imm116 == 0x0 {
					instruction.iType = InsnSrli
				} else if imm116 == 0x10 {
					instruction.iType = InsnSrai
				} else {
					UnReachable()
				}
				return
			case 0x6:
				instruction.iType = InsnOri
				return
			case 0x7:
				instruction.iType = InsnAndi
				return
			default:
				Fatal("unknow funct3")
			}
			UnReachable()
		case 0x5:
			InstrUTypeRead(instruction, data)
			instruction.iType = InsnAuipc
			return
		case 0x6:
			funct3 := Funct3(data)
			funct7 := Funct7(data)

			InstrITypeRead(instruction, data)

			switch funct3 {
			case 0x0:
				instruction.iType = InsnAddiw
				return
			case 0x1:
				assert(funct7 == 0, "funct7 is not 0")
				instruction.iType = InsnSlliw
				return
			case 0x5:
				switch funct7 {
				case 0x0:
					instruction.iType = InsnSrliw
					return
				case 0x20:
					instruction.iType = InsnSraiw
					return
				default:
					UnReachable()
				}
			default:
				UnReachable()
			}

			UnReachable()
		case 0x8:
			funct3 := Funct3(data)

			InstrSTypeRead(instruction, data)

			switch funct3 {
			case 0x0:
				instruction.iType = InsnSb
				return
			case 0x1:
				instruction.iType = InsnSh
				return
			case 0x2:
				instruction.iType = InsnSw
				return
			case 0x3:
				instruction.iType = InsnSd
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x9:
			funct3 := Funct3(data)

			InstrSTypeRead(instruction, data)

			switch funct3 {
			case 0x2: /* FSW */
				instruction.iType = InsnFsw
				return
			case 0x3: /* FSD */
				instruction.iType = InsnFsd
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0xc:
			InstrRTypeRead(instruction, data)

			funct3 := Funct3(data)
			funct7 := Funct7(data)

			switch funct7 {
			case 0x0:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnAdd
					return
				case 0x1:
					instruction.iType = InsnSll
					return
				case 0x2:
					instruction.iType = InsnSlt
					return
				case 0x3:
					instruction.iType = InsnSltu
					return
				case 0x4:
					instruction.iType = InsnXor
					return
				case 0x5:
					instruction.iType = InsnSrl
					return
				case 0x6:
					instruction.iType = InsnOr
					return
				case 0x7:
					instruction.iType = InsnAnd
					return
				default:
					UnReachable()
				}
			case 0x1:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnMul
					return
				case 0x1:
					instruction.iType = InsnMulh
					return
				case 0x2:
					instruction.iType = InsnMulhsu
					return
				case 0x3:
					instruction.iType = InsnMulhu
					return
				case 0x4:
					instruction.iType = InsnDiv
					return
				case 0x5:
					instruction.iType = InsnDivu
					return
				case 0x6:
					instruction.iType = InsnRem
					return
				case 0x7:
					instruction.iType = InsnRemu
					return
				default:
					UnReachable()
				}
			case 0x20:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnSub
					return
				case 0x5:
					instruction.iType = InsnSra
					return
				default:
					UnReachable()
				}
				UnReachable()
			default:
				UnReachable()
			}

			UnReachable()
		case 0xd:
			InstrUTypeRead(instruction, data)
			instruction.iType = InsnLui
			return
		case 0xe:
			InstrRTypeRead(instruction, data)

			funct3 := Funct3(data)
			funct7 := Funct7(data)

			switch funct7 {
			case 0x0:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnAddw
					return
				case 0x1:
					instruction.iType = InsnSllw
					return
				case 0x5:
					instruction.iType = InsnSrlw
					return
				default:
					UnReachable()
				}
			case 0x1:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnMulw
					return
				case 0x4:
					instruction.iType = InsnDivw
					return
				case 0x5:
					instruction.iType = InsnDivuw
					return
				case 0x6:
					instruction.iType = InsnRemw
					return
				case 0x7:
					instruction.iType = InsnRemuw
					return
				default:
					UnReachable()
				}
			case 0x20:
				switch funct3 {
				case 0x0:
					instruction.iType = InsnSubw
					return
				case 0x5:
					instruction.iType = InsnSraw
					return
				default:
					UnReachable()
				}
			default:
				UnReachable()
			}
			UnReachable()
		case 0x10:
			funct2 := Funct2(data)
			InstrFprTypeRead(instruction, data)
			switch funct2 {
			case 0x0:
				instruction.iType = InsnFmaddS
				return
			case 0x1:
				instruction.iType = InsnFmaddD
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x11:
			funct2 := Funct2(data)

			InstrFprTypeRead(instruction, data)
			switch funct2 {
			case 0x0:
				instruction.iType = InsnFmsubS
				return
			case 0x1:
				instruction.iType = InsnFmsubD
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x12:
			funct2 := Funct2(data)

			InstrFprTypeRead(instruction, data)
			switch funct2 {
			case 0x0:
				instruction.iType = InsnFnmsubS
				return
			case 0x1:
				instruction.iType = InsnFnmsubD
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x13:
			funct2 := Funct2(data)
			InstrFprTypeRead(instruction, data)

			switch funct2 {
			case 0x0:
				instruction.iType = InsnFnmaddS
				return
			case 0x1:
				instruction.iType = InsnFnmaddD
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x14:
			funct7 := Funct7(data)

			InstrRTypeRead(instruction, data)
			switch funct7 {
			case 0x0:
				instruction.iType = InsnFaddS
				return
			case 0x1:
				instruction.iType = InsnFaddD
				return
			case 0x4:
				instruction.iType = InsnFsubS
				return
			case 0x5:
				instruction.iType = InsnFsubD
				return
			case 0x8:
				instruction.iType = InsnFmulS
				return
			case 0x9:
				instruction.iType = InsnFmulD
				return
			case 0xc:
				instruction.iType = InsnFdivS
				return
			case 0xd:
				instruction.iType = InsnFdivD
				return
			case 0x10:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFsgnjS
					return
				case 0x1:
					instruction.iType = InsnFsgnjnS
					return
				case 0x2:
					instruction.iType = InsnFsgnjxS
					return
				default:
					UnReachable()
				}
			case 0x11:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFsgnjD
					return
				case 0x1:
					instruction.iType = InsnFsgnjnD
					return
				case 0x2:
					instruction.iType = InsnFsgnjxD
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x14:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFminS
					return
				case 0x1:
					instruction.iType = InsnFmaxS
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x15:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFminD
					return
				case 0x1:
					instruction.iType = InsnFmaxD
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x20:
				assert(Rs2(data) == 1, "FCVT.S.D rs2 is not zero")
				instruction.iType = InsnFcvtSD
				return
			case 0x21:
				assert(Rs2(data) == 0, "FCVT.D.S rs2 is not zero")
				instruction.iType = InsnFcvtDS
				return
			case 0x2c:
				assert(Rs2(data) == 0, "FSQRT.S rs2 is not zero")
				instruction.iType = InsnFsqrtS
				return
			case 0x2d:
				assert(Rs2(data) == 0, "FSQRT.D rs2 is not zero")
				instruction.iType = InsnFsqrtD
				return
			case 0x50:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFleS
					return
				case 0x1:
					instruction.iType = InsnFltS
					return
				case 0x2:
					instruction.iType = InsnFeqS
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x51:
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFleD
					return
				case 0x1:
					instruction.iType = InsnFltD
					return
				case 0x2:
					instruction.iType = InsnFeqD
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x60:
				rs2 := Rs2(data)

				switch rs2 {
				case 0x0:
					instruction.iType = InsnFcvtWS
					return
				case 0x1:
					instruction.iType = InsnFcvtWuS
					return
				case 0x2:
					instruction.iType = InsnFcvtLS
					return
				case 0x3:
					instruction.iType = InsnFcvtLuS
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x61:
				rs2 := Rs2(data)
				switch rs2 {
				case 0x0:
					instruction.iType = InsnFcvtWD
					return
				case 0x1:
					instruction.iType = InsnFcvtWud
					return
				case 0x2:
					instruction.iType = InsnFcvtLD
					return
				case 0x3:
					instruction.iType = InsnFcvtLuD
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x68:
				rs2 := Rs2(data)

				switch rs2 {
				case 0x0:
					instruction.iType = InsnFcvtSW
					return
				case 0x1:
					instruction.iType = InsnFcvtSWu
					return
				case 0x2:
					instruction.iType = InsnFcvtSL
					return
				case 0x3:
					instruction.iType = InscFcvtSLu
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x69:
				rs2 := Rs2(data)

				switch rs2 {
				case 0x0:
					instruction.iType = InsnFcvtDW
					return
				case 0x1:
					instruction.iType = InsnFcvtDWu
					return
				case 0x2:
					instruction.iType = InsnFcvtDL
					return
				case 0x3:
					instruction.iType = InsnFcvtDLu
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x70:
				assert(Rs2(data) == 0, "case 0x70 rs2 is not zero")
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFmvXW
					return
				case 0x1:
					instruction.iType = InsnFclassS
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x71:
				assert(Rs2(data) == 0, "case 0x71 rs2 is not zero")
				funct3 := Funct3(data)

				switch funct3 {
				case 0x0:
					instruction.iType = InsnFmvXD
					return
				case 0x1:
					instruction.iType = InsnFclassD
					return
				default:
					UnReachable()
				}
				UnReachable()
			case 0x78:
				assert(Rs2(data) == 0, "FMV.W.X rs2 is not zero")
				instruction.iType = InsnFmvWX
				return
			case 0x79:
				assert(Rs2(data) == 0, "FMV.D.X rs2 is not zero")
				instruction.iType = InsnFmvDX
				return
			default:
				UnReachable()
			}
			UnReachable()
		case 0x18:
			InstrBTypeRead(instruction, data)

			funct3 := Funct3(data)
			switch funct3 {
			case 0x0: /* BEQ */
				instruction.iType = InsnBeq
				return
			case 0x1: /* BNE */
				instruction.iType = InsnBne
				return
			case 0x4: /* BLT */
				instruction.iType = InsnBlt
				return
			case 0x5: /* BGE */
				instruction.iType = InsnBge
				return
			case 0x6: /* BLTU */
				instruction.iType = InsnBltu
				return
			case 0x7:
				instruction.iType = InsnBgeu
			default:
				UnReachable()
			}
			UnReachable()
		case 0x19: /* JALR */
			InstrITypeRead(instruction, data)
			instruction.iType = InsnJalr
			instruction.cont = true
			return
		case 0x1b:
			InstrJTypeRead(instruction, data)
			instruction.iType = InsnJal
			instruction.cont = true
			return
		case 0x1c:
			if data == 0x73 {
				instruction.iType = InsnEcall
				instruction.cont = true
				return
			}

			funct3 := Funct3(data)
			InstrCsrTypeRead(instruction, data)
			switch funct3 {
			case 0x1:
				instruction.iType = InsnCsrrw
				return
			case 0x2:
				instruction.iType = InsnCsrrs
				return
			case 0x3:
				instruction.iType = InsnCsrrc
				return
			case 0x5:
				instruction.iType = InsnCsrrwi
				return
			case 0x6:
				instruction.iType = InsnCsrrsi
				return
			case 0x7:
				instruction.iType = InsnCsrrci
				return
			default:
				UnReachable()
			}
			UnReachable()
		default:
			UnReachable()
		}
		UnReachable()
	default:
		UnReachable()
	}
}
