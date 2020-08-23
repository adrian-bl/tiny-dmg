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

	t := r.RomType

	if t == 0 {
		mbc = newMbc0()
	} else if t == 0x1 || t == 0x2 || t == 0x3 {
		mbc = newMbc1()
	} else if t == 0x0F || t == 0x10 || t == 0x11 || t == 0x12 || t == 0x13 {
		mbc = newMbc3()
	} else {
		panic(fmt.Errorf("ROM requested MBC%d, but no such implementation is known.\n", r.RomType))
	}
	return newBios(b, mbc)
}
