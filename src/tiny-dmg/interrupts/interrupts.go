package interrupts

import (
	"tiny-dmg/cpu"
	"tiny-dmg/memory"
)

// InterruptServiceRoutine addresses
const (
	IsrVblank    = 0x40
	IsrLcdStatus = 0x48
	IsrTimer     = 0x50
	IsrSerial    = 0x58
	IsrJoypad    = 0x60
)

type Interrupts struct {
}

func New() *Interrupts {
	return new(Interrupts)
}

func (i *Interrupts) Update(gb *cpu.GbCpu, m *memory.Memory) {

	if gb.InterruptsEnabled && m.GetByte(memory.RegInterruptEnable)&memory.BitIrVblank != 0 && m.GetByte(memory.RegInterruptFlag)&memory.BitIrVblank != 0 {
		gb.InterruptsEnabled = false
		m.WriteRawClear(memory.RegInterruptFlag, memory.BitIrVblank)
		gb.Reg.PC-- // Op_Rst counts+1, to skip over itself. However: This is not a real opcode, so we first step one back.
		cpu.Op_Rst(gb, IsrVblank)
	}
}
