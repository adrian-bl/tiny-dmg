package cpu

import (
	"fmt"
)

type OpEntry struct {
	Name   string
	Cycles uint8
	Cback  func(*GbCpu)
}

var OpCodes = map[uint8]OpEntry{
	0x00: {"NOOP", 4, Op_NOP},
	0x01: {"LDHL", 12, Op_LD_BC_nn},
	0x04: {"INCb", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.B) }},
	0x05: {"DECb", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.B) }},
	0x06: {"LD B, n    ;", 8, Op_LDBn},
	0x07: {"RLCA", 4, Op_RLCA},
	0x0B: {"DECbc", 8, func(gb *GbCpu) { Do_Dec_88(gb, &gb.Reg.B, &gb.Reg.C) }},
	0x0C: {"INCc", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.C) }},
	0x0D: {"DECc", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.C) }},
	0x0E: {"LD C, n    ;", 8, Op_LDCn},
	0x11: {"LDHL", 12, Op_LD_DE_nn},
	0x12: {"LDnA", 8, Op_LD_n_A},
	0x13: {"INC3", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.D, &gb.Reg.E) }},
	0x18: {"JRn", 8, Op_JR_n},
	0x20: {"JPNZ", 12, Op_JPnz}, // Fixme: This can be 12 or 8
	0x21: {"LDHL", 12, Op_LD_HL_nn},
	0x22: {"LDIx", 12, Op_LDI_HL_A},
	0x2A: {"LDA+", 8, Op_LD_A_HLi},
	0x31: {"LDSP", 12, Op_LD_SP_nn},
	0x3E: {"LDAn", 8, Op_LDAn},
	0x78: {"LDAB", 4, Op_LDAB},
	0x7E: {"LDAHL", 8, Op_LD_A_HL},
	0xAF: {"XOR A      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xB1: {"ORAC", 4, Op_OrAC},
	0xC3: {"JP  ", 16, Op_JP},
	0xC9: {"RET ", 16, Op_RET},
	0xCD: {"CALn", 24, Op_CALL},
	0xCB: {"CB! ", 12, Cb_Disp},  // fixme: cb takes 4 cycles + the code executed (mostly 8)
	0xD0: {"RENC", 20, Op_RetNC}, // 20 or 8 ?!
	0xE0: {"LDHn", 12, Op_LDHnA},
	0xE6: {"ANDa", 8, Op_ANDAn},
	0xF0: {"LDHA", 12, Op_LDHAn}, //
	0xF1: {"POP!", 12, Op_POP_AF},
	0xF3: {"DI  ", 4, Op_DI},
	0xF5: {"PSaf", 16, Op_PUSH_AF},
	0xFB: {"EI  ", 4, Op_EI},
	0xFE: {"CPd8", 8, Op_CPd8},
}

func Cb_Disp(gb *GbCpu) {
	op := gb.Mem.GetByte(gb.Reg.PC + 1)
	switch op {
	case 0xBF:
		Cb_ResetBit(0x07, &gb.Reg.A)
	default:
		panic(nil)
	}
	gb.Reg.PC += 2
}

func Op_RET(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.popFromStack())<<8 + uint16(gb.popFromStack())
}

func Op_JPnz(gb *GbCpu) {
	if gb.Reg.F&FlagZ == 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_CPd8(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)

	gb.Reg.F &= ^FlagMask // clear all bits
	gb.Reg.F |= FlagN     // this is always set

	if gb.Reg.A == val {
		gb.Reg.F |= FlagZ
	}

	fixmeSlowPrefixed := int8(gb.Reg.A&0xF) - int8(val&0xF)
	if fixmeSlowPrefixed < 0 {
		// (gb.Reg.A&0xF < val&0xF) ??
		gb.Reg.F |= FlagH
		panic(nil)
	}
	if gb.Reg.A < val {
		gb.Reg.F |= FlagC
	}

	gb.Reg.PC += 2
}

func Op_RetNC(gb *GbCpu) {
	if gb.Reg.F&FlagC == 0 {
		panic(nil) // not implemented yet
	}
	gb.Reg.PC++
}

func Op_RLCA(gb *GbCpu) {
	gb.Reg.PC++

	// Clear all bits and set carry flag if needed
	gb.Reg.F &= ^FlagMask
	if gb.Reg.A&0x80 != 0 {
		gb.Reg.F |= FlagC
	}

	gb.Reg.A = gb.Reg.A>>7 | gb.Reg.A<<1
	// gnuboy doesn't do this (?)
	if gb.Reg.A == 0 {
		gb.Reg.A |= FlagZ
		panic(nil)
	}
}

