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
	instruction.imm = data >> 20
	instruction.rs1 = uint8(Rs1(data))
	instruction.rd = uint8(Rd(data))

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
			UnReachable()
		case 0x4:
			UnReachable()
		case 0x5:
			UnReachable()
		case 0x6:
			UnReachable()
		case 0x8:
			UnReachable()
		case 0xc:
			UnReachable()
		case 0xd:
			UnReachable()
		case 0xe:
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
