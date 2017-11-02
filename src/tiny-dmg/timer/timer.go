package timer

import (
	"fmt"
	"tiny-dmg/cpu"
	"tiny-dmg/memory"
)

type Timer struct {
	doomsdayClock uint32
}

func NewTimer() *Timer {
	return new(Timer)
}

func (t *Timer) Update(gb *cpu.GbCpu, m *memory.Memory, cycles uint8) {
	t.updateTimer(gb, m, cycles)
	t.dividerHack(gb, m, cycles)
}

func (t *Timer) updateTimer(gb *cpu.GbCpu, m *memory.Memory, cycles uint8) {
	conf := m.GetRaw(memory.RegTimerControl)

	if conf&(1<<2) == 0 {
		// timer is disabled, nothing to do
		return
	}

	if t.timerFires(conf&0x03, cycles) {
		val := m.GetRaw(memory.RegTimerCounter)
		if val == 0xFF { // would overflow -> reset and fire interrupt
			val = m.GetRaw(memory.RegTimerModulo)
			m.WriteRawSet(memory.RegInterruptFlag, memory.BitIrTimer)
		} else {
			val++
		}
		m.WriteRaw(memory.RegTimerCounter, val)
	}
}

// timerFires returns true if we should increment the timer.
func (t *Timer) timerFires(conf uint8, cycles uint8) bool {
	steps := uint32(1)
	switch conf {
	case 0:
		// 1024 / 1024
	case 1:
		steps = 1024 / 16
	case 2:
		steps = 1024 / 64
	case 3:
		steps = 1024 / 256
	}
	t.doomsdayClock += uint32(cycles) * steps

	if t.doomsdayClock >= 1024 {
		t.doomsdayClock -= 1024 // fixme: is this correct?
		return true
	}
	return false
}

// Just a quick hack to get the divider register running, FIXME: Needs proper emulation.
func (t *Timer) dividerHack(gb *cpu.GbCpu, m *memory.Memory, cycles uint8) {
	if gb.ClockCycles%256 == 0 {
		v := m.GetByte(memory.RegDivider)
		v++
		m.WriteRaw(memory.RegDivider, v)
	}
}