func Op_LDAB(gb *GbCpu) {
	gb.Reg.A = gb.Reg.B
	gb.Reg.PC++
}

func Op_LDHAn(gb *GbCpu) {
	src := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + 0xFF00
	gb.Reg.A = gb.Mem.GetByte(src)
	fmt.Printf("READ %04X from %04X\n", gb.Reg.A, src)
	gb.Reg.PC += 2
}

func Op_CALL(gb *GbCpu) {
	spc := gb.Reg.PC + 3
	gb.pushToStack(uint8(spc & 0xFF))
	gb.pushToStack(uint8(spc >> 8 & 0xFF))
	Op_JP(gb)
}

func Op_PUSH_AF(gb *GbCpu) {
	gb.pushToStack(gb.Reg.A)
	gb.pushToStack(gb.Reg.F)
	gb.Reg.PC++
}

func Op_POP_AF(gb *GbCpu) {
	// fixme: is the order correct?
	gb.Reg.A = gb.popFromStack()
	gb.Reg.F = gb.popFromStack()
	gb.Reg.PC++
}

func Op_OrAC(gb *GbCpu) {
	gb.Reg.A |= gb.Reg.C
	gb.Reg.PC++

	gb.Reg.F &= ^FlagMask // clear all bits
	if gb.Reg.A == 0 {
		gb.Reg.F |= FlagZ
	}
}

// 0xe6 AND A, n
func Op_ANDAn(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	Do_And_88(gb, &gb.Reg.A, val)
	gb.Reg.PC += 2
}

func Op_LDHnA(gb *GbCpu) {
	dst := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + 0xFF00

	gb.Mem.WriteByte(dst, gb.Reg.A)
	gb.Reg.PC += 2
	fmt.Printf("WROTE %04X to %04X\n", gb.Mem.GetByte(dst), dst)
}

func Op_LDAn(gb *GbCpu) {
	gb.Reg.A = gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDBn(gb *GbCpu) {
	gb.Reg.B = gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDCn(gb *GbCpu) {
	gb.Reg.C = gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

// Put value of A into location specified by DE
func Op_LD_n_A(gb *GbCpu) {
	addr := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	gb.Mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_HL_nn(gb *GbCpu) {
	Do_Load_88(gb, &gb.Reg.L, &gb.Reg.H)
}
func Op_LD_DE_nn(gb *GbCpu) {
	Do_Load_88(gb, &gb.Reg.E, &gb.Reg.D)
}
func Op_LD_BC_nn(gb *GbCpu) {
	Do_Load_88(gb, &gb.Reg.C, &gb.Reg.B)
}

func Op_LD_SP_nn(gb *GbCpu) {
	gb.Reg.SP = uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
	gb.Reg.PC += 3
}

func Op_LD_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.Mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_HLi(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.Mem.GetByte(val)
	val++
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

// Put value of A into location specified by HL, increment HL
func Op_LDI_HL_A(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Mem.WriteByte(val, gb.Reg.A)
	val++
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

func Op_DI(gb *GbCpu) {
	gb.InterruptsEnabled = false
	gb.Reg.PC++
	fmt.Printf(">>> Code disabled interrupts (Fixme: not handled yet)\n")
}

func Op_EI(gb *GbCpu) {
	gb.InterruptsEnabled = true
	gb.Reg.PC++
	fmt.Printf(">>> Code enabled interrupts\n")
}

func Op_JP(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
}

func Op_JR_n(gb *GbCpu) {

	//if (gb.Reg.PC == 0x09A3) {
	fmt.Printf("FETCH: %d", gb.Mem.GetByte(gb.Reg.PC+1))
	gb.Reg.PC += 2 + uint16(gb.Mem.GetByte(gb.Reg.PC+1))
	fmt.Printf(" -> JUMP TO %X\n", gb.Reg.PC)
	if gb.Reg.PC != 0x84 && gb.Reg.PC != 0x9A && gb.Reg.PC != 0xB0 {
		panic(nil)
		/*
		   func (c *CPU) relativeJump(dist uint8) {
		   	c.PC = signedAdd(c.PC, dist)
		   }

		   func signedAdd(a uint16, b uint8) uint16 {
		   	bSigned := int8(b)
		   	if bSigned >= 0 {
		   		return a + uint16(bSigned)
		   	} else {
		   		return a - uint16(-bSigned)
		   	}
		   }*/
	}
}

func Op_NOP(gb *GbCpu) {
	gb.Reg.PC++
}
