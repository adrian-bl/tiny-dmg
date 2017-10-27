package cpu

import (
	"fmt"
)

func Cb_Disp(gb *GbCpu) {
	op := gb.Mem.GetByte(gb.Reg.PC + 1)
	fmt.Printf("-> CB %02X\n", op)

	switch op {
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
	case 0x27:
		Cb_sla(gb, &gb.Reg.A)
	case 0x33:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.E)
	case 0x34:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.H)
	case 0x35:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.L)
	case 0x37:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.A)
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
	case 0xBE:
		Cb_res_16(gb, 0x07, gb.Reg.H, gb.Reg.L)
	case 0xDE:
		Cb_SetHlpBit(gb, 0x03)
	case 0xBF:
		Cb_ResetBit(0x07, &gb.Reg.A)
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
	val := gb.Mem.GetByte(addr)
	Cb_set(bit, &val)
	gb.Mem.WriteByte(addr, val)
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

func Cb_checkBitn(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := uint8(gb.Mem.GetByte(addr))
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

func Cb_res_16(gb *GbCpu, bit, h, l uint8) {
	addr := uint16(h)<<8 + uint16(l)
	val := gb.Mem.GetByte(addr) & ^(1 << bit)
	gb.Mem.WriteByte(addr, val)
}
