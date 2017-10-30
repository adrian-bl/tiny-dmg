package mbc

import (
	"fmt"
)

type Mbc1 struct {
	bank uint8
}

func NewMbc1() MemoryBankController {
	mbc := new(Mbc1)
	mbc.bank = 1
	return mbc
}

func (mbc *Mbc1) ReadBank4000(m RawMemoryAccess, addr uint16) uint8 {
	addr = addr + (0x4000 * uint16(mbc.bank-1)) // bank 1 already equals the expected base, so it's +0
	return m.GetRaw(addr)
}

func (mbc *Mbc1) WriteToRom(addr uint16, val uint8) {
	if addr >= 2000 && addr <= 0x3FFF {
		bank := val & 0x1F
		if bank == 0 {
			bank = 1
		}
		mbc.bank = bank
		fmt.Printf("ROM selects bank %d\n", bank)
	} else {
		panic(fmt.Errorf("MBC1: Write %X TO %X\n", val, addr))
	}
}
