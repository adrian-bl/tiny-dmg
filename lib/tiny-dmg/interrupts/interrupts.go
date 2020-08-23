package interrupts

import (
	"log"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/cpu"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/memory"
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
	// Track enabled AND fired interrupts.
	flags := m.GetByte(memory.RegInterruptEnable) & m.GetByte(memory.RegInterruptFlag)
	if flags != 0 {
		// Disable halt if ANY interrupt could fire.
		gb.Halted = false
	}
	if gb.InterruptsEnabled {
		// Order matters: Vblank has highes prio, joypad lowest. Note that only one interrupt can fire per invocation of Update.
		switch {
		case flags&memory.BitIrVblank != 0:
			handle(gb, m, memory.BitIrVblank, IsrVblank)
		case flags&memory.BitIrLcdStatus != 0:
			handle(gb, m, memory.BitIrLcdStatus, IsrLcdStatus)
		case flags&memory.BitIrTimer != 0:
			handle(gb, m, memory.BitIrTimer, IsrTimer)
		case flags&memory.BitIrSerial != 0:
			handle(gb, m, memory.BitIrSerial, IsrSerial)
		case flags&memory.BitIrJoypad != 0:
			handle(gb, m, memory.BitIrJoypad, IsrJoypad)
		case flags != 0:
			log.Printf("InterruptFlag enabled but nothing fired?!\n")
		}
	}
}

// handle handles the actual interrupt work, that is: checking whether
// the interrupt should fire and call to the ISR if it should.
func handle(gb *cpu.GbCpu, m *memory.Memory, check uint8, isr uint16) {
	gb.InterruptsEnabled = false
	m.WriteRawClear(memory.RegInterruptFlag, check)
	gb.Reg.PC-- // Op_Rst counts+1, to skip over itself. However: This is not a real opcode, so we first step one back.
	cpu.Op_Rst(gb, isr)
}
