package memory

import (
	"fmt"
	"io/ioutil"
	"tiny-dmg/joypad"
	"tiny-dmg/mbc"
	"tiny-dmg/rom"
)

type Memory struct {
	memory [0x10000]byte // memory ranges from 0x0000 - 0xFFFF
	joypad *joypad.Joypad
	mbc    mbc.MemoryBankController
	rom    rom.RomImage
}

func New(r rom.RomImage, j *joypad.Joypad) (m *Memory, err error) {
	m = new(Memory)
	copy(m.memory[0:], r.GetBytes())
	m.mbc = mbc.GetMbc(r.RomType)
	m.joypad = j
	m.rom = r
	return
}

func (m *Memory) PowerOn() {
	//m.WriteByte(RegJoypadInput, 0xCF) // FIXME: Remove once we have joypad emulation
	m.WriteRaw(0xFF05, 0x00)
	m.WriteRaw(0xFF06, 0x00)
	m.WriteRaw(0xFF07, 0x00)
	m.WriteRaw(0xFF10, 0x80)
	m.WriteRaw(0xFF11, 0xBF)
	m.WriteRaw(0xFF12, 0xF3)
	m.WriteRaw(0xFF14, 0xBF)
	m.WriteRaw(0xFF16, 0x3F)
	m.WriteRaw(0xFF17, 0x00)
	m.WriteRaw(0xFF19, 0xBF)
	m.WriteRaw(0xFF1A, 0x7F)
	m.WriteRaw(0xFF1B, 0xFF)
	m.WriteRaw(0xFF1C, 0x9F)
	m.WriteRaw(0xFF1E, 0xBF)
	m.WriteRaw(0xFF20, 0xFF)
	m.WriteRaw(0xFF21, 0x00)
	m.WriteRaw(0xFF22, 0x00)
	m.WriteRaw(0xFF23, 0xBF)
	m.WriteRaw(0xFF24, 0x77)
	m.WriteRaw(0xFF25, 0xF3)
	m.WriteRaw(0xFF26, 0xF1)
	m.WriteRaw(RegLcdControl, 0x91)
	m.WriteRaw(RegLcdState, 0x81)
	m.WriteRaw(RegScrollY, 0x00)
	m.WriteRaw(RegScrollX, 0x00)
	m.WriteRaw(RegCurrentScanline, 0x00)
	m.WriteRaw(RegLYCompare, 0x00)
	m.WriteRaw(0xFF47, 0xFC)
	m.WriteRaw(0xFF48, 0xFF)
	m.WriteRaw(0xFF49, 0xFF)
	m.WriteRaw(0xFF4A, 0x00)
	m.WriteRaw(0xFF4B, 0x00)
	m.WriteRaw(0xFFFF, 0x00)
	fmt.Printf("# memory initialized\n")
}

func (m *Memory) GetByte(addr uint16) byte {
	if addr <= 0x7FFF {
		return m.mbc.ReadFromRom(m.rom, addr)
	} else if addr <= 0x9FFF {
		// Vram read
	} else if addr <= 0xBFFF {
		return m.mbc.ReadExternalRam(m, addr)
	} else if addr <= 0xCFFF {
		// Work RAM bank 0
	} else if addr <= 0xDFFF {
		// Work RAM bank 1
	} else if addr <= 0xFDFF {
		panic(fmt.Errorf("Mirror of C000~DDFF, not implemented yet!"))
	} else if addr <= 0xFE9F {
		// OAM
	} else if addr <= 0xFEFF {
		panic(fmt.Errorf("Read from unusable region: %X\n"))
	} else if addr <= 0xFF7F {
		switch addr {
		case RegJoypadInput:
			return m.joypad.GetJoypadByte(m.memory[addr])
		}
	} else if addr <= 0xFFFE {
		// HRAM
	} else {
		// IE
	}

	return m.memory[addr]
}

func (m *Memory) WriteByte(addr uint16, val byte) {
	if addr <= 0x7FFF {
		fmt.Printf("Write to ROM bank %X\n", addr)
		m.mbc.WriteToRom(addr, val)
		return
	} else if addr <= 0x9FFF {
		fmt.Printf("VRAM write: %X = %X\n", addr, val)
	} else if addr <= 0xBFFF {
		m.mbc.WriteExternalRam(m, addr, val)
		return
	} else if addr <= 0xCFFF {
		fmt.Printf("Write to Work RAM? %X = %X\n", addr, val)
	} else if addr <= 0xDFFF {
		fmt.Printf("Write to work RAM 1~N? %X = %X\n", addr, val)
	} else if addr <= 0xFDFF {
		panic(fmt.Errorf("Mirror of C000~DDFF, not implemented yet!"))
	} else if addr <= 0xFE9F {
		fmt.Printf("Write to OAM: %X = %X\n", addr, val)
	} else if addr <= 0xFEFF {
		fmt.Printf("Write to unusable region: %X\n", addr)
		return
	} else if addr <= 0xFF7F {
		fmt.Printf("Write to IO register %X = %X\n", addr, val)
		switch addr {
		case RegJoypadInput:
			m.regWriteJoypad(val)
			return
		case RegCurrentScanline:
			m.regWriteCurrentScanline(val)
			return
		case RegLcdState:
			m.regWriteLcdState(val)
			return
		case RegDoDMA:
			m.regWriteDoDma(val)
			return
		case RegDivider:
			val = 0
			// does NOT return: write just resets it.
		default:
			fmt.Printf("IO reg %X not handled, write %X will happen...\n", addr, val)
		}
	} else if addr <= 0xFFFE {
		fmt.Printf("Write to HRAM: %X = %X\n", addr, val)
	} else {
		fmt.Printf("Write to interrupts enable register! %X = %X\n", addr, val)
	}

	m.memory[addr] = val
}

func (m *Memory) regWriteJoypad(val byte) {
	m.memory[RegJoypadInput] &^= 0x30
	m.memory[RegJoypadInput] |= val & 0x30
}

func (m *Memory) regWriteDoDma(val byte) {
	// FIXME: This isn't free. we should count up cycles
	src := uint16(val) << 8 // val is divided by 0x100
	for i := uint16(0); i < 0xA0; i++ {
		m.memory[StartOamRange+i] = m.memory[src+i]
	}
	fmt.Printf("+++ DMA Transfer from %X done\n", src)
}

func (m *Memory) regWriteCurrentScanline(val byte) {
	fmt.Printf("+++ write to scanline register, setting it to ZERO")
	m.memory[RegCurrentScanline] = 0
}

func (m *Memory) regWriteLcdState(val byte) {
	fmt.Printf("+++ LCD STATUS WRITE: IGNORED: %X\n", val)
}

// WriteRaw writes a byte to the raw memory location, without
// any checks or mappings.
func (m *Memory) WriteRaw(addr uint16, val byte) {
	m.memory[addr] = val
}

func (m *Memory) GetRaw(addr uint16) byte {
	return m.memory[addr]
}

// WriteRawSet sets given bits at a raw memory location.
func (m *Memory) WriteRawSet(addr uint16, mask byte) {
	m.memory[addr] |= mask
}

// WriteRawClear clears given bits in a raw memory location.
func (m *Memory) WriteRawClear(addr uint16, mask byte) {
	m.memory[addr] &^= mask
}

func (m *Memory) Dump() {
	//	fmt.Printf("%v\n", m.memory[0x8000:0xa000])
	ioutil.WriteFile("/tmp/x.data", m.memory[0x0:0x10000], 0644)

}
