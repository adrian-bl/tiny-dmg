package cpu

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/memory"
)

type GbCpu struct {
	InterruptsEnabled bool
	Reg               Registers
	mem               *memory.Memory
	ClockCycles       uint32
	Halted            bool
}

const (
	FlagZ      = uint8(1 << 7)
	FlagN      = uint8(1 << 6)
	FlagH      = uint8(1 << 5)
	FlagC      = uint8(1 << 4)
	FlagMask   = uint8(FlagZ | FlagN | FlagH | FlagC)
	FlagUnused = uint8(0xF)
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

func New(m *memory.Memory) (gb *GbCpu, err error) {
	gb = new(GbCpu)
	gb.mem = m
	return
}

func (gb *GbCpu) PowerOn() {
	gb.Reg.SP = 0x00
	gb.Reg.PC = 0x00

	gb.Reg.A = 0x00
	gb.Reg.F = 0x00
	gb.Reg.B = 0x00
	gb.Reg.C = 0x00
	gb.Reg.D = 0x00
	gb.Reg.E = 0x00
	gb.Reg.H = 0x00
	gb.Reg.L = 0x00
}

func (gb *GbCpu) Execute(opcode uint8) uint8 {
	oc := OpCodes[opcode]

	oc.Callback(gb)

	// Bit 0..3 are ALWAYS zero, even if the code set them
	gb.Reg.F &^= FlagUnused
	return oc.ClockCycles
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
