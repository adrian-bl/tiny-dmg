package cpu

import (
	"fmt"
	"time"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
)

type GbCpu struct {
	InterruptsEnabled bool
	Reg               Registers
	Mem               *memory.Memory
	Lcd               *lcd.Lcd
	ClockCycles            uint32
	OpCode            OpEntry
}

const (
	FlagZ    = uint8(1 << 7)
	FlagN    = uint8(1 << 6)
	FlagH    = uint8(1 << 5)
	FlagC    = uint8(1 << 4)
	FlagMask = uint8(FlagZ | FlagN | FlagH | FlagC)
)

const (
	MaxCyclesPerSecond = 69905 // how many clock cycles the GB can perform in a second
)

type Registers struct {
	A  uint8
	F  uint8
	B  uint8
	C  uint8
	D  uint8
	E  uint8
	H  uint8
	L  uint8
	SP uint16
	PC uint16
}

func New(m *memory.Memory, l *lcd.Lcd) (gb GbCpu, err error) {
	gb.Reg.SP = 0xFFFE
	gb.Reg.PC = 0x0100

	gb.Reg.A = 0x01 // We are a 1stgen gameboy
	gb.Reg.F = FlagZ | FlagH | FlagC
	gb.Reg.C = 0x13
	gb.Reg.E = 0xD8
	gb.Reg.H = 0x01
	gb.Reg.L = 0x4D

	gb.Mem = m
	gb.Lcd = l
	return
}

func (gb *GbCpu) Boot() {
	fmt.Printf("Starting Z80 emulation, initial pc=%08X\n", gb.Reg.PC)

	op := byte(0)
	i := 1
	for {
		op = gb.Mem.GetByte(gb.Reg.PC) // raw opcode from ROM
		gb.OpCode = OpCodes[op]

		fmt.Printf("%04X %02X                        SP=%04X      BC=%02X%02X       DE=%02X%02X    ", gb.Reg.PC, op, gb.Reg.SP, gb.Reg.B, gb.Reg.C, gb.Reg.D, gb.Reg.E)
		fmt.Printf("HL=%02X%02X    A=%02X F=%02X [", gb.Reg.H, gb.Reg.L, gb.Reg.A, gb.Reg.F)
		if gb.Reg.F&FlagZ != 0 {
			fmt.Printf("Z")
		} else {
			fmt.Printf("-")
		}
		if gb.Reg.F&FlagN != 0 {
			fmt.Printf("N")
		} else {
			fmt.Printf("-")
		}
		if gb.Reg.F&FlagH != 0 {
			fmt.Printf("H")
		} else {
			fmt.Printf("-")
		}
		if gb.Reg.F&FlagC != 0 {
			fmt.Printf("C")
		} else {
			fmt.Printf("-")
		}

		fmt.Printf("]  op=%-18s, c=%d ## %d, c=%d, ff44 = %X, >> FIXME: STAT = %X, LCDC=%02X\n", gb.OpCode.Name, gb.OpCode.ClockCycles, i, gb.ClockCycles, gb.Mem.GetByte(0xFF44), gb.Mem.GetByte(0xFF41), gb.Mem.GetByte(0xFF40))
		i++

		if gb.OpCode.Cback == nil {
			for {
				fmt.Printf("BREAKPOINT HIT AT %X -> WE ARE HANGING HERE....\n", gb.Reg.PC)
				time.Sleep(100 * time.Second)
			}
		}

		gb.OpCode.Cback(gb)
		gb.ClockCycles += uint32(gb.OpCode.ClockCycles)

		gb.Lcd.Update(gb.OpCode.ClockCycles)

	}

}

func (gb *GbCpu) pushToStack(b byte) {
	gb.Reg.SP--
	gb.Mem.WriteByte(gb.Reg.SP, b)
	fmt.Printf("STACK WRITE: %02X @%X\n", b, gb.Reg.SP)
}

func (gb *GbCpu) popFromStack() byte {
	b := gb.Mem.GetByte(gb.Reg.SP)
	fmt.Printf("STACK READ: %02X @%X\n", b, gb.Reg.SP)
	gb.Reg.SP++
	return b
}
