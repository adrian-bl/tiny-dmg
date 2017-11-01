package cpu

import (
	"fmt"
)

func Do_Inc_Uint8(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^(FlagZ | FlagN | FlagH)
	result := *target + 1

	if result == 0 {
		gb.Reg.F |= FlagZ
	}

	if (result & 0x0F) == 0x00 {
		gb.Reg.F |= FlagH
		fmt.Printf("-> HALF CARRY SET\n")
	}

	*target = result
	gb.Reg.PC++
}

func Do_Dec_Uint8(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^(FlagZ | FlagH)
	gb.Reg.F |= FlagN
	result := *target - 1

	if result == 0 {
		gb.Reg.F |= FlagZ
	}
	if (result^0x01^*target)&0x10 == 0x10 {
		gb.Reg.F |= FlagH
	}
	*target = result
	gb.Reg.PC++
}

func Do_Inc_88(gb *GbCpu, target1 *uint8, target2 *uint8) {
	val := uint16(*target1)<<8 + uint16(*target2)
	val++
	*target1 = uint8(val >> 8 & 0xFF)
	*target2 = uint8(val & 0xFF)
	gb.Reg.PC++
}

func Do_Dec_88(gb *GbCpu, target1 *uint8, target2 *uint8) {
	val := uint16(*target1)<<8 + uint16(*target2)
	val--
	*target1 = uint8(val >> 8 & 0xFF)
	*target2 = uint8(val & 0xFF)
	gb.Reg.PC++
}

func Do_Load_88(gb *GbCpu, target1 *uint8, target2 *uint8) {
	*target1 = gb.mem.GetByte(gb.Reg.PC + 1)
	*target2 = gb.mem.GetByte(gb.Reg.PC + 2)
	gb.Reg.PC += 3
}

func Do_And_88(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^(FlagZ | FlagN | FlagC)
	gb.Reg.F |= FlagH

	result := *target & value
	if result == 0x00 {
		gb.Reg.F |= FlagZ
	}
	*target = result
	gb.Reg.PC++
}

func Do_Xor_88(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask

	*target ^= value
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
	gb.Reg.PC++
}

func Do_Add_1616(gb *GbCpu, target *uint16, value uint16) {
	gb.Reg.F &= ^(FlagC | FlagH | FlagN) // Zero flag is not touched

	result := uint32(*target) + uint32(value)

	if (uint32(*target)&0xFFF+uint32(value)&0xFFF)&0x1000 != 0 {
		gb.Reg.F |= FlagH
	}
	if result > 0xFFFF {
		gb.Reg.F |= FlagC
	}
	// No Z flag

	*target = uint16(result)
	gb.Reg.PC++
}

func Do_Add_88(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask

	result := uint16(*target) + uint16(value)

	if (result & 0xFF00) != 0 {
		gb.Reg.F |= FlagC
	}
	if result&0x00FF == 0 {
		gb.Reg.F |= FlagZ
	}
	if uint16(*target)&0xF+uint16(value)&0xF > 0xF {
		gb.Reg.F |= FlagH
	}
	*target = uint8(result)
	gb.Reg.PC++
}

func Do_Or_88(gb *GbCpu, target *uint8, value uint8) {
	*target |= value

	gb.Reg.F &= ^FlagMask // clear all bits
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}

	gb.Reg.PC++
}

func Do_Sub_88(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask

	result := uint16(*target) - uint16(value)
	gb.Reg.F |= FlagN

	if *target < value {
		gb.Reg.F |= FlagC
	}
	if result&0x00FF == 0 {
		gb.Reg.F |= FlagZ
	}
	if (*target & 0xF) < (value & 0xF) {
		gb.Reg.F |= FlagH
	}
	*target = uint8(result)
	gb.Reg.PC++
}

// "The sum of the two operands plus the carry flag (0 or 1) is calculated, and the result is written back into the first operand."
func Do_Adc_88(gb *GbCpu, target *uint8, value uint8) {
	carry := uint16(0)
	if gb.Reg.F&FlagC != 0 {
		carry = 1
	}

	gb.Reg.F &= ^FlagMask

	result := uint16(*target) + uint16(value) + carry

	if (result & 0xFF00) != 0 {
		gb.Reg.F |= FlagC
	}
	if result&0x00FF == 0 {
		gb.Reg.F |= FlagZ
	}
	if uint16(*target)&0xF+uint16(value)&0xF+carry > 0xF {
		gb.Reg.F |= FlagH
	}

	*target = uint8(result)
	gb.Reg.PC++
}

func Do_Sbc_88(gb *GbCpu, target *uint8, value uint8) {
	carry := uint16(0)
	if gb.Reg.F&FlagC != 0 {
		carry = 1
	}

	gb.Reg.F &= ^FlagMask

	result := uint16(*target) - uint16(value) - carry
	gb.Reg.F |= FlagN

	if (result & 0xFF00) != 0 {
		gb.Reg.F |= FlagC
	}
	if result&0x00FF == 0 {
		gb.Reg.F |= FlagZ
	}
	if uint16(*target)&0xF < uint16(value)&0xF+carry {
		gb.Reg.F |= FlagH
	}

	*target = uint8(result)
	gb.Reg.PC++
}

func Do_Cp(gb *GbCpu, a, b uint8) {
	gb.Reg.F &= ^FlagMask

	if a == b {
		gb.Reg.F |= FlagZ
	}
	if b > a {
		gb.Reg.F |= FlagC
	}
	if (b & 0x0F) > (a & 0x0F) {
		gb.Reg.F |= FlagH
	}

	gb.Reg.F |= FlagN
	gb.Reg.PC++
}

func Do_Rrc(gb *GbCpu, target *uint8) {
	gb.Reg.F &= ^FlagMask

	carry := *target & 0x01

	*target >>= 1
	if carry != 0 {
		gb.Reg.F |= FlagC
		*target |= 0x80
	}

	gb.Reg.PC++
}
