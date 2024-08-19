package main

import "log"

// Debugging
const Debug = true
const Error = true

func DPrintf(format string, a ...interface{}) {
	if Debug {
		log.Printf(format, a...)
	}
}

func EPrintf(format string, a ...interface{}) {
	if Error {
		log.Printf(format, a...)
	}
}
