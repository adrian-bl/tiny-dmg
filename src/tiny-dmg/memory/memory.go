package memory

import (
	"fmt"
	"io/ioutil"
	"tiny-dmg/rom"
)

type Memory struct {
	memory [0x10000]byte // memory ranges from 0x0000 - 0xFFFF
}

func New(r rom.RomImage) (m *Memory, err error) {
	m = new(Memory)
	copy(m.memory[0:], r.Blob[0:]) // fixme: maybe we should start at 0x100 ?
	return
}

func (m *Memory) PowerOn() {
	m.WriteByte(RegJoypadInput, 0xCF) // FIXME: Remove once we have joypad emulation
	m.WriteByte(0xFF05, 0x00)
	m.WriteByte(0xFF06, 0x00)
	m.WriteByte(0xFF07, 0x00)
	m.WriteByte(0xFF10, 0x80)
	m.WriteByte(0xFF11, 0xBF)
	m.WriteByte(0xFF12, 0xF3)
	m.WriteByte(0xFF14, 0xBF)
	m.WriteByte(0xFF16, 0x3F)
	m.WriteByte(0xFF17, 0x00)
	m.WriteByte(0xFF19, 0xBF)
	m.WriteByte(0xFF1A, 0x7F)
	m.WriteByte(0xFF1B, 0xFF)
	m.WriteByte(0xFF1C, 0x9F)
	m.WriteByte(0xFF1E, 0xBF)
	m.WriteByte(0xFF20, 0xFF)
	m.WriteByte(0xFF21, 0x00)
	m.WriteByte(0xFF22, 0x00)
	m.WriteByte(0xFF23, 0xBF)
	m.WriteByte(0xFF24, 0x77)
	m.WriteByte(0xFF25, 0xF3)
	m.WriteByte(0xFF26, 0xF1)
	m.WriteByte(RegLcdControl, 0x91)
	m.WriteByte(RegLcdState, 0x84)
	m.WriteByte(RegScrollY, 0x00)
	m.WriteByte(RegScrollX, 0x00)
	m.WriteByte(RegCurrentScanline, 0x00)
	m.WriteByte(RegLYCompare, 0x00)
	m.WriteByte(0xFF47, 0xFC)
	m.WriteByte(0xFF48, 0xFF)
	m.WriteByte(0xFF49, 0xFF)
	m.WriteByte(0xFF4A, 0x00)
	m.WriteByte(0xFF4B, 0x00)
	m.WriteByte(0xFFFF, 0x00)
	fmt.Printf("# memory initialized\n")
}

func (m *Memory) GetByte(addr uint16) byte {
	return m.memory[addr]
}

func (m *Memory) WriteByte(addr uint16, val byte) {

	if addr < 0x8000 {
		fmt.Printf("Ignoring bougous write of %X to %X (RO memory!)\n", val, addr)
		return
	}

	if addr >= 0xE000 && (addr < 0xFE00) {
		fmt.Printf("Implement me\n")
		panic(nil)
	}

	if (addr >= 0xFEA0) && (addr < 0xFEFF) {
		fmt.Printf("Ignoring bougous write of %X to %X (Restricted memory!)\n", val, addr)
		return
	}

	if RegJoypadInput == addr {
		switch val {
		case 0x10:
			val = 0xDF
		case 0x20:
			val = 0xEF
		case 0x30:
			val = 0xFF
		default:
			fmt.Printf("Unexpected joypad write: %X\n", val)
		}
		fmt.Printf("Write to joypad reg, faking write to be: %X\n", val)
	}

	if RegDivider == addr {
		fmt.Printf("Write to divider register!\n")
		panic(nil)
	}

	if RegCurrentScanline == addr {
		fmt.Printf("Write to scanline register -> RESETTING SCANLINE VALUE\n", val)
		val = 0
	}
	if RegLcdState == addr {
		fmt.Printf("LCD WRITE: %d\n", val)
	}

	if RegDoDMA == addr {
		panic(fmt.Errorf("DMA START: %X\n", val))
	}

	m.memory[addr] = val
}

func (m *Memory) WriteRaw(addr uint16, val byte) {
	m.memory[addr] = val
}

func (m *Memory) Dump() {
	//	fmt.Printf("%v\n", m.memory[0x8000:0xa000])
	ioutil.WriteFile("/tmp/x.data", m.memory[0x0:0x10000], 0644)

}
