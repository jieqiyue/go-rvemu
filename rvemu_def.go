package main

// ExitReason 指令执行跳出的原因
type ExitReason string

const (
	None           ExitReason = "None"
	DirectBranch   ExitReason = "DirectBranch"
	InDirectBranch ExitReason = "InDirectBranch"
	ECall          ExitReason = "ECall"
)

type InsnType string

const (
	InsnAddi InsnType = "InsnAddi"
	NumInsns InsnType = "NumInsns"
)

// GpRegType 32个通用寄存器
type GpRegType int

const (
	zero GpRegType = iota
	ra
	sp
	gp
	tp

	t0
	t1
	t2

	s0
	s1

	a0
	a1
	a2
	a3
	a4
	a5
	a6
	a7

	s2
	s3
	s4
	s5
	s6
	s7
	s8
	s9
	s10
	s11

	t3
	t4
	t5
	t6
)

type State struct {
	exitReason ExitReason
	gpRegs     [32]uint64
	pc         uint64
}

type Instruction struct {
	rd    uint8
	rs1   uint8
	rs2   uint8
	iType FuncName
	rvc   bool
	cont  bool
}
