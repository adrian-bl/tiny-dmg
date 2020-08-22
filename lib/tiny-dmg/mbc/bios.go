package mbc

import (
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
)

type Bios struct {
	bios []byte
	mbc  MemoryBankController
}

// newBios returns a new BIOS MBC. A BIOS MBC returns the loaded
// bios but dispatches every read beyond the bios to the original ROM.
// It is also the only MBC to implement the DisableBootRom() call
// which will cause itself to be destroyed and replaced with the real ROM.
func newBios(r *rom.RomImage, mbc MemoryBankController) MemoryBankController {
	b := new(Bios)
	b.bios = r.GetBytes()
	b.mbc = mbc
	return b
}

// ReadFromRom reads given address from the ROM.
func (b *Bios) ReadFromRom(r *rom.RomImage, addr uint16) uint8 {
	if addr < uint16(len(b.bios)) {
		return b.bios[addr]
	}
	return b.mbc.ReadFromRom(r, addr)
}

// WriteToRom dispatches a write request to a MBC.
func (b *Bios) WriteToRom(addr uint16, val uint8) {
	b.mbc.WriteToRom(addr, val)
}

// WriteExternalRam dispatches a write request to a MBC.
func (b *Bios) WriteExternalRam(m RawMemoryAccess, addr uint16, val uint8) {
	b.mbc.WriteExternalRam(m, addr, val)
}

// ReadExternalRam dispatches a read request to a MBC.
func (b *Bios) ReadExternalRam(m RawMemoryAccess, addr uint16) uint8 {
	return b.mbc.ReadExternalRam(m, addr)
}

// DisableBootRom causes the *real* mbc to get loaded into &mbc.
func (b *Bios) DisableBootRom(mbc *MemoryBankController, val uint8) {
	*mbc = b.mbc
}
