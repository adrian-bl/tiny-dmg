package mbc

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
)

type Mbc1 struct {
	bank       uint8
	ramEnabled bool
}

func newMbc1() MemoryBankController {
	mbc := new(Mbc1)
	mbc.bank = 1
	mbc.ramEnabled = false
	return mbc
}

func (mbc *Mbc1) ReadFromRom(r *rom.RomImage, addr uint16) uint8 {
	real := uint32(addr)
	if addr >= 0x4000 && addr <= 0x7fff {
		real = uint32(addr) + (0x4000 * uint32(mbc.bank-1))
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		// handled by memory, should never reach this.
		panic(nil)
	}
	return r.GetByte(real)
}

func (mbc *Mbc1) WriteToRom(addr uint16, val uint8) {
	if addr <= 0x1FFF {
		if val&0xA == 0xA {
			mbc.ramEnabled = true
		} else {
			mbc.ramEnabled = false
		}
	} else if addr >= 2000 && addr <= 0x3FFF {
		bank := val & 0x1F
		switch bank {
		case 0:
			bank = 1
		case 0x20:
			fallthrough
		case 0x40:
			fallthrough
		case 0x60:
			bank++
		}
		mbc.bank = bank
		fmt.Printf("ROM selects bank %d\n", bank)
	} else {
		panic(fmt.Errorf("MBC1: Write %X TO %X\n", val, addr))
	}
}

func (mbc *Mbc1) WriteExternalRam(m RawMemoryAccess, addr uint16, val uint8) {
	if mbc.ramEnabled {
		m.WriteRaw(addr, val)
	}
}
func (mbc *Mbc1) ReadExternalRam(m RawMemoryAccess, addr uint16) uint8 {
	if mbc.ramEnabled {
		return m.GetRaw(addr)
	}
	return 0xFF
}

func (mbc *Mbc1) DisableBootRom(x *MemoryBankController, val uint8) {
	// noop
}
