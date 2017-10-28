package cpu

import (
	"fmt"
)

func Cb_Disp(gb *GbCpu) {
	op := gb.mem.GetByte(gb.Reg.PC + 1)
	fmt.Printf("-> CB %02X\n", op)

	switch op {
	case 0x09:
		Do_Rrc(gb, &gb.Reg.C)
		gb.Reg.PC-- // undo
	case 0x10:
		Cb_rl(gb, &gb.Reg.B, gb.Reg.B)
	case 0x11:
		Cb_rl(gb, &gb.Reg.C, gb.Reg.C)
	case 0x12:
		Cb_rl(gb, &gb.Reg.D, gb.Reg.D)
	case 0x13:
		Cb_rl(gb, &gb.Reg.E, gb.Reg.E)
	case 0x14:
		Cb_rl(gb, &gb.Reg.H, gb.Reg.H)
	case 0x15:
		Cb_rl(gb, &gb.Reg.L, gb.Reg.L)
	case 0x19:
		Cb_rr(gb, &gb.Reg.C, gb.Reg.C)
	case 0x1A:
		Cb_rr(gb, &gb.Reg.D, gb.Reg.D)
	case 0x1B:
		Cb_rr(gb, &gb.Reg.E, gb.Reg.E)
	case 0x1C:
		Cb_rr(gb, &gb.Reg.H, gb.Reg.H)
	case 0x1D:
		Cb_rr(gb, &gb.Reg.L, gb.Reg.L)
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
	case 0x33:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.E)
	case 0x34:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.H)
	case 0x35:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.L)
	case 0x37:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.A)
	case 0x38:
		Cb_srl(gb, &gb.Reg.B, gb.Reg.B)
	case 0x39:
		Cb_srl(gb, &gb.Reg.C, gb.Reg.C)
	case 0x3A:
		Cb_srl(gb, &gb.Reg.D, gb.Reg.D)
	case 0x3B:
		Cb_srl(gb, &gb.Reg.E, gb.Reg.E)
	case 0x3C:
		Cb_srl(gb, &gb.Reg.H, gb.Reg.H)
	case 0x3D:
		Cb_srl(gb, &gb.Reg.L, gb.Reg.L)
	case 0x3F:
		Cb_srl(gb, &gb.Reg.A, gb.Reg.A)
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
	case 0x87:
		Cb_ResetBit(0x00, &gb.Reg.A)
	case 0x86:
		Cb_res_16(gb, 0x00, gb.Reg.H, gb.Reg.L)
	case 0x9E:
		Cb_res_16(gb, 0x03, gb.Reg.H, gb.Reg.L)
	case 0xBE:
		Cb_res_16(gb, 0x07, gb.Reg.H, gb.Reg.L)
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

func Cb_SwapReg(flags *uint8, target *uint8) {
	*target = ((*target & 0xF) << 4) | ((*target & 0xF0) >> 4)

	*flags &= ^FlagMask // clear all bits

	if *target == 0 {
		*flags |= FlagZ
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

func Cb_srl(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask

	if (value & 0x01) != 0 {
		gb.Reg.F |= FlagC
	}

	value >>= 1

	if value == 0 {
		gb.Reg.F |= FlagZ
	}

	*target = value
}

func Cb_rl(gb *GbCpu, target *uint8, value uint8) {
	oldMask := gb.Reg.F
	gb.Reg.F &= ^FlagMask

	carry := uint8(0)
	if oldMask&FlagC != 0 {
		carry = 1
	}

	if value&0x80 != 0 {
		gb.Reg.F |= FlagC
	}

	value <<= 1
	value += carry

	if value == 0 {
		gb.Reg.F |= FlagZ
	}

	*target = value
}

func Cb_rr(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask
	// FIXME: rr 0xFF on mgba leads to 0xff, we do 0x7f...
	*target >>= 1
	if (gb.Reg.F & FlagH) != 0 {
		*target |= 0x80
	}

	if *target&0x01 != 0 {
		gb.Reg.F |= FlagC
	} else {
		gb.Reg.F &= ^FlagC
	}

	if *target != 0 {
		gb.Reg.F &= ^FlagZ
	} else {
		gb.Reg.F |= FlagZ
	}
}

func Cb_res_16(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := gb.mem.GetByte(addr) & ^(1 << bit)
	gb.mem.WriteByte(addr, val)
}
