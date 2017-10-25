package cpu

import (
	"fmt"
)

func Cb_Disp(gb *GbCpu) {
	op := gb.Mem.GetByte(gb.Reg.PC + 1)
	fmt.Printf("-> CB %02X\n", op)

	switch op {
	case 0x11:
		Cb_rl(gb, &gb.Reg.C, gb.Reg.C)
	case 0x27:
		Cb_sla(gb, &gb.Reg.A)
	case 0x37:
		Cb_SwapReg(&gb.Reg.F, &gb.Reg.A)
	case 0x3F:
		Cb_srl(gb, &gb.Reg.A, gb.Reg.A)
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
	case 0x7F:
		Cb_checkBit(gb, 0x07, gb.Reg.A)
	case 0x87:
		Cb_ResetBit(0x00, &gb.Reg.A)
	case 0xDE:
		Cb_SetHlpBit(gb, 0x03)
	case 0xBF:
		Cb_ResetBit(0x07, &gb.Reg.A)
	default:
		fmt.Printf("Unknown CB opcode: %02X", op)
		for {
		}
	}
	gb.Reg.PC += 2
}

func Cb_ResetBit(bit uint8, target *uint8) {
	*target &= ^(1 << bit)
}

func Cb_SetBit(bit uint8, target *uint8) {
	*target |= (1 << bit)
}

func Cb_SetHlpBit(gb *GbCpu, bit uint8) {
	addr := uint16(gb.Reg.H)<<8 + uint16(gb.Reg.L)
	val := gb.Mem.GetByte(addr)
	Cb_SetBit(bit, &val)
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
	if oldMask & FlagC != 0 {
		carry = 1
	}

	if value & 0x80 != 0 {
		gb.Reg.F |= FlagC
	}

	value <<= 1;
	value += carry;

	if value == 0 {
		gb.Reg.F |= FlagZ
	}

	*target = value
}
