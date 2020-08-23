package mbc

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
)

type Mbc3 struct {
	romBank    uint8
	sramBank   uint8
	ramEnabled bool
	clockLatch uint8
}

func newMbc3() MemoryBankController {
	mbc := &Mbc3{}
	mbc.romBank = 1
	mbc.ramEnabled = false
	return mbc
}

func (mbc *Mbc3) ReadFromRom(r *rom.RomImage, addr uint16) uint8 {
	real := uint32(addr)
	if addr >= 0x4000 {
		real = uint32(addr) + (0x4000 * uint32(mbc.romBank-1))
	}
	if addr >= 0xA000 && addr <= 0xBFFF {
		// handled by memory.go, should never reach this.
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
		romBank := val & 0x7F
		if romBank == 0 {
			romBank = 1
		}
		mbc.romBank = romBank
	} else if addr >= 0x4000 && addr <= 0x5fff {
		mbc.sramBank = val
	} else if addr >= 0x6000 && addr <= 0x7fff {
		if mbc.clockLatch == 0 && val == 1 {
			fmt.Printf("TODO: latch RTC\n")
		}
		mbc.clockLatch = val
	} else {
		panic(fmt.Errorf("MBC3: Write %X TO %X\n", val, addr))
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
