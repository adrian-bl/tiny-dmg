package mbc

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
)

type Mbc3 struct {
	bank       uint8
	ramEnabled bool
}

func newMbc3() MemoryBankController {
	mbc := &Mbc3{}
	mbc.bank = 1
	mbc.ramEnabled = false
	return mbc
}

func (mbc *Mbc3) ReadFromRom(r *rom.RomImage, addr uint16) uint8 {
	real := uint32(addr)
	if addr >= 0x4000 {
		real = uint32(addr) + (0x4000 * uint32(mbc.bank-1))
	}
	if addr >= 0xA000 && addr <= 0xBFFF {
		panic(fmt.Errorf("Not implemented"))
	}
	return r.GetByte(real)
}

func (mbc *Mbc3) WriteToRom(addr uint16, val uint8) {
	if addr <= 0x1FFF {
		if val&0xA == 0xA {
			mbc.ramEnabled = true
		} else {
			mbc.ramEnabled = false
		}
	} else if addr >= 2000 && addr <= 0x3FFF {
		bank := val & 0x7F
		if bank == 0 {
			bank = 1
		}
		mbc.bank = bank
		fmt.Printf("ROM selects bank %d\n", bank)
	} else if addr >= 0x4000 && addr <= 0x5fff {
		panic(fmt.Errorf("not implemented yet"))
	} else if addr >= 0x6000 && addr <= 0x7fff {
		panic(fmt.Errorf("RTC not done yet"))
	} else {
		panic(fmt.Errorf("MBC1: Write %X TO %X\n", val, addr))
	}
}

func (mbc *Mbc3) WriteExternalRam(m RawMemoryAccess, addr uint16, val uint8) {
	if mbc.ramEnabled {
		m.WriteRaw(addr, val)
	}
}
func (mbc *Mbc3) ReadExternalRam(m RawMemoryAccess, addr uint16) uint8 {
	if mbc.ramEnabled {
		return m.GetRaw(addr)
	}
	return 0xFF
}

func (mbc *Mbc3) DisableBootRom(x *MemoryBankController, val uint8) {
	// noop
}
