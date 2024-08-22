package main

// 在go语言中逻辑右移还是算术右移是通过远算数的类型来决定的
func GetQuadrant(data uint32) uint32 {
	return ((data) >> 0) & 0x3
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
		Fatal("unimplemented")
	default:
		UnReachable()
	}
}
