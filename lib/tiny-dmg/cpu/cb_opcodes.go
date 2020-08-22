package cpu

import (
	"fmt"
)

func Cb_Disp(gb *GbCpu) {
	op := gb.mem.GetByte(gb.Reg.PC + 1)

	switch op {
	case 0x00:
		cb_rlc(gb, &gb.Reg.B)
	case 0x01:
		cb_rlc(gb, &gb.Reg.C)
	case 0x02:
		cb_rlc(gb, &gb.Reg.D)
	case 0x03:
		cb_rlc(gb, &gb.Reg.E)
	case 0x04:
		cb_rlc(gb, &gb.Reg.H)
	case 0x05:
		cb_rlc(gb, &gb.Reg.L)
	case 0x06:
		cb_rlc_hlp(gb)
	case 0x07:
		cb_rlc(gb, &gb.Reg.A)
	case 0x08:
		cb_rrc(gb, &gb.Reg.B)
	case 0x09:
		cb_rrc(gb, &gb.Reg.C)
	case 0x0A:
		cb_rrc(gb, &gb.Reg.D)
	case 0x0B:
		cb_rrc(gb, &gb.Reg.E)
	case 0x0C:
		cb_rrc(gb, &gb.Reg.H)
	case 0x0D:
		cb_rrc(gb, &gb.Reg.L)
	case 0x0E:
		cb_rrc_hlp(gb)
	case 0x0F:
		cb_rrc(gb, &gb.Reg.A)
	case 0x10:
		cb_rl(gb, &gb.Reg.B)
	case 0x11:
		cb_rl(gb, &gb.Reg.C)
	case 0x12:
		cb_rl(gb, &gb.Reg.D)
	case 0x13:
		cb_rl(gb, &gb.Reg.E)
	case 0x14:
		cb_rl(gb, &gb.Reg.H)
	case 0x15:
		cb_rl(gb, &gb.Reg.L)
	case 0x16:
		cb_rl_hlp(gb)
	case 0x17:
		cb_rl(gb, &gb.Reg.A)
	case 0x18:
		cb_rr(gb, &gb.Reg.B)
	case 0x19:
		cb_rr(gb, &gb.Reg.C)
	case 0x1A:
		cb_rr(gb, &gb.Reg.D)
	case 0x1B:
		cb_rr(gb, &gb.Reg.E)
	case 0x1C:
		cb_rr(gb, &gb.Reg.H)
	case 0x1D:
		cb_rr(gb, &gb.Reg.L)
	case 0x1E:
		cb_rr_hlp(gb)
	case 0x1F:
		cb_rr(gb, &gb.Reg.A)
	case 0x20:
		cb_sla(gb, &gb.Reg.B)
	case 0x21:
		cb_sla(gb, &gb.Reg.C)
	case 0x22:
		cb_sla(gb, &gb.Reg.D)
	case 0x23:
		cb_sla(gb, &gb.Reg.E)
	case 0x24:
		cb_sla(gb, &gb.Reg.H)
	case 0x25:
		cb_sla(gb, &gb.Reg.L)
	case 0x26:
		cb_sla_hlp(gb)
	case 0x27:
		cb_sla(gb, &gb.Reg.A)
	case 0x28:
		cb_sra(gb, &gb.Reg.B)
	case 0x29:
		cb_sra(gb, &gb.Reg.C)
	case 0x2A:
		cb_sra(gb, &gb.Reg.D)
	case 0x2B:
		cb_sra(gb, &gb.Reg.E)
	case 0x2C:
		cb_sra(gb, &gb.Reg.H)
	case 0x2D:
		cb_sra(gb, &gb.Reg.L)
	case 0x2F:
		cb_sra(gb, &gb.Reg.A)
	case 0x2E:
		cb_sra_hlp(gb)
	case 0x30:
		cb_swap(gb, &gb.Reg.B)
	case 0x31:
		cb_swap(gb, &gb.Reg.C)
	case 0x32:
		cb_swap(gb, &gb.Reg.D)
	case 0x33:
		cb_swap(gb, &gb.Reg.E)
	case 0x34:
		cb_swap(gb, &gb.Reg.H)
	case 0x35:
		cb_swap(gb, &gb.Reg.L)
	case 0x36:
		cb_swap_hlp(gb)
	case 0x37:
		cb_swap(gb, &gb.Reg.A)
	case 0x38:
		cb_srl(gb, &gb.Reg.B)
	case 0x39:
		cb_srl(gb, &gb.Reg.C)
	case 0x3A:
		cb_srl(gb, &gb.Reg.D)
	case 0x3B:
		cb_srl(gb, &gb.Reg.E)
	case 0x3C:
		cb_srl(gb, &gb.Reg.H)
	case 0x3D:
		cb_srl(gb, &gb.Reg.L)
	case 0x3E:
		cb_srl_hlp(gb)
	case 0x3F:
		cb_srl(gb, &gb.Reg.A)
	case 0x40:
		cb_checkbit(gb, 0x00, gb.Reg.B)
	case 0x41:
		cb_checkbit(gb, 0x00, gb.Reg.C)
	case 0x42:
		cb_checkbit(gb, 0x00, gb.Reg.D)
	case 0x43:
		cb_checkbit(gb, 0x00, gb.Reg.E)
	case 0x44:
		cb_checkbit(gb, 0x00, gb.Reg.H)
	case 0x45:
		cb_checkbit(gb, 0x00, gb.Reg.L)
	case 0x46:
		cb_checkbit_n(gb, 0x00, gb.Reg.H, gb.Reg.L)
	case 0x47:
		cb_checkbit(gb, 0x00, gb.Reg.A)
	case 0x48:
		cb_checkbit(gb, 0x01, gb.Reg.B)
	case 0x49:
		cb_checkbit(gb, 0x01, gb.Reg.C)
	case 0x4A:
		cb_checkbit(gb, 0x01, gb.Reg.D)
	case 0x4B:
		cb_checkbit(gb, 0x01, gb.Reg.E)
	case 0x4C:
		cb_checkbit(gb, 0x01, gb.Reg.H)
	case 0x4D:
		cb_checkbit(gb, 0x01, gb.Reg.L)
	case 0x4E:
		cb_checkbit_n(gb, 0x01, gb.Reg.H, gb.Reg.L)
	case 0x4F:
		cb_checkbit(gb, 0x01, gb.Reg.A)
	case 0x50:
		cb_checkbit(gb, 0x02, gb.Reg.B)
	case 0x51:
		cb_checkbit(gb, 0x02, gb.Reg.C)
	case 0x52:
		cb_checkbit(gb, 0x02, gb.Reg.D)
	case 0x53:
		cb_checkbit(gb, 0x02, gb.Reg.E)
	case 0x54:
		cb_checkbit(gb, 0x02, gb.Reg.H)
	case 0x55:
		cb_checkbit(gb, 0x02, gb.Reg.L)
	case 0x56:
		cb_checkbit_n(gb, 0x02, gb.Reg.H, gb.Reg.L)
	case 0x57:
		cb_checkbit(gb, 0x02, gb.Reg.A)
	case 0x58:
		cb_checkbit(gb, 0x03, gb.Reg.B)
	case 0x59:
		cb_checkbit(gb, 0x03, gb.Reg.C)
	case 0x5A:
		cb_checkbit(gb, 0x03, gb.Reg.D)
	case 0x5B:
		cb_checkbit(gb, 0x03, gb.Reg.E)
	case 0x5C:
		cb_checkbit(gb, 0x03, gb.Reg.H)
	case 0x5D:
		cb_checkbit(gb, 0x03, gb.Reg.L)
	case 0x5E:
		cb_checkbit_n(gb, 0x03, gb.Reg.H, gb.Reg.L)
	case 0x5F:
		cb_checkbit(gb, 0x03, gb.Reg.A)
	case 0x60:
		cb_checkbit(gb, 0x04, gb.Reg.B)
	case 0x61:
		cb_checkbit(gb, 0x04, gb.Reg.C)
	case 0x62:
		cb_checkbit(gb, 0x04, gb.Reg.D)
	case 0x63:
		cb_checkbit(gb, 0x04, gb.Reg.E)
	case 0x64:
		cb_checkbit(gb, 0x04, gb.Reg.H)
	case 0x65:
		cb_checkbit(gb, 0x04, gb.Reg.L)
	case 0x66:
		cb_checkbit_n(gb, 0x04, gb.Reg.H, gb.Reg.L)
	case 0x67:
		cb_checkbit(gb, 0x04, gb.Reg.A)
	case 0x68:
		cb_checkbit(gb, 0x05, gb.Reg.B)
	case 0x69:
		cb_checkbit(gb, 0x05, gb.Reg.C)
	case 0x6A:
		cb_checkbit(gb, 0x05, gb.Reg.D)
	case 0x6B:
		cb_checkbit(gb, 0x05, gb.Reg.E)
	case 0x6C:
		cb_checkbit(gb, 0x05, gb.Reg.H)
	case 0x6D:
		cb_checkbit(gb, 0x05, gb.Reg.L)
	case 0x6E:
		cb_checkbit_n(gb, 0x05, gb.Reg.H, gb.Reg.L)
	case 0x6F:
		cb_checkbit(gb, 0x05, gb.Reg.A)
	case 0x70:
		cb_checkbit(gb, 0x06, gb.Reg.B)
	case 0x71:
		cb_checkbit(gb, 0x06, gb.Reg.C)
	case 0x72:
		cb_checkbit(gb, 0x06, gb.Reg.D)
	case 0x73:
		cb_checkbit(gb, 0x06, gb.Reg.E)
	case 0x74:
		cb_checkbit(gb, 0x06, gb.Reg.H)
	case 0x75:
		cb_checkbit(gb, 0x06, gb.Reg.L)
	case 0x76:
		cb_checkbit_n(gb, 0x06, gb.Reg.H, gb.Reg.L)
	case 0x77:
		cb_checkbit(gb, 0x06, gb.Reg.A)
	case 0x78:
		cb_checkbit(gb, 0x07, gb.Reg.B)
	case 0x79:
		cb_checkbit(gb, 0x07, gb.Reg.C)
	case 0x7A:
		cb_checkbit(gb, 0x07, gb.Reg.D)
	case 0x7B:
		cb_checkbit(gb, 0x07, gb.Reg.E)
	case 0x7C:
		cb_checkbit(gb, 0x07, gb.Reg.H)
	case 0x7D:
		cb_checkbit(gb, 0x07, gb.Reg.L)
	case 0x7E:
		cb_checkbit_n(gb, 0x07, gb.Reg.H, gb.Reg.L)
	case 0x7F:
		cb_checkbit(gb, 0x07, gb.Reg.A)
	case 0x80:
		cb_resetbit(0x00, &gb.Reg.B)
	case 0x81:
		cb_resetbit(0x00, &gb.Reg.C)
	case 0x82:
		cb_resetbit(0x00, &gb.Reg.D)
	case 0x83:
		cb_resetbit(0x00, &gb.Reg.E)
	case 0x84:
		cb_resetbit(0x00, &gb.Reg.H)
	case 0x85:
		cb_resetbit(0x00, &gb.Reg.L)
	case 0x86:
		cb_res_16(gb, 0x00, gb.Reg.H, gb.Reg.L)
	case 0x87:
		cb_resetbit(0x00, &gb.Reg.A)
	case 0x88:
		cb_resetbit(0x01, &gb.Reg.B)
	case 0x89:
		cb_resetbit(0x01, &gb.Reg.C)
	case 0x8A:
		cb_resetbit(0x01, &gb.Reg.D)
	case 0x8B:
		cb_resetbit(0x01, &gb.Reg.E)
	case 0x8C:
		cb_resetbit(0x01, &gb.Reg.H)
	case 0x8D:
		cb_resetbit(0x01, &gb.Reg.L)
	case 0x8E:
		cb_res_16(gb, 0x01, gb.Reg.H, gb.Reg.L)
	case 0x8F:
		cb_resetbit(0x01, &gb.Reg.A)
	case 0x90:
		cb_resetbit(0x02, &gb.Reg.B)
	case 0x91:
		cb_resetbit(0x02, &gb.Reg.C)
	case 0x92:
		cb_resetbit(0x02, &gb.Reg.D)
	case 0x93:
		cb_resetbit(0x02, &gb.Reg.E)
	case 0x94:
		cb_resetbit(0x02, &gb.Reg.H)
	case 0x95:
		cb_resetbit(0x02, &gb.Reg.L)
	case 0x96:
		cb_res_16(gb, 0x02, gb.Reg.H, gb.Reg.L)
	case 0x97:
		cb_resetbit(0x02, &gb.Reg.A)
	case 0x98:
		cb_resetbit(0x03, &gb.Reg.B)
	case 0x99:
		cb_resetbit(0x03, &gb.Reg.C)
	case 0x9A:
		cb_resetbit(0x03, &gb.Reg.D)
	case 0x9B:
		cb_resetbit(0x03, &gb.Reg.E)
	case 0x9C:
		cb_resetbit(0x03, &gb.Reg.H)
	case 0x9D:
		cb_resetbit(0x03, &gb.Reg.L)
	case 0x9E:
		cb_res_16(gb, 0x03, gb.Reg.H, gb.Reg.L)
	case 0x9F:
		cb_resetbit(0x03, &gb.Reg.A)
	case 0xA0:
		cb_resetbit(0x04, &gb.Reg.B)
	case 0xA1:
		cb_resetbit(0x04, &gb.Reg.C)
	case 0xA2:
		cb_resetbit(0x04, &gb.Reg.D)
	case 0xA3:
		cb_resetbit(0x04, &gb.Reg.E)
	case 0xA4:
		cb_resetbit(0x04, &gb.Reg.H)
	case 0xA5:
		cb_resetbit(0x04, &gb.Reg.L)
	case 0xA6:
		cb_res_16(gb, 0x04, gb.Reg.H, gb.Reg.L)
	case 0xA7:
		cb_resetbit(0x04, &gb.Reg.A)
	case 0xA8:
		cb_resetbit(0x05, &gb.Reg.B)
	case 0xA9:
		cb_resetbit(0x05, &gb.Reg.C)
	case 0xAA:
		cb_resetbit(0x05, &gb.Reg.D)
	case 0xAB:
		cb_resetbit(0x05, &gb.Reg.E)
	case 0xAC:
		cb_resetbit(0x05, &gb.Reg.H)
	case 0xAD:
		cb_resetbit(0x05, &gb.Reg.L)
	case 0xAE:
		cb_res_16(gb, 0x05, gb.Reg.H, gb.Reg.L)
	case 0xAF:
		cb_resetbit(0x05, &gb.Reg.A)
	case 0xB0:
		cb_resetbit(0x06, &gb.Reg.B)
	case 0xB1:
		cb_resetbit(0x06, &gb.Reg.C)
	case 0xB2:
		cb_resetbit(0x06, &gb.Reg.D)
	case 0xB3:
		cb_resetbit(0x06, &gb.Reg.E)
	case 0xB4:
		cb_resetbit(0x06, &gb.Reg.H)
	case 0xB5:
		cb_resetbit(0x06, &gb.Reg.L)
	case 0xB6:
		cb_res_16(gb, 0x06, gb.Reg.H, gb.Reg.L)
	case 0xB7:
		cb_resetbit(0x06, &gb.Reg.A)
	case 0xB8:
		cb_resetbit(0x07, &gb.Reg.B)
	case 0xB9:
		cb_resetbit(0x07, &gb.Reg.C)
	case 0xBA:
		cb_resetbit(0x07, &gb.Reg.D)
	case 0xBB:
		cb_resetbit(0x07, &gb.Reg.E)
	case 0xBC:
		cb_resetbit(0x07, &gb.Reg.H)
	case 0xBD:
		cb_resetbit(0x07, &gb.Reg.L)
	case 0xBE:
		cb_res_16(gb, 0x07, gb.Reg.H, gb.Reg.L)
	case 0xBF:
		cb_resetbit(0x07, &gb.Reg.A)
	case 0xC0:
		cb_set(0x00, &gb.Reg.B)
	case 0xC1:
		cb_set(0x00, &gb.Reg.C)
	case 0xC2:
		cb_set(0x00, &gb.Reg.D)
	case 0xC3:
		cb_set(0x00, &gb.Reg.E)
	case 0xC4:
		cb_set(0x00, &gb.Reg.H)
	case 0xC5:
		cb_set(0x00, &gb.Reg.L)
	case 0xC6:
		cb_setbit_hlp(gb, 0x00)
	case 0xC7:
		cb_set(0x00, &gb.Reg.A)
	case 0xC8:
		cb_set(0x01, &gb.Reg.B)
	case 0xC9:
		cb_set(0x01, &gb.Reg.C)
	case 0xCA:
		cb_set(0x01, &gb.Reg.D)
	case 0xCB:
		cb_set(0x01, &gb.Reg.E)
	case 0xCC:
		cb_set(0x01, &gb.Reg.H)
	case 0xCD:
		cb_set(0x01, &gb.Reg.L)
	case 0xCE:
		cb_setbit_hlp(gb, 0x01)
	case 0xCF:
		cb_set(0x01, &gb.Reg.A)
	case 0xD0:
		cb_set(0x02, &gb.Reg.B)
	case 0xD1:
		cb_set(0x02, &gb.Reg.C)
	case 0xD2:
		cb_set(0x02, &gb.Reg.D)
	case 0xD3:
		cb_set(0x02, &gb.Reg.E)
	case 0xD4:
		cb_set(0x02, &gb.Reg.H)
	case 0xD5:
		cb_set(0x02, &gb.Reg.L)
	case 0xD6:
		cb_setbit_hlp(gb, 0x02)
	case 0xD7:
		cb_set(0x02, &gb.Reg.A)
	case 0xD8:
		cb_set(0x03, &gb.Reg.B)
	case 0xD9:
		cb_set(0x03, &gb.Reg.C)
	case 0xDA:
		cb_set(0x03, &gb.Reg.D)
	case 0xDB:
		cb_set(0x03, &gb.Reg.E)
	case 0xDC:
		cb_set(0x03, &gb.Reg.H)
	case 0xDD:
		cb_set(0x03, &gb.Reg.L)
	case 0xDE:
		cb_setbit_hlp(gb, 0x03)
	case 0xDF:
		cb_set(0x03, &gb.Reg.A)
	case 0xE0:
		cb_set(0x04, &gb.Reg.B)
	case 0xE1:
		cb_set(0x04, &gb.Reg.C)
	case 0xE2:
		cb_set(0x04, &gb.Reg.D)
	case 0xE3:
		cb_set(0x04, &gb.Reg.E)
	case 0xE4:
		cb_set(0x04, &gb.Reg.H)
	case 0xE5:
		cb_set(0x04, &gb.Reg.L)
	case 0xE6:
		cb_setbit_hlp(gb, 0x04)
	case 0xE7:
		cb_set(0x04, &gb.Reg.A)
	case 0xE8:
		cb_set(0x05, &gb.Reg.B)
	case 0xE9:
		cb_set(0x05, &gb.Reg.C)
	case 0xEA:
		cb_set(0x05, &gb.Reg.D)
	case 0xEB:
		cb_set(0x05, &gb.Reg.E)
	case 0xEC:
		cb_set(0x05, &gb.Reg.H)
	case 0xED:
		cb_set(0x05, &gb.Reg.L)
	case 0xEE:
		cb_setbit_hlp(gb, 0x05)
	case 0xEF:
		cb_set(0x05, &gb.Reg.A)
	case 0xF0:
		cb_set(0x06, &gb.Reg.B)
	case 0xF1:
		cb_set(0x06, &gb.Reg.C)
	case 0xF2:
		cb_set(0x06, &gb.Reg.D)
	case 0xF3:
		cb_set(0x06, &gb.Reg.E)
	case 0xF4:
		cb_set(0x06, &gb.Reg.H)
	case 0xF5:
		cb_set(0x06, &gb.Reg.L)
	case 0xF6:
		cb_setbit_hlp(gb, 0x06)
	case 0xF7:
		cb_set(0x06, &gb.Reg.A)
	case 0xF8:
		cb_set(0x07, &gb.Reg.B)
	case 0xF9:
		cb_set(0x07, &gb.Reg.C)
	case 0xFA:
		cb_set(0x07, &gb.Reg.D)
	case 0xFB:
		cb_set(0x07, &gb.Reg.E)
	case 0xFC:
		cb_set(0x07, &gb.Reg.H)
	case 0xFD:
		cb_set(0x07, &gb.Reg.L)
	case 0xFE:
		cb_setbit_hlp(gb, 0x07)
	case 0xFF:
		cb_set(0x07, &gb.Reg.A)
	default:
		fmt.Printf("Unknown CB opcode: %02X", op)
		gb.crash()
	}
	gb.Reg.PC += 2
}

