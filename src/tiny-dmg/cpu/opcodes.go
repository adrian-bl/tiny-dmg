package cpu

import (
	"fmt"
)

type OpEntry struct {
	Name        string
	ClockCycles uint8
	Callback    func(*GbCpu)
}

var OpCodes = map[uint8]OpEntry{
	0x00: {"NOOP			", 4, Op_NOP},
	0x01: {"LD BC, d16		", 12, Op_LD_BC_nn},
	0x02: {"LD (BC), A		", 8, Op_LD_BC_A},
	0x03: {"INC BC			", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.B, &gb.Reg.C) }},
	0x04: {"INC B			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.B) }},
	0x05: {"DEC B			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.B) }},
	0x06: {"LD B, d8		", 8, Op_LDBn},
	0x07: {"RLCA			", 4, Op_RLCA},
	0x08: {"LD (a16), SP	", 20, Op_LD_a16_SP},
	0x09: {"ADD HL, BC		", 8, Op_ADD_HL_BC},
	0x10: {"STOP 0			", 4, func(gb *GbCpu) { fmt.Printf("!! stop ignored\n"); gb.Reg.PC++ }},
	0x0A: {"LD A, (BC)		", 8, Op_LD_A_BC},
	0x0B: {"DEC BC			", 8, func(gb *GbCpu) { Do_Dec_88(gb, &gb.Reg.B, &gb.Reg.C) }},
	0x0C: {"INC C			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.C) }},
	0x0D: {"DEC C			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.C) }},
	0x0E: {"LD C, d8		", 8, Op_LDCn},
	0x0F: {"RRC A			", 4, func(gb *GbCpu) { Do_Rrc(gb, &gb.Reg.A) }},
	// 0x10 STOP 0, 4
	0x11: {"LD DE, d16		", 12, Op_LD_DE_nn},
	0x12: {"LD (DE), A		", 8, Op_LD_DE_A},
	0x13: {"INC DE			", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.D, &gb.Reg.E) }},
	0x14: {"INC D			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.D) }},
	0x15: {"DEC D			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.D) }},
	0x16: {"LD D, d8		", 8, Op_LDDn},
	0x17: {"RLA				", 4, func(gb *GbCpu) { Cb_rl(gb, &gb.Reg.A, gb.Reg.A); gb.Reg.PC++ }},
	0x18: {"JR r8			", 12, Op_JR_n},
	0x19: {"ADD HL, DE		", 8, Op_ADD_HL_DE},
	0x1A: {"LD A, (DE)		", 8, Op_LD_A_DE},
	0x1B: {"DEC DE			", 8, func(gb *GbCpu) { Do_Dec_88(gb, &gb.Reg.D, &gb.Reg.E) }},
	0x1C: {"INC E			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.E) }},
	0x1D: {"DEC E			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.E) }},
	0x1E: {"LD E, d8		", 8, Op_LDEn},
	0x1F: {"RRA				", 4, Op_RRA},
	0x20: {"JR NZ, r8		", 8, Op_JPnz}, // Fixme: This can be 12 or 8
	0x21: {"LD HL, d16		", 12, Op_LD_HL_nn},
	0x22: {"LD (HL+), A		", 8, Op_LDI_HL_A},
	0x23: {"INC HL			", 8, func(gb *GbCpu) { Do_Inc_88(gb, &gb.Reg.H, &gb.Reg.L) }},
	0x24: {"INC H			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.H) }},
	0x25: {"DEC H			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.H) }},
	0x26: {"LD H, d8		", 8, Op_LD_H_n},
	0x27: {"DAA				", 4, Op_DAA},
	0x28: {"JR Z,r8			", 8, Op_JPz}, // Fixme: This can be 12 or 8
	0x29: {"ADD HL, HL		", 8, Op_ADD_HL_HL},
	0x2A: {"LD A, (HL+)		", 8, Op_LD_A_HLi},
	0x2B: {"DEC HL			", 8, func(gb *GbCpu) { Do_Dec_88(gb, &gb.Reg.H, &gb.Reg.L) }},
	0x2C: {"INC L			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.L) }},
	0x2D: {"DEC L			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.L) }},
	0x2E: {"LD L, d8		", 8, Op_LD_L_n},
	0x2F: {"CPL				", 4, Op_CPL},
	0x30: {"JR NC, r8		", 8, Op_JR_NC_r8}, // Fixme: this can be 12 or 8
	0x31: {"LD SP, d16		", 12, Op_LD_SP_nn},
	0x32: {"LD (HL-), A		", 8, Op_LDD_HL_A},
	0x33: {"INC SP			", 8, func(gb *GbCpu) { gb.Reg.SP++; gb.Reg.PC++ }},
	0x34: {"INC (HL)		", 12, Op_INC_HL},
	0x35: {"DEC (HL)		", 12, Op_DEC_HL},
	0x36: {"LD (HL), d8		", 12, Op_LD_HL_d8},
	// 0x37 SCF 4
	0x38: {"JR C, r8		", 8, Op_JR_C_n}, // Fixme: this can be 12 or 8
	0x39: {"ADD HL, SP		", 8, Op_ADD_HL_SP},
	0x3A: {"LD A, (HL-)		", 8, Op_LD_A_HLdec},
	0x3B: {"DEC SP			", 8, func(gb *GbCpu) { gb.Reg.SP--; gb.Reg.PC++ }},
	0x3C: {"INC A			", 4, func(gb *GbCpu) { Do_Inc_Uint8(gb, &gb.Reg.A) }},
	0x3D: {"DEC A			", 4, func(gb *GbCpu) { Do_Dec_Uint8(gb, &gb.Reg.A) }},
	0x3E: {"LD A, d8		", 8, Op_LDAn},
	0x3F: {"CCF				", 4, Op_CCF},
	0x40: {"LD B, B			", 4, Op_NOP},
	0x41: {"LD B, C			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.C; gb.Reg.PC++ }},
	0x42: {"LD B, D			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.D; gb.Reg.PC++ }},
	0x43: {"LD B, E			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.E; gb.Reg.PC++ }},
	0x44: {"LD B, H			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.H; gb.Reg.PC++ }},
	0x45: {"LD B, L			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.L; gb.Reg.PC++ }},
	0x46: {"LD B, (HL)		", 8, Op_LD_B_HL},
	0x47: {"LD B, A			", 4, func(gb *GbCpu) { gb.Reg.B = gb.Reg.A; gb.Reg.PC++ }},
	0x48: {"LD C, B			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.B; gb.Reg.PC++ }},
	//
	0x4A: {"LD C, D			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.D; gb.Reg.PC++ }},
	0x4B: {"LD C, E			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.E; gb.Reg.PC++ }},
	0x4C: {"LD C, H			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.H; gb.Reg.PC++ }},
	0x4D: {"LD C, L			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.L; gb.Reg.PC++ }},
	0x4F: {"LD C, A			", 4, func(gb *GbCpu) { gb.Reg.C = gb.Reg.A; gb.Reg.PC++ }},
	0x4E: {"LD C, (HL)		", 8, Op_LD_C_HL},
	0x50: {"LD D, B			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.B; gb.Reg.PC++ }},
	0x51: {"LD D, C			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.C; gb.Reg.PC++ }},
	0x53: {"LD D, E			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.E; gb.Reg.PC++ }},
	0x54: {"LD D, H			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.H; gb.Reg.PC++ }},
	0x55: {"LD D, L			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.L; gb.Reg.PC++ }},
	0x56: {"LD D, (HL)		", 8, Op_LD_D_HL},
	0x57: {"LD D, A			", 4, func(gb *GbCpu) { gb.Reg.D = gb.Reg.A; gb.Reg.PC++ }},
	0x58: {"LD E, B			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.B; gb.Reg.PC++ }},
	0x59: {"LD E, C			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.C; gb.Reg.PC++ }},
	0x5A: {"LD E, D			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.D; gb.Reg.PC++ }},
	//
	0x5C: {"LD E, H			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.H; gb.Reg.PC++ }},
	0x5D: {"LD E, L			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.L; gb.Reg.PC++ }},
	0x5E: {"LD E, (HL)		", 8, Op_LD_E_HL},
	0x5F: {"LD E, A			", 4, func(gb *GbCpu) { gb.Reg.E = gb.Reg.A; gb.Reg.PC++ }},
	0x60: {"LD H, B			", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.B; gb.Reg.PC++ }},
	0x61: {"LD H, C			", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.C; gb.Reg.PC++ }},
	0x62: {"LD H, D			", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.D; gb.Reg.PC++ }},
	0x63: {"LD H, E			", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.E; gb.Reg.PC++ }},
	0x66: {"LD H, (HL)		", 8, Op_LD_H_HL},
	0x67: {"LD H, A			", 4, func(gb *GbCpu) { gb.Reg.H = gb.Reg.A; gb.Reg.PC++ }},
	0x68: {"LD L, B			", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.B; gb.Reg.PC++ }},
	0x69: {"LD L, C			", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.C; gb.Reg.PC++ }},
	0x6A: {"LD L, D			", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.D; gb.Reg.PC++ }},
	0x6B: {"LD L, E			", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.E; gb.Reg.PC++ }},
	0x6E: {"LD L,(HL)		", 8, Op_LD_L_HL},
	0x6F: {"LD L,A			", 4, func(gb *GbCpu) { gb.Reg.L = gb.Reg.A; gb.Reg.PC++ }},
	0x70: {"LD (HL), B		", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.B) }},
	0x71: {"LD (HL), C		", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.C) }},
	0x72: {"LD (HL), D		", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.D) }},
	0x73: {"LD (HL), E		", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.E) }},
	0x76: {"HALT			", 4, func(gb *GbCpu) { fmt.Printf("--> HALT!"); gb.Reg.PC++ }},
	0x77: {"LD (HL), A		", 8, func(gb *GbCpu) { Op_LD_HL_x(gb, gb.Reg.A) }},
	0x78: {"LD A, B			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.B; gb.Reg.PC++ }},
	0x79: {"LD A, C			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.C; gb.Reg.PC++ }},
	0x7A: {"LD A, D			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.D; gb.Reg.PC++ }},
	0x7B: {"LD A, E			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.E; gb.Reg.PC++ }},
	0x7C: {"LD A, H			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.H; gb.Reg.PC++ }},
	0x7D: {"LD A, L			", 4, func(gb *GbCpu) { gb.Reg.A = gb.Reg.L; gb.Reg.PC++ }},
	0x7E: {"LD A, (HL)		", 8, Op_LD_A_HL},
	0x80: {"ADD A, B		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x81: {"ADD A, C		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x82: {"ADD A, D		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x83: {"ADD A, E		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x84: {"ADD A, H		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x85: {"ADD A, L		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x86: {"ADD A,(HL)		", 8, Op_ADD_A_HL},
	0x87: {"ADD A, A		", 4, func(gb *GbCpu) { Do_Add_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0x88: {"ADC B			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x89: {"ADC C			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x8A: {"ADC D			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x8B: {"ADC E			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x8C: {"ADC H			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x8D: {"ADC L			", 4, func(gb *GbCpu) { Do_Adc_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x8E: {"ADC A,(HL)		", 8, Op_ADC_A_HL},
	0x90: {"SUB A, B		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x91: {"SUB A, C		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x92: {"SUB A, D		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x93: {"SUB A, E		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x94: {"SUB A, H		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x95: {"SUB A, L		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x96: {"SUB A,(HL)		", 8, Op_SUB_A_HL},
	0x97: {"SUB A, A		", 4, func(gb *GbCpu) { Do_Sub_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0x98: {"SBC A, B		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0x99: {"SBC A, C		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0x9A: {"SBC A, D		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0x9B: {"SBC A, E		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0x9C: {"SBC A, H		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0x9D: {"SBC A, L		", 4, func(gb *GbCpu) { Do_Sbc_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0x9E: {"SBC A,(HL)		", 8, Op_SBC_A_HL},
	0xA0: {"AND B			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xA1: {"AND C			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xA2: {"AND D			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0xA3: {"AND E			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0xA4: {"AND H			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0xA5: {"AND L			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0xA7: {"AND A			", 4, func(gb *GbCpu) { Do_And_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xA8: {"XOR B			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xA9: {"XOR C			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xAA: {"XOR D			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0xAB: {"XOR E			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0xAC: {"XOR H			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0xAD: {"XOR L			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0xAE: {"XOR (HL)		", 8, Op_XOR_HL},
	0xAF: {"XOR A			", 4, func(gb *GbCpu) { Do_Xor_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xB0: {"OR B			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.B) }},
	0xB1: {"OR C			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.C) }},
	0xB2: {"OR D			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.D) }},
	0xB3: {"OR E			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.E) }},
	0xB4: {"OR H			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.H) }},
	0xB5: {"OR L			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.L) }},
	0xB6: {"OR (HL)			", 8, Op_OR_HL},
	0xB7: {"OR A			", 4, func(gb *GbCpu) { Do_Or_88(gb, &gb.Reg.A, gb.Reg.A) }},
	0xB8: {"CP B			", 4, func(gb *GbCpu) { Do_Cp(gb, gb.Reg.A, gb.Reg.B) }},
	0xB9: {"CP C			", 4, func(gb *GbCpu) { Do_Cp(gb, gb.Reg.A, gb.Reg.C) }},
	0xBA: {"CP D			", 4, func(gb *GbCpu) { Do_Cp(gb, gb.Reg.A, gb.Reg.D) }},
	0xBB: {"CP E			", 4, func(gb *GbCpu) { Do_Cp(gb, gb.Reg.A, gb.Reg.E) }},
	0xBE: {"CP (HL)			", 8, Op_CP_HL},
	0xC0: {"RET NZ			", 8, Op_RET_NZ}, // Fixme: can be 8 or 20
	0xC1: {"POP BC			", 12, Op_POP_BC},
	0xC2: {"JP NZ, a16		", 12, Op_JP_NZ}, // Fixme: can be 12 or 16
	0xC3: {"JP a16			", 16, Op_JP},
	0xC4: {"CALL NZ, a16	", 12, Op_CALL_NZ_a16}, // fixme: 12 or 24
	0xC5: {"PUSH BC			", 16, Op_PUSH_BC},
	0xC6: {"ADD A, d8		", 8, Op_ADD_A_n},
	0xC8: {"RET Z			", 8, Op_RET_Z}, // Fixme: can be 8 or 20
	0xC9: {"RET				", 16, Op_RET},
	0xCA: {"JP Z, a16		", 12, Op_JP_Z_NN}, // Fixme: can be 12 or 16
	0xCB: {"PREFIX CB		", 4, Cb_Disp}, // fixme: cb takes 4 cycles + the code executed (mostly 8)
	0xCC: {"CALL Z, a16		", 12, Op_CALL_Z_a16}, // fixme: can be 12 or 24
	0xCD: {"CALL a16		", 24, Op_CALL},
	0xCE: {"ADC A, d8		", 8, Op_ADC_A_d8},
	0xCF: {"RST 8			", 16, func(gb *GbCpu) { Op_Rst(gb, 0x08) }},
	0xD0: {"RET NC			", 8, Op_RetNC}, // fixme: 20 or 8 ?!
	0xD1: {"POP DE			", 12, Op_POP_DE},
	0xD2: {"JP NC, a16		", 12, Op_JP_NC}, // fixme: can be 12 or 16
	0xD5: {"PUSH DE			", 16, Op_PUSH_DE},
	0xD6: {"SUB d8			", 8, Op_SUB_d8},
	0xD8: {"RET C			", 8, Op_RET_C}, // Fixme: can be 8 or 20
	0xD9: {"RETI			", 16, Op_RETI},
	0xDA: {"JP C,a16		", 12, Op_JP_C_NN}, // Fixme: can be 12 or 16
	0xDE: {"SBC A,d8		", 8, Op_SBC_A_d8},
	0xDF: {"RST	18			", 16, func(gb *GbCpu) { Op_Rst(gb, 0x18) }},
	0xE0: {"LDH (a8), A		", 12, Op_LDHnA},
	0xE1: {"POP HL			", 12, Op_POP_HL},
	0xE2: {"LD (C), A		", 8, Op_LD_C_A},
	0xE5: {"PUSH HL			", 16, Op_PUSH_HL},
	0xE6: {"AND d8			", 8, Op_ANDAn},
	0xE8: {"ADD SP, r8		", 16, Op_ADD_SP_n},
	0xE9: {"JP (HL)			", 4, Op_JP_HL},
	0xEA: {"LD (a16), A		", 16, Op_LD_a16_A},
	0xEE: {"XOR d8			", 8, Op_XOR_d8},
	0xEF: {"RST 28			", 16, func(gb *GbCpu) { Op_Rst(gb, 0x28) }},
	0xF0: {"LD A, (a8)		", 12, Op_LDHAn}, //
	0xF1: {"POP AF			", 12, Op_POP_AF},
	0xF3: {"DI				", 4, Op_DI},
	0xF5: {"PUSH AF			", 16, Op_PUSH_AF},
	0xF6: {"OR d8			", 8, Op_OR_n},
	0xF8: {"LD HL,SP+r8		", 12, Op_LD_HL_SP_r8},
	0xF9: {"LD SP,HL		", 8, Op_LD_SP_HL},
	0xFA: {"LD A, (a16)		", 16, Op_LD_A_a16},
	0xFB: {"EI				", 4, Op_EI},
	0xFE: {"CP d8			", 8, Op_CPd8},
	0xFF: {"RST 38			", 16, func(gb *GbCpu) { Op_Rst(gb, 0x38) }},
}

func (gb *GbCpu) crash() {
	fmt.Printf(">>> crashing at sp=%X, pc=%X, hl=%02X%02X, a=%X, f=%X\n", gb.Reg.SP, gb.Reg.PC, gb.Reg.H, gb.Reg.L, gb.Reg.A, gb.Reg.F)
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

func Op_RET_C(gb *GbCpu) {
	if gb.Reg.F&FlagC != 0 {
		gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
	} else {
		gb.Reg.PC++
	}
}

func Op_RETI(gb *GbCpu) {
	gb.InterruptsEnabled = true
	Op_RET(gb)
}

func Op_JP_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.PC = hl
}

func Op_JPnz(gb *GbCpu) {
	if gb.Reg.F&FlagZ == 0 {
		add := int8(gb.mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JPz(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		add := int8(gb.mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JP_Z_NN(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		addr := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
		gb.Reg.PC = addr
	} else {
		gb.Reg.PC += 3
	}
}

func Op_JP_C_NN(gb *GbCpu) {
	if gb.Reg.F&FlagC != 0 {
		addr := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
		gb.Reg.PC = addr
	} else {
		gb.Reg.PC += 3
	}
}

func Op_CPd8(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)

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
		gb.Reg.PC = uint16(gb.popFromStack()) + uint16(gb.popFromStack())<<8
	} else {
		gb.Reg.PC++
	}
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
	src := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + 0xFF00
	gb.Reg.A = gb.mem.GetByte(src)
	fmt.Printf("READ %04X from %04X\n", gb.Reg.A, src)
	gb.Reg.PC += 2
}

func Op_CALL(gb *GbCpu) {
	spc := gb.Reg.PC + 3
	gb.pushToStack(uint8(spc >> 8 & 0xFF))
	gb.pushToStack(uint8(spc & 0xFF))
	Op_JP(gb)
}

func Op_CALL_Z_a16(gb *GbCpu) {
	if gb.Reg.F&FlagZ != 0 {
		Op_CALL(gb)
	} else {
		gb.Reg.PC += 3
	}
}

func Op_CALL_NZ_a16(gb *GbCpu) {
	if gb.Reg.F&FlagZ == 0 {
		Op_CALL(gb)
	} else {
		gb.Reg.PC += 3
	}
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
	val := gb.mem.GetByte(addr)
	Do_Dec_Uint8(gb, &val)
	gb.mem.WriteByte(addr, val)
}

func Op_INC_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(addr)
	Do_Inc_Uint8(gb, &val)
	gb.mem.WriteByte(addr, val)
}

// 0xe6 AND A, n
func Op_ANDAn(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_And_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
}

// 0xe6 AND A, n
func Op_SUB_d8(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_Sub_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
}

func Op_SBC_A_d8(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_Sbc_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
}

func Op_LDHnA(gb *GbCpu) {
	dst := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + 0xFF00

	old := gb.mem.GetByte(dst)
	gb.mem.WriteByte(dst, gb.Reg.A)
	gb.Reg.PC += 2
	fmt.Printf("WROTE %04X to %04X, it was %04X -> %c\n", gb.mem.GetByte(dst), dst, old, gb.Reg.A)
}

func Op_LDAn(gb *GbCpu) {
	gb.Reg.A = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDBn(gb *GbCpu) {
	gb.Reg.B = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDCn(gb *GbCpu) {
	gb.Reg.C = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDDn(gb *GbCpu) {
	gb.Reg.D = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LDEn(gb *GbCpu) {
	gb.Reg.E = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

// Put value of A into location specified by DE
func Op_LD_DE_A(gb *GbCpu) {
	addr := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	gb.mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_BC_A(gb *GbCpu) {
	addr := uint16(gb.Reg.B)<<8 + uint16(gb.Reg.C)
	gb.mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_HL_d8(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	gb.mem.WriteByte(addr, val)
	gb.Reg.PC += 2
}

func Op_LD_HL_SP_r8(gb *GbCpu) {
	gb.Reg.F &= ^FlagMask
	operand := gb.mem.GetByte(gb.Reg.PC + 1)
	val := uint32(gb.Reg.SP) + uint32(operand)

	if val&0xFFFF0000 != 0 {
		gb.Reg.F |= FlagC
		fmt.Printf("Set C?\n")
	}

	if (gb.Reg.SP&0x0F)+uint16(operand&0x0F) > 0x0F {
		gb.Reg.F |= FlagH
	}

	hl := val & 0xFFFF

	gb.Reg.L = uint8(hl & 0xFF)
	gb.Reg.H = uint8((hl >> 8 & 0xFF))

	gb.Reg.PC += 2
}

func Op_LD_SP_HL(gb *GbCpu) {
	gb.Reg.SP = uint16(gb.Reg.L) + uint16(gb.Reg.H)<<8
	gb.Reg.PC++
}

func Op_LD_C_A(gb *GbCpu) {
	gb.mem.WriteByte(0xFF00+uint16(gb.Reg.C), gb.Reg.A)
	gb.Reg.PC++
}

func Op_LD_A_a16(gb *GbCpu) {
	addr := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	gb.Reg.A = gb.mem.GetByte(addr)
	gb.Reg.PC += 3
}

func Op_LD_a16_A(gb *GbCpu) {
	addr := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	gb.mem.WriteByte(addr, gb.Reg.A)
	gb.Reg.PC += 3
	fmt.Printf("LD %X -> %X\n", addr, gb.Reg.A)
}

func Op_LD_a16_SP(gb *GbCpu) {
	addr := uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	gb.mem.WriteByte(addr, uint8(gb.Reg.SP&0xFF))
	gb.mem.WriteByte(addr+1, uint8(gb.Reg.SP>>8)&0xFF)
	fmt.Printf("%X = %X, %X = %X\n", addr, gb.mem.GetByte(addr), addr+1, gb.mem.GetByte(addr+1))
	gb.Reg.PC += 3
}

func Op_LD_H_n(gb *GbCpu) {
	gb.Reg.H = gb.mem.GetByte(gb.Reg.PC + 1)
	gb.Reg.PC += 2
}

func Op_LD_L_n(gb *GbCpu) {
	gb.Reg.L = gb.mem.GetByte(gb.Reg.PC + 1)
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
	gb.Reg.SP = uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	gb.Reg.PC += 3
}

func Op_SUB_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Sub_88(gb, &gb.Reg.A, gb.mem.GetByte(addr))
}

func Op_SBC_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Sbc_88(gb, &gb.Reg.A, gb.mem.GetByte(addr))
}

func Op_ADC_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Adc_88(gb, &gb.Reg.A, gb.mem.GetByte(addr))
}

func Op_ADD_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Add_88(gb, &gb.Reg.A, gb.mem.GetByte(addr))
}

func Op_LD_A_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_B_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.B = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_C_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.C = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_D_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.D = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_E_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.E = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_H_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.H = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_L_HL(gb *GbCpu) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.L = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_BC(gb *GbCpu) {
	addr := uint16(gb.Reg.B)<<8 + uint16(gb.Reg.C)
	gb.Reg.A = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_DE(gb *GbCpu) {
	addr := uint16(gb.Reg.D)<<8 + uint16(gb.Reg.E)
	gb.Reg.A = gb.mem.GetByte(addr)
	gb.Reg.PC++
}

func Op_LD_A_HLi(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.mem.GetByte(val)
	val++
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

// Put value of A into location specified by HL, increment HL
func Op_LDI_HL_A(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.mem.WriteByte(val, gb.Reg.A)
	val++
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

// Put value of A into location specified by HL, decrement HL
func Op_LDD_HL_A(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.mem.WriteByte(val, gb.Reg.A)
	val--
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

// Store value specified by HL in A, decrement HL
func Op_LD_A_HLdec(gb *GbCpu) {
	val := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.Reg.A = gb.mem.GetByte(val)
	val--
	gb.Reg.L = uint8(val & 0xFF)
	gb.Reg.H = uint8((val >> 8 & 0xFF))
	gb.Reg.PC++
}

func Op_LD_HL_x(gb *GbCpu, value uint8) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	gb.mem.WriteByte(addr, value)
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

func Op_ADD_HL_SP(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	Do_Add_1616(gb, &hl, gb.Reg.SP)

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
	val := uint16(int8(gb.mem.GetByte(gb.Reg.PC)))

	half := uint8(0)
	if (gb.Reg.SP&0xF + val&0xF) > 0xF {
		half = FlagH
	}

	Do_Add_1616(gb, &gb.Reg.SP, val)

	// unlike raw add_1616, this does always clear
	// the zero flag and has a different understanding of
	// halfcarry :-/
	gb.Reg.F &= ^(FlagZ | FlagH)
	gb.Reg.F |= half
}

func Op_ADD_A_n(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
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

func Op_JP_NC(gb *GbCpu) {
	if (gb.Reg.F & FlagC) != 0 {
		gb.Reg.PC += 3
	} else {
		gb.Reg.PC = uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	}
}

func Op_JP_NZ(gb *GbCpu) {
	if (gb.Reg.F & FlagZ) != 0 {
		gb.Reg.PC += 3
	} else {
		gb.Reg.PC = uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
	}
}

func Op_JP(gb *GbCpu) {
	gb.Reg.PC = uint16(gb.mem.GetByte(gb.Reg.PC+1)) + uint16(gb.mem.GetByte(gb.Reg.PC+2))<<8
}

func Op_JR_n(gb *GbCpu) {
	add := int8(gb.mem.GetByte(gb.Reg.PC + 1))
	gb.Reg.PC += 2 + uint16(add)
}

func Op_JR_C_n(gb *GbCpu) {
	if gb.Reg.F&FlagC != 0 {
		add := int8(gb.mem.GetByte(gb.Reg.PC + 1))
		gb.Reg.PC += uint16(add)
	}
	gb.Reg.PC += 2
}

func Op_JR_NC_r8(gb *GbCpu) {
	if gb.Reg.F&FlagC == 0 {
		add := int8(gb.mem.GetByte(gb.Reg.PC + 1))
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

func Op_OR_n(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_Or_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++
}

func Op_OR_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Do_Or_88(gb, &gb.Reg.A, val)
}

func Op_XOR_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Do_Xor_88(gb, &gb.Reg.A, val)
}

func Op_CP_HL(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Do_Cp(gb, gb.Reg.A, val)
}

func Op_Rst(gb *GbCpu, pc uint16) {
	gb.Reg.PC++
	gb.pushToStack(uint8(gb.Reg.PC >> 8 & 0xFF))
	gb.pushToStack(uint8(gb.Reg.PC & 0xFF))
	gb.Reg.PC = pc
}

// Stolen from Sameboy's z80_cpu.c:319
// Would https://forums.nesdev.com/viewtopic.php?f=20&t=15944 also work?
func Op_DAA(gb *GbCpu) {
	a := gb.Reg.A
	f := gb.Reg.F & (FlagN | FlagC) // FlagN and FlagC are not touched

	if gb.Reg.F&FlagN != 0 {
		if gb.Reg.F&FlagH != 0 {
			gb.Reg.F &^= FlagH
			if gb.Reg.F&FlagC != 0 {
				a += 0x9A
			} else {
				a += 0xFA
			}
		} else if gb.Reg.F&FlagC != 0 {
			a += 0xA0
		}
	} else {
		n := uint16(a)

		if gb.Reg.F&FlagC != 0 {
			n += 0x100
		}
		if gb.Reg.F&FlagH != 0 {
			n += 0x06
			if n >= 0xA0 {
				n -= 0xA0
				f |= FlagC
			}

		} else {
			if n > 0x99 {
				n += 0x60
			}
			// WTF?!
			wtf := uint16(0)
			if n&0x0F > 9 {
				wtf = 6
			}
			n = (n & 0x0F) + wtf + (n & 0xFF0)
			if n&0xFF00 != 0 {
				f |= FlagC
			}
		}
		a = uint8(n)
	}

	if a == 0 {
		f |= FlagZ
	}

	gb.Reg.A = a
	gb.Reg.F = f
	gb.Reg.PC++

}

func Op_RRA(gb *GbCpu) {
	carry := byte(0)

	if gb.Reg.F&FlagC != 0 {
		carry = 1 << 7
	}

	gb.Reg.F &= ^FlagMask // clear all bits
	if gb.Reg.A&0x01 != 0 {
		gb.Reg.F |= FlagC
	}

	gb.Reg.A >>= 1
	gb.Reg.A += carry

	gb.Reg.PC++
}

func Op_CCF(gb *GbCpu) {
	if gb.Reg.F&FlagC != 0 {
		gb.Reg.F &= ^FlagC
	} else {
		gb.Reg.F |= FlagC
	}
	gb.Reg.F &= ^(FlagH | FlagN)
	gb.Reg.PC++
}

func Op_ADC_A_d8(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_Adc_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++
}

func Op_XOR_d8(gb *GbCpu) {
	val := gb.mem.GetByte(gb.Reg.PC + 1)
	Do_Xor_88(gb, &gb.Reg.A, val)
	gb.Reg.PC++ // +1 because we read one byte
}
