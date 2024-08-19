package main

import (
	"os"
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
