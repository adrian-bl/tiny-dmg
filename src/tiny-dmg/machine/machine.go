package machine

import (
	"fmt"
	"time"
	"tiny-dmg/cpu"
	"tiny-dmg/interrupts"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
)

type Machine struct {
	cpu *cpu.GbCpu
	lcd *lcd.Lcd
	mem *memory.Memory
	itr *interrupts.Interrupts
}

func New(c *cpu.GbCpu, m *memory.Memory, l *lcd.Lcd) (mach *Machine, err error) {
	mach = new(Machine)
	mach.cpu = c
	mach.lcd = l
	mach.mem = m
	mach.itr = interrupts.New()
	return
}

func (mach *Machine) PowerOn() {
	mach.mem.PowerOn()
	mach.lcd.PowerOn()
	mach.cpu.PowerOn()
}

var XLOG = "mgba"

func (mach *Machine) Run() {
	fmt.Printf("Starting Z80 emulation, initial pc=%08X\n", mach.cpu.Reg.PC)

	op := byte(0)
	i := 1
	for {
		op = mach.mem.GetByte(mach.cpu.Reg.PC) // raw opcode from ROM
		dbgopcode := cpu.OpCodes[op]

		if XLOG == "verbose" {
			fmt.Printf("%04X %02X AF=%02X%02X      BC=%02X%02X       DE=%02X%02X    ", mach.cpu.Reg.PC, mach.cpu.Reg.A, mach.cpu.Reg.F, mach.cpu.Reg.B, mach.cpu.Reg.C, mach.cpu.Reg.D, mach.cpu.Reg.E)
			fmt.Printf("HL=%02X%02X   SP=%04X [", mach.cpu.Reg.H, mach.cpu.Reg.L, mach.cpu.Reg.SP)
			if mach.cpu.Reg.F&cpu.FlagZ != 0 {
				fmt.Printf("Z")
			} else {
				fmt.Printf("-")
			}
			if mach.cpu.Reg.F&cpu.FlagN != 0 {
				fmt.Printf("N")
			} else {
				fmt.Printf("-")
			}
			if mach.cpu.Reg.F&cpu.FlagH != 0 {
				fmt.Printf("H")
			} else {
				fmt.Printf("-")
			}
			if mach.cpu.Reg.F&cpu.FlagC != 0 {
				fmt.Printf("C")
			} else {
				fmt.Printf("-")
			}

			fmt.Printf("] c=%d ## %d, c=%d, LY(FF44) = %X, >> FIXME: STAT = %X (%X), LCDC=%02X, op=%s\n", dbgopcode.ClockCycles, i, mach.cpu.ClockCycles, mach.mem.GetByte(0xFF44), mach.mem.GetByte(0xFF41), mach.mem.GetByte(0xFF41)&0x3, mach.mem.GetByte(0xFF40), dbgopcode.Name)
		}

		if XLOG == "mgba" {
			fmt.Printf("A: %02X F: %02X B: %02X C: %02X D: %02X E: %02X H: %02X L: %02X SP: %04X PC: %04X |\n",
				mach.cpu.Reg.A, mach.cpu.Reg.F, mach.cpu.Reg.B, mach.cpu.Reg.C,
				mach.cpu.Reg.D, mach.cpu.Reg.E, mach.cpu.Reg.H, mach.cpu.Reg.L, mach.cpu.Reg.SP, mach.cpu.Reg.PC)
		}

		i++

		if dbgopcode.Callback == nil {
			for {
				fmt.Printf("BREAKPOINT HIT AT %X -> WE ARE HANGING HERE....\n", mach.cpu.Reg.PC)
				time.Sleep(100 * time.Second)
			}
		}

		cycles := mach.cpu.Execute(op)

		mach.cpu.ClockCycles += uint32(cycles)
		mach.lcd.Update(cycles, mach.CallbackVblank)
		mach.itr.Update(mach.cpu, mach.mem, mach.lcd)
	}
}

func (mach *Machine) CallbackVblank() {
	if mach.cpu.InterruptsEnabled {
		mach.cpu.InterruptsEnabled = false
		//	mach.cpu.Reg.PC--
		//	cpu.Op_Rst(mach.cpu, 0x40)
	}
}
