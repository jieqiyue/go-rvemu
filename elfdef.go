package main

// EiNident 定义ELF文件的魔数字节的数量
const EiNident = 16

// ElfMagic 定义ELF文件的魔术字节为全局常量
const ElfMagic = "\x7fELF"

// EmRiscV 规定了CPU架构
const EmRiscV = 243

// EiClass Elf文件位数相关
const EiClass = 4
const ElfClassNone = 0
const ElfClass32 = 1
const ElfClass64 = 2
const ElfClassSum = 3

// Elf64PhdrPTypeLoad 代表这个段是否是要加载到内存的
const Elf64PhdrPTypeLoad = 1

type Elf64EhdrT struct {
	EIdent     [EiNident]uint8
	EType      uint16
	EMachine   uint16
	EVersion   uint32
	EEntry     uint64
	EPhoff     uint64
	EShoff     uint64
	EFlags     uint32
	EEhsize    uint16
	EPhentsize uint16
	EPhnum     uint16
	EShentsize uint16
	EShnum     uint16
	EShstrndx  uint16
}

type Elf64PhdrT struct {
	PType   uint32
	PFlags  uint32
	POffset uint64
	PVaddr  uint64
	PPaddr  uint64
	PFilesz uint64
	PMemsz  uint64
	PAlign  uint64
}