func cb_resetbit(bit uint8, target *uint8) {
	*target &= ^(1 << bit)
}

func cb_set(bit uint8, target *uint8) {
	*target |= (1 << bit)
}

func cb_setbit_hlp(gb *GbCpu, bit uint8) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(addr)
	cb_set(bit, &val)
	gb.mem.WriteByte(addr, val)
}

func cb_swap(gb *GbCpu, target *uint8) {
	*target = ((*target & 0xF) << 4) | ((*target & 0xF0) >> 4)

	gb.Reg.F &= ^FlagMask // clear all bits

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func cb_sla(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	if (*target & 0x80) != 0 {
		gb.Reg.F |= FlagC
	}

	*target <<= 1

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func cb_sra(gb *GbCpu, target *uint8) {
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

func cb_checkbit_n(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := uint8(gb.mem.GetByte(addr))
	cb_checkbit(gb, bit, val)
}

func cb_checkbit(gb *GbCpu, bit uint8, value uint8) {
	gb.Reg.F &= ^(FlagZ | FlagN)
	gb.Reg.F |= FlagH

	if value&(1<<bit) == 0 {
		gb.Reg.F |= FlagZ
	}
}

func cb_srl(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	if *target&0x01 != 0 {
		gb.Reg.F |= FlagC
	}

	*target = *target >> 1

	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}

func cb_rl(gb *GbCpu, target *uint8) {
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

func cb_rlc(gb *GbCpu, target *uint8) {
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

func cb_rlc_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_rlc(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_rrc_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_rrc(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_rl_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_rl(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_rr_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_rr(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_sla_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_sla(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_srl_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_srl(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_sra_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_sra(gb, &val)
	gb.mem.WriteByte(hl, val)
}

func cb_swap_hlp(gb *GbCpu) {
	hl := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.mem.GetByte(hl)
	cb_swap(gb, &val)
	gb.mem.WriteByte(hl, val)
}

// 9 bit rotation to right
func cb_rr(gb *GbCpu, target *uint8) {
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
func cb_rrc(gb *GbCpu, target *uint8) {
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

func cb_res_16(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := gb.mem.GetByte(addr) & ^(1 << bit)
	gb.mem.WriteByte(addr, val)
}
