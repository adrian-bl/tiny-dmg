package cpu

import (
	"fmt"
)

func Cb_Disp(gb *GbCpu) {
	op := gb.mem.GetByte(gb.Reg.PC + 1)

	switch op {
	case 0x00:
		Cb_rlc(gb, &gb.Reg.B)
	case 0x01:
		Cb_rlc(gb, &gb.Reg.C)
	case 0x02:
		Cb_rlc(gb, &gb.Reg.D)
	case 0x03:
		Cb_rlc(gb, &gb.Reg.E)
	case 0x04:
		Cb_rlc(gb, &gb.Reg.H)
	case 0x05:
		Cb_rlc(gb, &gb.Reg.L)
	case 0x06:
		Cb_rlc_hlp(gb)
	case 0x07:
		Cb_rlc(gb, &gb.Reg.A)
	case 0x08:
		Cb_rrc(gb, &gb.Reg.B)
	case 0x09:
		Cb_rrc(gb, &gb.Reg.C)
	case 0x0A:
		Cb_rrc(gb, &gb.Reg.D)
	case 0x0B:
		Cb_rrc(gb, &gb.Reg.E)
	case 0x0C:
		Cb_rrc(gb, &gb.Reg.H)
	case 0x0D:
		Cb_rrc(gb, &gb.Reg.L)
	case 0x0E:
		Cb_rrc_hlp(gb)
	case 0x0F:
		Cb_rrc(gb, &gb.Reg.A)
	case 0x10:
		Cb_rl(gb, &gb.Reg.B)
	case 0x11:
		Cb_rl(gb, &gb.Reg.C)
	case 0x12:
		Cb_rl(gb, &gb.Reg.D)
	case 0x13:
		Cb_rl(gb, &gb.Reg.E)
	case 0x14:
		Cb_rl(gb, &gb.Reg.H)
	case 0x15:
		Cb_rl(gb, &gb.Reg.L)
	case 0x16:
		Cb_rl_hlp(gb)
	case 0x17:
		Cb_rl(gb, &gb.Reg.A)
	case 0x18:
		Cb_rr(gb, &gb.Reg.B)
	case 0x19:
		Cb_rr(gb, &gb.Reg.C)
	case 0x1A:
		Cb_rr(gb, &gb.Reg.D)
	case 0x1B:
		Cb_rr(gb, &gb.Reg.E)
	case 0x1C:
		Cb_rr(gb, &gb.Reg.H)
	case 0x1D:
		Cb_rr(gb, &gb.Reg.L)
	case 0x1E:
		Cb_rr_hlp(gb)
	case 0x1F:
		Cb_rr(gb, &gb.Reg.A)
	case 0x20:
		Cb_sla(gb, &gb.Reg.B)
	case 0x21:
		Cb_sla(gb, &gb.Reg.C)
	case 0x22:
		Cb_sla(gb, &gb.Reg.D)
	case 0x23:
		Cb_sla(gb, &gb.Reg.E)
	case 0x24:
		Cb_sla(gb, &gb.Reg.H)
	case 0x25:
		Cb_sla(gb, &gb.Reg.L)
	case 0x26:
		Cb_sla_hlp(gb)
	case 0x27:
		Cb_sla(gb, &gb.Reg.A)
	case 0x28:
		Cb_sra(gb, &gb.Reg.B)
	case 0x29:
		Cb_sra(gb, &gb.Reg.C)
	case 0x2A:
		Cb_sra(gb, &gb.Reg.D)
	case 0x2B:
		Cb_sra(gb, &gb.Reg.E)
	case 0x2C:
		Cb_sra(gb, &gb.Reg.H)
	case 0x2D:
		Cb_sra(gb, &gb.Reg.L)
	case 0x2F:
		Cb_sra(gb, &gb.Reg.A)
	case 0x2E:
		Cb_sra_hlp(gb)
	case 0x30:
		Cb_SwapReg(gb, &gb.Reg.B)
	case 0x31:
		Cb_SwapReg(gb, &gb.Reg.C)
	case 0x32:
		Cb_SwapReg(gb, &gb.Reg.D)
	case 0x33:
		Cb_SwapReg(gb, &gb.Reg.E)
	case 0x34:
		Cb_SwapReg(gb, &gb.Reg.H)
	case 0x35:
		Cb_SwapReg(gb, &gb.Reg.L)
	case 0x36:
		Cb_SwapReg_hlp(gb)
	case 0x37:
		Cb_SwapReg(gb, &gb.Reg.A)
	case 0x38:
		Cb_srl(gb, &gb.Reg.B)
	case 0x39:
		Cb_srl(gb, &gb.Reg.C)
	case 0x3A:
		Cb_srl(gb, &gb.Reg.D)
	case 0x3B:
		Cb_srl(gb, &gb.Reg.E)
	case 0x3C:
		Cb_srl(gb, &gb.Reg.H)
	case 0x3D:
		Cb_srl(gb, &gb.Reg.L)
	case 0x3E:
		Cb_srl_hlp(gb)
	case 0x3F:
		Cb_srl(gb, &gb.Reg.A)
	case 0x40:
		Cb_checkBit(gb, 0x00, gb.Reg.B)
	case 0x41:
		Cb_checkBit(gb, 0x00, gb.Reg.C)
	case 0x42:
		Cb_checkBit(gb, 0x00, gb.Reg.D)
	case 0x43:
		Cb_checkBit(gb, 0x00, gb.Reg.E)
	case 0x44:
		Cb_checkBit(gb, 0x00, gb.Reg.H)
	case 0x45:
		Cb_checkBit(gb, 0x00, gb.Reg.L)
	case 0x46:
		Cb_checkBitn(gb, 0x00, gb.Reg.H, gb.Reg.L)
	case 0x47:
		Cb_checkBit(gb, 0x00, gb.Reg.A)
	case 0x48:
		Cb_checkBit(gb, 0x01, gb.Reg.B)
	case 0x49:
		Cb_checkBit(gb, 0x01, gb.Reg.C)
	case 0x4A:
		Cb_checkBit(gb, 0x01, gb.Reg.D)
	case 0x4B:
		Cb_checkBit(gb, 0x01, gb.Reg.E)
	case 0x4C:
		Cb_checkBit(gb, 0x01, gb.Reg.H)
	case 0x4D:
		Cb_checkBit(gb, 0x01, gb.Reg.L)
	case 0x4E:
		Cb_checkBitn(gb, 0x01, gb.Reg.H, gb.Reg.L)
	case 0x4F:
		Cb_checkBit(gb, 0x01, gb.Reg.A)
	case 0x50:
		Cb_checkBit(gb, 0x02, gb.Reg.B)
	case 0x51:
		Cb_checkBit(gb, 0x02, gb.Reg.C)
	case 0x52:
		Cb_checkBit(gb, 0x02, gb.Reg.D)
	case 0x53:
		Cb_checkBit(gb, 0x02, gb.Reg.E)
	case 0x54:
		Cb_checkBit(gb, 0x02, gb.Reg.H)
	case 0x55:
		Cb_checkBit(gb, 0x02, gb.Reg.L)
	case 0x56:
		Cb_checkBitn(gb, 0x02, gb.Reg.H, gb.Reg.L)
	case 0x57:
		Cb_checkBit(gb, 0x02, gb.Reg.A)
	case 0x58:
		Cb_checkBit(gb, 0x03, gb.Reg.B)
	case 0x59:
		Cb_checkBit(gb, 0x03, gb.Reg.C)
	case 0x5A:
		Cb_checkBit(gb, 0x03, gb.Reg.D)
	case 0x5B:
		Cb_checkBit(gb, 0x03, gb.Reg.E)
	case 0x5C:
		Cb_checkBit(gb, 0x03, gb.Reg.H)
	case 0x5D:
		Cb_checkBit(gb, 0x03, gb.Reg.L)
	case 0x5E:
		Cb_checkBitn(gb, 0x03, gb.Reg.H, gb.Reg.L)
	case 0x5F:
		Cb_checkBit(gb, 0x03, gb.Reg.A)
	case 0x60:
		Cb_checkBit(gb, 0x04, gb.Reg.B)
	case 0x61:
		Cb_checkBit(gb, 0x04, gb.Reg.C)
	case 0x62:
		Cb_checkBit(gb, 0x04, gb.Reg.D)
	case 0x63:
		Cb_checkBit(gb, 0x04, gb.Reg.E)
	case 0x64:
		Cb_checkBit(gb, 0x04, gb.Reg.H)
	case 0x65:
		Cb_checkBit(gb, 0x04, gb.Reg.L)
	case 0x66:
		Cb_checkBitn(gb, 0x04, gb.Reg.H, gb.Reg.L)
	case 0x67:
		Cb_checkBit(gb, 0x04, gb.Reg.A)
	case 0x68:
		Cb_checkBit(gb, 0x05, gb.Reg.B)
	case 0x69:
		Cb_checkBit(gb, 0x05, gb.Reg.C)
	case 0x6A:
		Cb_checkBit(gb, 0x05, gb.Reg.D)
	case 0x6B:
		Cb_checkBit(gb, 0x05, gb.Reg.E)
	case 0x6C:
		Cb_checkBit(gb, 0x05, gb.Reg.H)
	case 0x6D:
		Cb_checkBit(gb, 0x05, gb.Reg.L)
	case 0x6E:
		Cb_checkBitn(gb, 0x05, gb.Reg.H, gb.Reg.L)
	case 0x6F:
		Cb_checkBit(gb, 0x05, gb.Reg.A)
	case 0x70:
		Cb_checkBit(gb, 0x06, gb.Reg.B)
	case 0x71:
		Cb_checkBit(gb, 0x06, gb.Reg.C)
	case 0x72:
		Cb_checkBit(gb, 0x06, gb.Reg.D)
	case 0x73:
		Cb_checkBit(gb, 0x06, gb.Reg.E)
	case 0x74:
		Cb_checkBit(gb, 0x06, gb.Reg.H)
	case 0x75:
		Cb_checkBit(gb, 0x06, gb.Reg.L)
	case 0x76:
		Cb_checkBitn(gb, 0x06, gb.Reg.H, gb.Reg.L)
	case 0x77:
		Cb_checkBit(gb, 0x06, gb.Reg.A)
	case 0x78:
		Cb_checkBit(gb, 0x07, gb.Reg.B)
	case 0x79:
		Cb_checkBit(gb, 0x07, gb.Reg.C)
	case 0x7A:
		Cb_checkBit(gb, 0x07, gb.Reg.D)
	case 0x7B:
		Cb_checkBit(gb, 0x07, gb.Reg.E)
	case 0x7C:
		Cb_checkBit(gb, 0x07, gb.Reg.H)
	case 0x7D:
		Cb_checkBit(gb, 0x07, gb.Reg.L)
	case 0x7E:
		Cb_checkBitn(gb, 0x07, gb.Reg.H, gb.Reg.L)
	case 0x7F:
		Cb_checkBit(gb, 0x07, gb.Reg.A)
	case 0x80:
		Cb_ResetBit(0x00, &gb.Reg.B)
	case 0x81:
		Cb_ResetBit(0x00, &gb.Reg.C)
	case 0x82:
		Cb_ResetBit(0x00, &gb.Reg.D)
	case 0x83:
		Cb_ResetBit(0x00, &gb.Reg.E)
	case 0x84:
		Cb_ResetBit(0x00, &gb.Reg.H)
	case 0x85:
		Cb_ResetBit(0x00, &gb.Reg.L)
	case 0x86:
		Cb_res_16(gb, 0x00, gb.Reg.H, gb.Reg.L)
	case 0x87:
		Cb_ResetBit(0x00, &gb.Reg.A)
	case 0x88:
		Cb_ResetBit(0x01, &gb.Reg.B)
	case 0x89:
		Cb_ResetBit(0x01, &gb.Reg.C)
	case 0x8A:
		Cb_ResetBit(0x01, &gb.Reg.D)
	case 0x8B:
		Cb_ResetBit(0x01, &gb.Reg.E)
	case 0x8C:
		Cb_ResetBit(0x01, &gb.Reg.H)
	case 0x8D:
		Cb_ResetBit(0x01, &gb.Reg.L)
	case 0x8E:
		Cb_res_16(gb, 0x01, gb.Reg.H, gb.Reg.L)
	case 0x8F:
		Cb_ResetBit(0x01, &gb.Reg.A)
	case 0x90:
		Cb_ResetBit(0x02, &gb.Reg.B)
	case 0x91:
		Cb_ResetBit(0x02, &gb.Reg.C)
	case 0x92:
		Cb_ResetBit(0x02, &gb.Reg.D)
	case 0x93:
		Cb_ResetBit(0x02, &gb.Reg.E)
	case 0x94:
		Cb_ResetBit(0x02, &gb.Reg.H)
	case 0x95:
		Cb_ResetBit(0x02, &gb.Reg.L)
	case 0x96:
		Cb_res_16(gb, 0x02, gb.Reg.H, gb.Reg.L)
	case 0x97:
		Cb_ResetBit(0x02, &gb.Reg.A)
	case 0x98:
		Cb_ResetBit(0x03, &gb.Reg.B)
	case 0x99:
		Cb_ResetBit(0x03, &gb.Reg.C)
	case 0x9A:
		Cb_ResetBit(0x03, &gb.Reg.D)
	case 0x9B:
		Cb_ResetBit(0x03, &gb.Reg.E)
	case 0x9C:
		Cb_ResetBit(0x03, &gb.Reg.H)
	case 0x9D:
		Cb_ResetBit(0x03, &gb.Reg.L)
	case 0x9E:
		Cb_res_16(gb, 0x03, gb.Reg.H, gb.Reg.L)
	case 0x9F:
		Cb_ResetBit(0x03, &gb.Reg.A)
	case 0xA0:
		Cb_ResetBit(0x04, &gb.Reg.B)
	case 0xA1:
		Cb_ResetBit(0x04, &gb.Reg.C)
	case 0xA2:
		Cb_ResetBit(0x04, &gb.Reg.D)
	case 0xA3:
		Cb_ResetBit(0x04, &gb.Reg.E)
	case 0xA4:
		Cb_ResetBit(0x04, &gb.Reg.H)
	case 0xA5:
		Cb_ResetBit(0x04, &gb.Reg.L)
	case 0xA6:
		Cb_res_16(gb, 0x04, gb.Reg.H, gb.Reg.L)
	case 0xA7:
		Cb_ResetBit(0x04, &gb.Reg.A)
	case 0xA8:
		Cb_ResetBit(0x05, &gb.Reg.B)
	case 0xA9:
		Cb_ResetBit(0x05, &gb.Reg.C)
	case 0xAA:
		Cb_ResetBit(0x05, &gb.Reg.D)
	case 0xAB:
		Cb_ResetBit(0x05, &gb.Reg.E)
	case 0xAC:
		Cb_ResetBit(0x05, &gb.Reg.H)
	case 0xAD:
		Cb_ResetBit(0x05, &gb.Reg.L)
	case 0xAE:
		Cb_res_16(gb, 0x05, gb.Reg.H, gb.Reg.L)
	case 0xAF:
		Cb_ResetBit(0x05, &gb.Reg.A)
	case 0xB0:
		Cb_ResetBit(0x06, &gb.Reg.B)
	case 0xB1:
		Cb_ResetBit(0x06, &gb.Reg.C)
	case 0xB2:
		Cb_ResetBit(0x06, &gb.Reg.D)
	case 0xB3:
		Cb_ResetBit(0x06, &gb.Reg.E)
	case 0xB4:
		Cb_ResetBit(0x06, &gb.Reg.H)
	case 0xB5:
		Cb_ResetBit(0x06, &gb.Reg.L)
	case 0xB6:
		Cb_res_16(gb, 0x06, gb.Reg.H, gb.Reg.L)
	case 0xB7:
		Cb_ResetBit(0x06, &gb.Reg.A)
	case 0xB8:
		Cb_ResetBit(0x07, &gb.Reg.B)
	case 0xB9:
		Cb_ResetBit(0x07, &gb.Reg.C)
	case 0xBA:
		Cb_ResetBit(0x07, &gb.Reg.D)
	case 0xBB:
		Cb_ResetBit(0x07, &gb.Reg.E)
	case 0xBC:
		Cb_ResetBit(0x07, &gb.Reg.H)
	case 0xBD:
		Cb_ResetBit(0x07, &gb.Reg.L)
	case 0xBE:
		Cb_res_16(gb, 0x07, gb.Reg.H, gb.Reg.L)
	case 0xBF:
		Cb_ResetBit(0x07, &gb.Reg.A)
	case 0xC0:
		Cb_set(0x00, &gb.Reg.B)
	case 0xC1:
		Cb_set(0x00, &gb.Reg.C)
	case 0xC2:
		Cb_set(0x00, &gb.Reg.D)
	case 0xC3:
		Cb_set(0x00, &gb.Reg.E)
	case 0xC4:
		Cb_set(0x00, &gb.Reg.H)
	case 0xC5:
		Cb_set(0x00, &gb.Reg.L)
	case 0xC6:
		Cb_SetHlpBit(gb, 0x00)
	case 0xC7:
		Cb_set(0x00, &gb.Reg.A)
	case 0xC8:
		Cb_set(0x01, &gb.Reg.B)
	case 0xC9:
		Cb_set(0x01, &gb.Reg.C)
	case 0xCA:
		Cb_set(0x01, &gb.Reg.D)
	case 0xCB:
		Cb_set(0x01, &gb.Reg.E)
	case 0xCC:
		Cb_set(0x01, &gb.Reg.H)
	case 0xCD:
		Cb_set(0x01, &gb.Reg.L)
	case 0xCE:
		Cb_SetHlpBit(gb, 0x01)
	case 0xCF:
		Cb_set(0x01, &gb.Reg.A)
	case 0xD0:
		Cb_set(0x02, &gb.Reg.B)
	case 0xD1:
		Cb_set(0x02, &gb.Reg.C)
	case 0xD2:
		Cb_set(0x02, &gb.Reg.D)
	case 0xD3:
		Cb_set(0x02, &gb.Reg.E)
	case 0xD4:
		Cb_set(0x02, &gb.Reg.H)
	case 0xD5:
		Cb_set(0x02, &gb.Reg.L)
	case 0xD6:
		Cb_SetHlpBit(gb, 0x02)
	case 0xD7:
		Cb_set(0x02, &gb.Reg.A)
	case 0xD8:
		Cb_set(0x03, &gb.Reg.B)
	case 0xD9:
		Cb_set(0x03, &gb.Reg.C)
	case 0xDA:
		Cb_set(0x03, &gb.Reg.D)
	case 0xDB:
		Cb_set(0x03, &gb.Reg.E)
	case 0xDC:
		Cb_set(0x03, &gb.Reg.H)
	case 0xDD:
		Cb_set(0x03, &gb.Reg.L)
	case 0xDE:
		Cb_SetHlpBit(gb, 0x03)
	case 0xDF:
		Cb_set(0x03, &gb.Reg.A)
	case 0xE0:
		Cb_set(0x04, &gb.Reg.B)
	case 0xE1:
		Cb_set(0x04, &gb.Reg.C)
	case 0xE2:
		Cb_set(0x04, &gb.Reg.D)
	case 0xE3:
		Cb_set(0x04, &gb.Reg.E)
	case 0xE4:
		Cb_set(0x04, &gb.Reg.H)
	case 0xE5:
		Cb_set(0x04, &gb.Reg.L)
	case 0xE6:
		Cb_SetHlpBit(gb, 0x04)
	case 0xE7:
		Cb_set(0x04, &gb.Reg.A)
	case 0xE8:
		Cb_set(0x05, &gb.Reg.B)
	case 0xE9:
		Cb_set(0x05, &gb.Reg.C)
	case 0xEA:
		Cb_set(0x05, &gb.Reg.D)
	case 0xEB:
		Cb_set(0x05, &gb.Reg.E)
	case 0xEC:
		Cb_set(0x05, &gb.Reg.H)
	case 0xED:
		Cb_set(0x05, &gb.Reg.L)
	case 0xEE:
		Cb_SetHlpBit(gb, 0x05)
	case 0xEF:
		Cb_set(0x05, &gb.Reg.A)
	case 0xF0:
		Cb_set(0x06, &gb.Reg.B)
	case 0xF1:
		Cb_set(0x06, &gb.Reg.C)
	case 0xF2:
		Cb_set(0x06, &gb.Reg.D)
	case 0xF3:
		Cb_set(0x06, &gb.Reg.E)
	case 0xF4:
		Cb_set(0x06, &gb.Reg.H)
	case 0xF5:
		Cb_set(0x06, &gb.Reg.L)
	case 0xF6:
		Cb_SetHlpBit(gb, 0x06)
	case 0xF7:
		Cb_set(0x06, &gb.Reg.A)
	case 0xF8:
		Cb_set(0x07, &gb.Reg.B)
	case 0xF9:
		Cb_set(0x07, &gb.Reg.C)
	case 0xFA:
		Cb_set(0x07, &gb.Reg.D)
	case 0xFB:
		Cb_set(0x07, &gb.Reg.E)
	case 0xFC:
		Cb_set(0x07, &gb.Reg.H)
	case 0xFD:
		Cb_set(0x07, &gb.Reg.L)
	case 0xFE:
		Cb_SetHlpBit(gb, 0x07)
	case 0xFF:
		Cb_set(0x07, &gb.Reg.A)
	default:
		fmt.Printf("Unknown CB opcode: %02X", op)
		gb.crash()
	}
	gb.Reg.PC += 2
}

