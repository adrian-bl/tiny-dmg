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

// Upate fires all pending interrupts.
func (i *Interrupts) Update(gb *cpu.GbCpu, m *memory.Memory) {
	// Order matters: Vblank has highes prio, joypad lowest
	handle(gb, m, memory.BitIrVblank, IsrVblank)
	handle(gb, m, memory.BitIrLcdStatus, IsrLcdStatus)
	handle(gb, m, memory.BitIrTimer, IsrTimer)
	handle(gb, m, memory.BitIrSerial, IsrSerial)
	handle(gb, m, memory.BitIrJoypad, IsrJoypad)
}

// handle handles the actual interrupt work, that is: checking whether
// the interrupt should fire and call to the ISR if it should.
func handle(gb *cpu.GbCpu, m *memory.Memory, check uint8, isr uint16) {
	if gb.InterruptsEnabled && m.GetByte(memory.RegInterruptEnable)&check != 0 && m.GetByte(memory.RegInterruptFlag)&check != 0 {
		gb.InterruptsEnabled = false
		m.WriteRawClear(memory.RegInterruptFlag, check)
		gb.Reg.PC-- // Op_Rst counts+1, to skip over itself. However: This is not a real opcode, so we first step one back.
		cpu.Op_Rst(gb, isr)
	}
}
