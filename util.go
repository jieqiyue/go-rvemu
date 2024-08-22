package main

import (
	"os"
	"unsafe"
)

func assert(condition bool, message string) {
	if !condition {
		EPrintf("assert fail, %v\n", message)
		os.Exit(1) // 终止程序
	}
}

func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}

func UnReachable() {
	panic("UnReachable Place")
}

func Fatal(message string) {
	//panic(message)
	EPrintf(message)
}

type MemSize uint64

const (
	Byte  MemSize = 0xff
	Word  MemSize = 0xffff
	Dword MemSize = 0xffffffff
	Qword MemSize = 0xffffffffffffffff
)

func GetProcessMemory(address uint64, len uint64) uint64 {
	var addrs = unsafe.Pointer(uintptr(address))
	ptr := (*uint64)(addrs)

	origin := *ptr
	return origin & len
}
