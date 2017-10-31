package mbc

import (
	"fmt"
	"tiny-dmg/rom"
)

type MemoryBankController interface {
	ReadFromRom(rom.RomImage, uint16) byte
	WriteToRom(uint16, uint8)
	WriteExternalRam(RawMemoryAccess, uint16, uint8)
	ReadExternalRam(RawMemoryAccess, uint16) byte
}

type RawMemoryAccess interface {
	GetRaw(uint16) byte
	WriteRaw(uint16, byte)
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
