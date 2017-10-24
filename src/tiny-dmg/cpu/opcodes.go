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
	0x01: {"LD BC,d16", 12, Op_LD_BC_nn},
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
	0x16: {"LD D, n    ;", 8, Op_LDDn},
	0x18: {"JRn", 8, Op_JR_n},
	0x19: {"ADD HL,DE", 8, Op_ADD_HL_DE},
	0x1A: {"LD A, (DE)", 8, Op_LD_A_DE},
	0x1C: {"INC E", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.E) }},
	0x20: {"JPNZ", 12, Op_JPnz}, // Fixme: This can be 12 or 8
	0x21: {"LD HL,d16", 12, Op_LD_HL_nn},
	0x22: {"LD (HL+),A", 8, Op_LDI_HL_A},
	0x23: {"INC HL", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.H, &gb.Reg.L) }},
	0x26: {"LD H,n", 8, Op_LD_H_n},
	0x28: {"JP Z", 12, Op_JPz},
	0x2A: {"LDA+", 8, Op_LD_A_HLi},
	0x2F: {"CPL", 4, Op_CPL},
	0x31: {"LDSP", 12, Op_LD_SP_nn},
	0x32: {"LD (HL-),A", 8, Op_LDD_HL_A},
	0x33: {"INC SP", 8, func(gb *GbCpu) { gb.Reg.SP++; gb.Reg.PC++ }},
	0x36: {"LD (HL),d8", 12, Op_LD_HL_d8},
	0x38: {"JR C N", 8, Op_JR_C_n},
	0x3E: {"LDAn", 8, Op_LDAn},
	0x47: {"LD B,A", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.A; gb.Reg.PC++ }},
	0x4F: {"LD C,A", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.A; gb.Reg.PC++ }},
	0x56: {"LD D,(HL)", 8, Op_LD_D_HL},
	0x57: {"LD D,A", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.A; gb.Reg.PC++ }},
	0x5E: {"LD E,(HL)", 8, Op_LD_E_HL},
	0x5F: {"LD E,A", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.A; gb.Reg.PC++ }},
	0x6F: {"LD L,A", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.A; gb.Reg.PC++ }},
	0x77: {"LD (HL),A", 8, Op_LD_HL_A},
	0x78: {"LD A,B", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.B; gb.Reg.PC++ }},
	0x79: {"LD A,C", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.C; gb.Reg.PC++ }},
	0x7A: {"LD A,D", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.D; gb.Reg.PC++ }},
	0x7B: {"LD A,E", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.E; gb.Reg.PC++ }},
	0x7C: {"LD A,H", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.H; gb.Reg.PC++ }},
	0x7D: {"LD A,L", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.L; gb.Reg.PC++ }},
	0x7E: {"LD A,(HL)", 8, Op_LD_A_HL},
	0x87: {"ADD A,A", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xA1: {"AND C      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xA7: {"AND A      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xA9: {"XOR C      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xAF: {"XOR A      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xB0: {"ORAB", 4, func(gb *GbCpu) { Do_Or_8(gb, &gb.Reg.A, gb.Reg.B) }},
	0xB1: {"ORAC", 4, func(gb *GbCpu) { Do_Or_8(gb, &gb.Reg.A, gb.Reg.C) }},
	0xC1: {"POP BC", 16, Op_POP_BC},
	0xC3: {"JP  ", 16, Op_JP},
	0xC5: {"PUSH BC", 16, Op_PUSH_BC},
	0xC8: {"RET Z", 20, Op_RET_Z},
	0xC9: {"RET ", 16, Op_RET},
	0xCA: {"JP Z NN", 16, Op_JP_Z_NN},
	0xCD: {"CALn", 24, Op_CALL},
	0xCB: {"CB! ", 12, Cb_Disp},  // fixme: cb takes 4 cycles + the code executed (mostly 8)
	0xD0: {"RENC", 20, Op_RetNC}, // 20 or 8 ?!
	0xD1: {"POP DE", 16, Op_POP_DE},
	0xD5: {"PUSH DE", 16, Op_PUSH_DE},
	0xE0: {"LDHn", 12, Op_LDHnA},
	0xE1: {"POP HL", 8, Op_POP_HL},
	0xE2: {"LD (C),A", 8, Op_LD_C_A},
	0xE5: {"PUSH HL", 16, Op_PUSH_HL},
	0xE6: {"ANDa", 8, Op_ANDAn},
	0xE9: {"JP HL", 8, Op_JP_HL},
	0xEA: {"LD (a16),A", 16, Op_LD_a16_A},
	0xEF: {"RST28", 8, Op_Rst28},
	0xF0: {"LDHA", 12, Op_LDHAn}, //
	0xF1: {"POP AF", 12, Op_POP_AF},
	0xF3: {"DI  ", 4, Op_DI},
	0xF5: {"PUSH AF", 16, Op_PUSH_AF},
	0xFA: {"LD A, (a16)", 16, Op_LD_A_a16},
	0xFB: {"EI  ", 4, Op_EI},
	0xFE: {"CP d8", 8, Op_CPd8},
}

func (gb *GbCpu) crash() {
	fmt.Printf(">>> crashing at pc=%X\n", gb.Reg.PC)
	panic(nil)
}

func Cb_Disp(gb *GbCpu) {
	op := gb.Mem.GetByte(gb.Reg.PC + 1)
	fmt.Printf("-> CB %02X\n", op)

	switch op {
	case 0xBF:
		Cb_ResetBit(0x07, &gb.Reg.A)
	case 0x37:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.A)
	case 0x87:
		Cb_RES(0, &gb.Reg.A)
	default:
		panic(fmt.Errorf("Unknown cb opcode: %02X", op))
	}
	gb.Reg.PC += 2
}

func Op_RET(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
}

func Op_RET_Z(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
	} else {
		// inc pc++?
		gb.crash()
	}
}

