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
			UnReachable()
		case 0x11:
			UnReachable()
		case 0x12:
			UnReachable()
		case 0x13:
			UnReachable()
		case 0x14:
			UnReachable()
		case 0x18:
			UnReachable()
		case 0x19:
			UnReachable()
		case 0x1b:
			UnReachable()
		case 0x1c:
			UnReachable()
		default:
			UnReachable()
		}
		UnReachable()
	default:
		UnReachable()
	}
}
