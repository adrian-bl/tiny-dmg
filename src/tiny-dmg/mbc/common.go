package mbc

import (
	"fmt"
)

type MemoryBankController interface {
	ReadBank4000(RawMemoryAccess, uint16) byte
	WriteToRom(uint16, uint8)
}

type RawMemoryAccess interface {
	GetRaw(uint16) byte
}

func GetMbc(t uint8) MemoryBankController {
	switch t {
	case 0:
		return NewMbc0()
	case 1:
		return NewMbc1()
	default:
		panic(fmt.Errorf("ROM requested MBC%d, but no such implementation is known.\n", t))
	}
}