func Cb_ResetBit(bit uint8, target *uint8) {
	*target &= ^(1 << bit)
}

func Cb_set(bit uint8, target *uint8) {
	*target |= (1 << bit)
}

func Cb_SetHlpBit(gb *GbCpu, bit uint8) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(addr)
	Cb_set(bit, &val)
	gb.mem.WriteByte(addr, val)
}

func Cb_SwapReg(gb *GbCpu, target *uint8) {
	*target = ((*target & 0xF) << 4) | ((*target & 0xF0) >> 4)

	gb.Reg.F &= ^FlagMask // clear all bits

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_sla(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	if (*target & 0x80) != 0 {
		gb.Reg.F |= FlagC
	}

	*target <<= 1

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_sra(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	carry := *target & 0x01
	if carry != 0 {
		gb.Reg.F |= FlagC
	} else {
		gb.Reg.F &= ^FlagC
	}

	*target = (*target & 0x80) | (*target >> 1)

	if *target != 0 {
		gb.Reg.F &= ^FlagZ
	} else {
		gb.Reg.F |= FlagZ
	}
}

func Cb_checkBitn(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := uint8(gb.mem.GetByte(addr))
	Cb_checkBit(gb, bit, val)
}

func Cb_checkBit(gb *GbCpu, bit uint8, value uint8) {
	gb.Reg.F &= ^(FlagZ | FlagN)
	gb.Reg.F |= FlagH

	if value&(1<<bit) == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_srl(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	if *target&0x01 != 0 {
		gb.Reg.F |= FlagC
	}

	*target = *target >> 1

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_rl(gb *GbCpu, target *uint8) {
	oldMask := gb.Reg.F
	gb.Reg.F &= ^FlagMask

	carry := uint8(0)
	if oldMask&FlagC != 0 {
		carry = 1
	}

	if *target&0x80 != 0 {
		gb.Reg.F |= FlagC
	}

	*target = (*target << 1) | carry

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_rlc(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	carry := uint8(0)
	if *target&0x80 != 0 {
		carry = 1
		gb.Reg.F |= FlagC
	}
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
	*target = (*target << 1) | carry
}

func Cb_rlc_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_rlc(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_rrc_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_rrc(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_rl_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_rl(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_rr_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_rr(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_sla_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_sla(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_srl_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_srl(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_sra_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_sra(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func Cb_SwapReg_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	Cb_SwapReg(gb, &val)
	gb.mem.WriteByte(hl, val)
}

// 9 bit rotation to right
func Cb_rr(gb *GbCpu, target *uint8) {
	carry := uint8(0)
	bit0 := (*target & 0x1) != 0

	if gb.Reg.F&FlagC != 0 {
		carry = 1 << 7 // carry is copied to bit 7
	}

	gb.Reg.F &= ^FlagMask
	*target = (*target >> 1) | carry

	if bit0 {
		gb.Reg.F |= FlagC
	}
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

// Note that this is NOT the same as math.go:Do_Rrc
func Cb_rrc(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	carry := *target & 0x01
	*target = (*target >> 1) | (carry << 7)

	if carry != 0 {
		gb.Reg.F |= FlagC
	}
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func Cb_res_16(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := gb.mem.GetByte(addr) & ^(1 << bit)
	gb.mem.WriteByte(addr, val)
}
