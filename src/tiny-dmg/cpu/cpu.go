package cpu

import (
	"fmt"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
)

type GbCpu struct {
	InterruptsEnabled bool
	Reg               Registers
	mem               *memory.Memory
	lcd               *lcd.Lcd
	ClockCycles       uint32
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

func New(m *memory.Memory, l *lcd.Lcd) (gb *GbCpu, err error) {
	gb = new(GbCpu)
	gb.mem = m
	gb.lcd = l
	return
}

func (gb *GbCpu) PowerOn() {
	gb.Reg.SP = 0xFFFE
	gb.Reg.PC = 0x0100

	gb.Reg.A = 0x01 // We are a 1stgen gameboy
	gb.Reg.F = FlagZ | FlagH | FlagC
	gb.Reg.C = 0x13
	gb.Reg.E = 0xD8
	gb.Reg.H = 0x01
	gb.Reg.L = 0x4D
}

func (gb *GbCpu) pushToStack(b byte) {
	gb.Reg.SP--
	gb.mem.WriteByte(gb.Reg.SP, b)
	fmt.Printf("STACK WRITE: %02X @%X\n", b, gb.Reg.SP)
}

func (gb *GbCpu) popFromStack() byte {
	b := gb.mem.GetByte(gb.Reg.SP)
	fmt.Printf("STACK READ: %02X @%X\n", b, gb.Reg.SP)
	gb.Reg.SP++
	return b
}
