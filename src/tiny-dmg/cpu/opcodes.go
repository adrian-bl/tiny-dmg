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
	0x03: {"INC BC", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.B, &gb.Reg.C) }},
	0x04: {"INCb", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.B) }},
	0x05: {"DECb", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.B) }},
	0x06: {"LD B, n    ;", 8, Op_LDBn},
	0x07: {"RLCA", 4, Op_RLCA},
	0x09: {"ADD HL, BC", 4, Op_ADD_HL_BC},
	0x0A: {"LD A, (BC)", 8, Op_LD_A_BC},
	0x0B: {"DECbc", 8, func(gb *GbCpu) { Do_Dec_88(gb, &gb.Reg.B, &gb.Reg.C) }},
	0x0C: {"INCc", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.C) }},
	0x0D: {"DECc", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.C) }},
	0x0E: {"LD C, n    ;", 8, Op_LDCn},
	0x11: {"LDHL", 12, Op_LD_DE_nn},
	0x12: {"LDnA", 8, Op_LD_n_A},
	0x13: {"INC DE", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.D, &gb.Reg.E) }},
	0x14: {"INC D", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.D) }},
	0x15: {"DEC D", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.D) }},
	0x16: {"LD D, n    ;", 8, Op_LDDn},
	0x17: {"RL A", 8, func(gb *GbCpu) { Cb_rl(gb, &gb.Reg.A, gb.Reg.A); gb.Reg.PC++ }},
	0x18: {"JRn", 8, Op_JR_n},
	0x19: {"ADD HL,DE", 8, Op_ADD_HL_DE},
	0x1A: {"LD A, (DE)", 8, Op_LD_A_DE},
	0x1C: {"INC E", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.E) }},
	0x1D: {"DEC E", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.E) }},
	0x1E: {"LDEn", 8, Op_LDEn},
	0x20: {"JPNZ", 12, Op_JPnz}, // Fixme: This can be 12 or 8
	0x21: {"LD HL,d16", 12, Op_LD_HL_nn},
	0x22: {"LD (HL+),A", 8, Op_LDI_HL_A},
	0x23: {"INC HL", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.H, &gb.Reg.L) }},
	0x24: {"INC H", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.H) }},
	0x26: {"LD H,n", 8, Op_LD_H_n},
	0x27: {"DAA", 8, Op_DAA},
	0x28: {"JP Z", 12, Op_JPz},
	0x29: {"ADD HL, HL", 4, Op_ADD_HL_HL},
	0x2A: {"LDA+", 8, Op_LD_A_HLi},
	0x2C: {"INC L", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.L) }},
	0x2D: {"DEC L", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.L) }},
	0x2E: {"LD L,n", 8, Op_LD_L_n},
	0x2F: {"CPL", 4, Op_CPL},
	0x30: {"JR NC", 4, Op_JR_NC_n},
	0x31: {"LDSP", 12, Op_LD_SP_nn},
	0x32: {"LD (HL-),A", 8, Op_LDD_HL_A},
	0x33: {"INC SP", 8, func(gb *GbCpu) { gb.Reg.SP++; gb.Reg.PC++ }},
	0x34: {"INC (HL)", 8, Op_INC_HL},
	0x35: {"DEC (HL)", 8, Op_DEC_HL},
	0x36: {"LD (HL),d8", 12, Op_LD_HL_d8},
	0x38: {"JR C N", 8, Op_JR_C_n},
	0x3A: {"LD A, (HL-)", 8, Op_LD_A_HLdec},
	0x3C: {"INC A", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.A) }},
	0x3D: {"DEC A", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.A) }},
	0x3E: {"LDAn", 8, Op_LDAn},
	0x40: {"LD B B", 4, Op_NOP},
	0x46: {"LD B,(HL)", 4, Op_LD_B_HL},
	0x47: {"LD B,A", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.A; gb.Reg.PC++ }},
	0x4F: {"LD C,A", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.A; gb.Reg.PC++ }},
	0x4E: {"LD C,(HL)", 8, Op_LD_C_HL},
	0x54: {"LD D,H", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.H; gb.Reg.PC++ }},
	0x56: {"LD D,(HL)", 8, Op_LD_D_HL},
	0x57: {"LD D,A", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.A; gb.Reg.PC++ }},
	0x5D: {"LD E,L", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.L; gb.Reg.PC++ }},
	0x5E: {"LD E,(HL)", 8, Op_LD_E_HL},
	0x5F: {"LD E,A", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.A; gb.Reg.PC++ }},
	0x60: {"LD H,B", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.B; gb.Reg.PC++ }},
	0x61: {"LD H,C", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.C; gb.Reg.PC++ }},
	0x62: {"LD H,D", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.D; gb.Reg.PC++ }},
	0x63: {"LD H,E", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.E; gb.Reg.PC++ }},
	0x66: {"LD H,(HL)", 8, Op_LD_H_HL},
	0x67: {"LD H,A", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.A; gb.Reg.PC++ }},
	0x68: {"LD L,B", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.B; gb.Reg.PC++ }},
	0x69: {"LD L,C", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.C; gb.Reg.PC++ }},
	0x6A: {"LD L,D", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.D; gb.Reg.PC++ }},
	0x6B: {"LD L,E", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.E; gb.Reg.PC++ }},
	0x6F: {"LD L,A", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.A; gb.Reg.PC++ }},
	0x70: {"LD (HL),B", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.B) }},
	0x71: {"LD (HL),C", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.C) }},
	0x72: {"LD (HL),D", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.D) }},
	0x73: {"LD (HL),E", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.E) }},
	0x76: {"HALT", 4, func(gb *GbCpu) { fmt.Printf("--> HALT!"); gb.Reg.PC++ }},
	0x77: {"LD (HL),A", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.A) }},
	0x78: {"LD A,B", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.B; gb.Reg.PC++ }},
	0x79: {"LD A,C", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.C; gb.Reg.PC++ }},
	0x7A: {"LD A,D", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.D; gb.Reg.PC++ }},
	0x7B: {"LD A,E", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.E; gb.Reg.PC++ }},
	0x7C: {"LD A,H", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.H; gb.Reg.PC++ }},
	0x7D: {"LD A,L", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.L; gb.Reg.PC++ }},
	0x7E: {"LD A,(HL)", 8, Op_LD_A_HL},
	0x80: {"ADD A,B", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x81: {"ADD A,C", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x82: {"ADD A,D", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x83: {"ADD A,E", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x84: {"ADD A,H", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x85: {"ADD A,L", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x87: {"ADD A,A", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0x88: {"ADC B", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x89: {"ADC C", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x8A: {"ADC D", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x8B: {"ADC E", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x8C: {"ADC H", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x8D: {"ADC L", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x90: {"SUB A,B", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x91: {"SUB A,C", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x92: {"SUB A,D", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x93: {"SUB A,E", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x94: {"SUB A,H", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x95: {"SUB A,L", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x97: {"SUB A,A", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xA0: {"AND B      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xA1: {"AND C      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xA2: {"AND D      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0xA3: {"AND E      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0xA4: {"AND H      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0xA5: {"AND L      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0xA7: {"AND A      ;", 8, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xA8: {"XOR B      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xA9: {"XOR C      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xAA: {"XOR D      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0xAB: {"XOR E      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0xAC: {"XOR H      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0xAD: {"XOR L      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.L) }},
	// 0xae
	0xAF: {"XOR A      ;", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xB0: {"ORAB", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xB1: {"ORAC", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xB8: {"CP B", 4, func(gb *GbCpu) { Do_Cp(gb, gb.Reg.A, gb.Reg.B) }},
	0xC0: {"RET nz", 4, Op_RET_NZ},
	0xC1: {"POP BC", 16, Op_POP_BC},
	0xC2: {"JP NZ", 16, Op_JP_NZ},
	0xC3: {"JP  ", 16, Op_JP},
	0xC5: {"PUSH BC", 16, Op_PUSH_BC},
	0xC6: {"ADD A n", 8, Op_ADD_A_n},
	0xC8: {"RET Z", 20, Op_RET_Z},
	0xC9: {"RET ", 16, Op_RET},
	0xCA: {"JP Z NN", 16, Op_JP_Z_NN},
	0xCD: {"CALn", 24, Op_CALL},
	0xCB: {"CB! ", 12, Cb_Disp}, // fixme: cb takes 4 cycles + the code executed (mostly 8)
	0xCF: {"RST8", 8, Op_Rst8},
	0xD0: {"RENC", 20, Op_RetNC}, // 20 or 8 ?!
	0xD1: {"POP DE", 16, Op_POP_DE},
	0xD5: {"PUSH DE", 16, Op_PUSH_DE},
	0xDF: {"RST18", 8, Op_Rst18},
	0xE0: {"LDHn", 12, Op_LDHnA},
	0xE1: {"POP HL", 8, Op_POP_HL},
	0xE2: {"LD (C),A", 8, Op_LD_C_A},
	0xE5: {"PUSH HL", 16, Op_PUSH_HL},
	0xE6: {"ANDa", 8, Op_ANDAn},
	0xE8: {"ADD SP n", 8, Op_ADD_SP_n},
	0xE9: {"JP HL", 8, Op_JP_HL},
	0xEA: {"LD (a16),A", 16, Op_LD_a16_A},
	0xEF: {"RST28", 8, Op_Rst28},
	0xF0: {"LDHA", 12, Op_LDHAn}, //
	0xF1: {"POP AF", 12, Op_POP_AF},
	0xF3: {"DI  ", 4, Op_DI},
	0xF5: {"PUSH AF", 16, Op_PUSH_AF},
	0xF6: {"OR n", 8, Op_OR_n},
	0xFA: {"LD A, (a16)", 16, Op_LD_A_a16},
	0xFB: {"EI  ", 4, Op_EI},
	0xFE: {"CP d8", 8, Op_CPd8},
}

func (gb *GbCpu) crash() {
	fmt.Printf(">>> crashing at sp=%X, pc=%X, hl=%02X%02X\n", gb.Reg.SP, gb.Reg.PC, gb.Reg.H, gb.Reg.L)
	for {
	}
}

func Op_RET(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
}

func Op_RET_Z(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
	} else {
		gb.Reg.PC++
	}
}

func Op_RET_NZ(gb *GbCpu) {
	if gb.Reg.F&FlagZ == 0 {
		gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
	} else {
		gb.Reg.PC++
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
	} else {
		gb.Reg.PC += 3
	}
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

func Op_DEC_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.Mem.GetByte(addr)
	Do_Dec_Uint8(gb, &val)
	gb.Mem.WriteByte(addr, val)
}

func Op_INC_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.Mem.GetByte(addr)
	Do_Inc_Uint8(gb, &val)
	gb.Mem.WriteByte(addr, val)
}

// 0xe6 AND A, n
func Op_ANDAn(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	Do_And_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
}

func Op_LDHnA(gb *GbCpu) {
	dst := uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + 0xFF00

	old := gb.Mem.GetByte(dst)
	gb.Mem.WriteByte(dst, gb.Reg.A)
	gb.Reg.PC += 2
	fmt.Printf("WROTE %04X to %04X, it was %04X\n", gb.Mem.GetByte(dst), dst, old)
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

func Op_LDEn(gb *GbCpu) {
	gb.Reg.E = gb.Mem.GetByte(gb.Reg.PC + 1)
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

func Op_LD_L_n(gb *GbCpu) {
	gb.Reg.L = gb.Mem.GetByte(gb.Reg.PC + 1)
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

func Op_LD_B_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.B = gb.Mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_C_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.C = gb.Mem.GetByte(addr)
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

func Op_LD_H_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.H = gb.Mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_BC(gb *GbCpu) {
	addr := uint16(gb.Reg.B)<<8 + uint16(gb.Reg.C)
	gb.Reg.A = gb.Mem.GetByte(addr)
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

// Store value specified by HL in A, decrement HL
func Op_LD_A_HLdec(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.Mem.GetByte(val)
	val--
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

func Op_LD_HL_x(gb *GbCpu, value uint8) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Mem.WriteByte(addr, value)
	gb.Reg.PC++
}

func Op_ADD_HL_DE(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	de := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	Do_Add_1616(gb, &hl, de)

	gb.Reg.L = uint8(hl & 0xFF)
	gb.Reg.H = uint8((hl >> 8 & 0xFF))
}

func Op_ADD_HL_BC(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	bc := uint16(gb.Reg.B)<<8 + uint16(gb.Reg.C)
	Do_Add_1616(gb, &hl, bc)

	gb.Reg.L = uint8(hl & 0xFF)
	gb.Reg.H = uint8((hl >> 8 & 0xFF))
}

func Op_ADD_HL_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Add_1616(gb, &hl, hl)

	gb.Reg.L = uint8(hl & 0xFF)
	gb.Reg.H = uint8((hl >> 8 & 0xFF))
}

func Op_ADD_SP_n(gb *GbCpu) {
	gb.Reg.PC++
	val := uint16(int8(gb.Mem.GetByte(gb.Reg.PC)))
	Do_Add_1616(gb, &gb.Reg.SP, val)

	// unlike raw add_1616, this does always clear
	// the zero flag
	gb.Reg.F &= ^FlagZ
}

func Op_ADD_A_n(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	Do_Add_88(gb, &gb.Reg.A, val)
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

func Op_JP_NZ(gb *GbCpu) {
	if (gb.Reg.F & FlagZ) != 0 {
		gb.Reg.PC += 3
	} else {
		gb.Reg.PC = uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
	}
}

func Op_JP(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.Mem.GetByte(gb.Reg.PC+1)) + uint16(gb.Mem.GetByte(gb.Reg.PC+2))<<8
}

func Op_JR_n(gb *GbCpu) {
	add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
	gb.Reg.PC += 2 + uint16(add)
}

func Op_JR_C_n(gb *GbCpu) {
	if gb.Reg.F&FlagC != 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JR_NC_n(gb *GbCpu) {
	if gb.Reg.F&FlagC == 0 {
		add := int8(gb.Mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	} else {
		gb.crash()
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

func Op_OR_n(gb *GbCpu) {
	val := gb.Mem.GetByte(gb.Reg.PC + 1)
	Do_Or_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++
}

func Op_Rst8(gb *GbCpu) {
	gb.Reg.PC++
	gb.pushToStack(uint8(gb.Reg.PC >> 8 & 0xFF))
	gb.pushToStack(uint8(gb.Reg.PC & 0xFF))
	gb.Reg.PC = 0x8
}

func Op_Rst18(gb *GbCpu) {
	gb.Reg.PC++
	gb.pushToStack(uint8(gb.Reg.PC >> 8 & 0xFF))
	gb.pushToStack(uint8(gb.Reg.PC & 0xFF))
	gb.Reg.PC = 0x18
}

func Op_Rst28(gb *GbCpu) {
	gb.Reg.PC++
	gb.pushToStack(uint8(gb.Reg.PC >> 8 & 0xFF))
	gb.pushToStack(uint8(gb.Reg.PC & 0xFF))
	gb.Reg.PC = 0x28
}

// Stolen from Cinoop
func Op_DAA(gb *GbCpu) {
	s := uint16(gb.Reg.A)

	if gb.Reg.F&FlagN != 0 {
		if gb.Reg.F&FlagH != 0 {
			s = (s - 0x06) & 0xFF
		}
		if gb.Reg.F&FlagH != 0 {
			s -= 0x60
		}
	} else {
		if gb.Reg.F&FlagH != 0 || (s&0xF) > 9 {
			s += 0x06
		}
		if gb.Reg.F&FlagH != 0 || s > 0x9F {
			s += 0x60
		}
	}

	gb.Reg.A = uint8(s)
	gb.Reg.F &= ^(FlagH | FlagZ)

	if gb.Reg.A == 0 {
		gb.Reg.F |= FlagZ
	}
	if s >= 0x100 {
		gb.Reg.F |= FlagC
	}

	gb.Reg.PC++
}
