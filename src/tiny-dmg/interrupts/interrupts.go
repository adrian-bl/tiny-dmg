package interrupts

import (
	"tiny-dmg/cpu"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
)

type Interrupts struct {
	lastScanline uint8
}

func New() *Interrupts {
	return new(Interrupts)
}

func (i *Interrupts) Update(gb *cpu.GbCpu, m *memory.Memory, l *lcd.Lcd) {
	scanline := m.GetByte(memory.RegCurrentScanline)

	if scanline != i.lastScanline && scanline == lcd.LastVisibleScanline+4 {
		//		fmt.Printf("----------------------YYXX: Scanline vblank !::: %d\n", scanline)
		//		if gb.InterruptsEnabled {
		//			gb.InterruptsEnabled = false
		//			gb.Reg.PC--
		//			cpu.Op_Rst(gb, 0x40)
		//		}
	}
	i.lastScanline = scanline
}