func Op_JP_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.PC = hl
}

func Op_JPnz(gb *GbCpu) {
	if gb.Reg.F&FlagZ == 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JPz(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JP_Z_NN(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		addr := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
		gb.Reg.PC = addr
		gb.crash()
	}
	gb.Reg.PC += 3
}

func Op_CPd8(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)

	gb.Reg.F &= ^FlagMask // clear all bits
	gb.Reg.F |= FlagN     // this is always set

	if gb.Reg.A == val {
		gb.Reg.F |= FlagZ
	}

	if gb.Reg.A < val {
		gb.Reg.F |= FlagC
	}

	if (gb.Reg.A & 0x0F) < (val & 0x0F) {
		gb.Reg.F |= FlagH
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

func Op_LDHAn(gb *GbCpu) {
	src := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + 0xFF00
	gb.Reg.A = gb.Mem.GetByte(src)
	fmt.Printf("READ %04X from %04X\n", gb.Reg.A, src)
	gb.Reg.PC += 2
}

func Op_CALL(gb *GbCpu) {
	spc := gb.Reg.PC + 3
	gb.pushToStack(uint8(spc >> 8 & 0xFF))
	gb.pushToStack(uint8(spc & 0xFF))
	Op_JP(gb)
}

func Op_PUSH_AF(gb *GbCpu) {
	gb.pushToStack(gb.Reg.A)
	gb.pushToStack(gb.Reg.F)
	gb.Reg.PC++
}

func Op_PUSH_BC(gb *GbCpu) {
	gb.pushToStack(gb.Reg.B)
	gb.pushToStack(gb.Reg.C)
	gb.Reg.PC++
}

func Op_PUSH_DE(gb *GbCpu) {
	gb.pushToStack(gb.Reg.D)
	gb.pushToStack(gb.Reg.E)
	gb.Reg.PC++
}

func Op_PUSH_HL(gb *GbCpu) {
	gb.pushToStack(gb.Reg.H)
	gb.pushToStack(gb.Reg.L)
	gb.Reg.PC++
}

func Op_POP_AF(gb *GbCpu) {
	gb.Reg.F = gb.popFromStack()
	gb.Reg.A = gb.popFromStack()
	gb.Reg.PC++
}

func Op_POP_BC(gb *GbCpu) {
	gb.Reg.C = gb.popFromStack()
	gb.Reg.B = gb.popFromStack()
	gb.Reg.PC++
}

func Op_POP_HL(gb *GbCpu) {
	gb.Reg.L = gb.popFromStack()
	gb.Reg.H = gb.popFromStack()
	gb.Reg.PC++
}

func Op_POP_DE(gb *GbCpu) {
	gb.Reg.E = gb.popFromStack()
	gb.Reg.D = gb.popFromStack()
	gb.Reg.PC++
}

// 0xe6 AND A, n
func Op_ANDAn(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	Do_And_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
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

func Op_LDDn(gb *GbCpu) {
	gb.Reg.D = gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

// Put value of A into location specified by DE
func Op_LD_n_A(gb *GbCpu) {
	addr := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	gb.Mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_HL_d8(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Mem.WriteByte(addr, val)
	gb.Reg.PC += 2
}

func Op_LD_C_A(gb *GbCpu) {
	gb.Mem.WriteByte(0xFF00+uint16(gb.Reg.C), gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_A_a16(gb *GbCpu) {
	addr := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
	gb.Reg.A = gb.Mem.GetByte(addr)
	gb.Reg.PC += 3
}

func Op_LD_a16_A(gb *GbCpu) {
	addr := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
	gb.Mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC += 3
	fmt.Printf("LD %X -> %X\n", addr, gb.Reg.A)
}

func Op_LD_H_n(gb *GbCpu) {
	gb.Reg.H = gb.Mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
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

func Op_LD_D_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.D = gb.Mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_E_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.E = gb.Mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_DE(gb *GbCpu) {
	addr := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
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

// Put value of A into location specified by HL, decrement HL
func Op_LDD_HL_A(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Mem.WriteByte(val, gb.Reg.A)
	val--
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

func Op_LD_HL_A(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Mem.WriteByte(val, gb.Reg.A)
	gb.Reg.PC++
}

func Op_ADD_HL_DE(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	de := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	Do_Add_1616(gb, &hl, de)

	gb.Reg.L = uint8(hl & 0xFF)
	gb.Reg.H = uint8((hl >> 8 & 0xFF))
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
	add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
	gb.Reg.PC += 2 + uint16(add)
}

func Op_JR_C_n(gb *GbCpu) {
	if gb.Reg.F&FlagH != 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_CPL(gb *GbCpu) {
	gb.Reg.F |= (FlagN | FlagH) // N and H are always set
	gb.Reg.A = ^gb.Reg.A        // complement A
	gb.Reg.PC++
}

func Op_NOP(gb *GbCpu) {
	gb.Reg.PC++
}

func Op_Rst28(gb *GbCpu) {
	gb.Reg.PC++
	gb.pushToStack(uint8(gb.Reg.PC >> 8 & 0xFF))
	gb.pushToStack(uint8(gb.Reg.PC & 0xFF))
	gb.Reg.PC = 0x28
}
