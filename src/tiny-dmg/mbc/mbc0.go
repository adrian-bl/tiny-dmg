package mbc

import (
	"fmt"
)

type Mbc0 struct {
}

func NewMbc0() MemoryBankController {
	return new(Mbc0)
}

func (mbc *Mbc0) ReadBank4000(m RawMemoryAccess, addr uint16) uint8 {
	return m.GetRaw(addr)
}

func (mbc *Mbc0) WriteToRom(addr uint16, val uint8) {
	fmt.Printf("MBC0: Write to ROM add %X (%X) ignored\n", addr, val)
}
