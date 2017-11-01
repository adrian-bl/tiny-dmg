package mbc

import (
	"fmt"
	"tiny-dmg/rom"
)

type Mbc0 struct {
}

func newMbc0() MemoryBankController {
	return new(Mbc0)
}

func (mbc *Mbc0) ReadFromRom(r rom.RomImage, addr uint16) uint8 {
	return r.GetByte(uint32(addr))
}

func (mbc *Mbc0) WriteToRom(addr uint16, val uint8) {
	fmt.Printf("MBC0: Write to ROM add %X (%X) ignored\n", addr, val)
}

func (mbc *Mbc0) WriteExternalRam(m RawMemoryAccess, addr uint16, val uint8) {
	m.WriteRaw(addr, val)
}
func (mbc *Mbc0) ReadExternalRam(m RawMemoryAccess, addr uint16) uint8 {
	return m.GetRaw(addr)
}

func (mbc *Mbc0) DisableBootRom(x *MemoryBankController, val uint8) {
	// noop
}
