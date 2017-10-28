package machine

import (
	"fmt"
	"time"
	"tiny-dmg/cpu"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
)

type Machine struct {
	cpu *cpu.GbCpu
	lcd *lcd.Lcd
	mem *memory.Memory
}

func New(cpu *cpu.GbCpu, lcd *lcd.Lcd, mem *memory.Memory) (mach *Machine, err error) {
	mach = new(Machine)
	mach.cpu = cpu
	mach.lcd = lcd
	mach.mem = mem
	return
}

func (mach *Machine) PowerOn() {
	mach.mem.PowerOn()
	mach.lcd.PowerOn()
	mach.cpu.PowerOn()
}

func (mach *Machine) Run() {
	fmt.Printf("Starting Z80 emulation, initial pc=%08X\n", mach.cpu.Reg.PC)

	op := byte(0)
	i := 1
	for {
		op = mach.mem.GetByte(mach.cpu.Reg.PC) // raw opcode from ROM
		mach.cpu.OpCode = cpu.OpCodes[op]

		fmt.Printf("%04X %02X                        SP=%04X      BC=%02X%02X       DE=%02X%02X    ", mach.cpu.Reg.PC, op, mach.cpu.Reg.SP, mach.cpu.Reg.B, mach.cpu.Reg.C, mach.cpu.Reg.D, mach.cpu.Reg.E)
		fmt.Printf("HL=%02X%02X    A=%02X F=%02X [", mach.cpu.Reg.H, mach.cpu.Reg.L, mach.cpu.Reg.A, mach.cpu.Reg.F)
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

		fmt.Printf("] c=%d ## %d, c=%d, LY(FF44) = %X, >> FIXME: STAT = %X (%X), LCDC=%02X, op=%s\n", mach.cpu.OpCode.ClockCycles, i, mach.cpu.ClockCycles, mach.mem.GetByte(0xFF44), mach.mem.GetByte(0xFF41), mach.mem.GetByte(0xFF41)&0x3, mach.mem.GetByte(0xFF40), mach.cpu.OpCode.Name)
		i++

		if mach.cpu.OpCode.Cback == nil {
			for {
				fmt.Printf("BREAKPOINT HIT AT %X -> WE ARE HANGING HERE....\n", mach.cpu.Reg.PC)
				time.Sleep(100 * time.Second)
			}
		}

		mach.cpu.OpCode.Cback(mach.cpu)
		mach.cpu.ClockCycles += uint32(mach.cpu.OpCode.ClockCycles)

		mach.lcd.Update(mach.cpu.OpCode.ClockCycles)

	}
}
