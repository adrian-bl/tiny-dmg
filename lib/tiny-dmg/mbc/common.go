package mbc

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
)

type MemoryBankController interface {
	ReadFromRom(*rom.RomImage, uint16) byte
	WriteToRom(uint16, uint8)
	WriteExternalRam(RawMemoryAccess, uint16, uint8)
	ReadExternalRam(RawMemoryAccess, uint16) byte
	DisableBootRom(*MemoryBankController, uint8)
}

type RawMemoryAccess interface {
	GetRaw(uint16) byte
	WriteRaw(uint16, byte)
}

func GetMbc(b, r *rom.RomImage) MemoryBankController {
	var mbc MemoryBankController
	switch r.RomType {
	case 0:
		mbc = newMbc0()
	case 1:
		mbc = newMbc1()
	case 3:
		mbc = newMbc3()
	default:
		panic(fmt.Errorf("ROM requested MBC%d, but no such implementation is known.\n", r.RomType))
	}
	return newBios(b, mbc)
}
