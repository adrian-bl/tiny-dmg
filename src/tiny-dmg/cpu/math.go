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

	if (result & 0x0f) == 0x0f {
		gb.Reg.F |= FlagH
		fmt.Printf("-> HALF CARRY SET.")
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
	*target1 = gb.Mem.GetByte(gb.Reg.PC + 1)
	*target2 = gb.Mem.GetByte(gb.Reg.PC + 2)
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

func Do_Add_88(gb *GbCpu, target *uint8, value uint8) {
	gb.Reg.F &= ^FlagMask

	result := uint16(*target) + uint16(value)
	*target = uint8(result & 0xFF)

	if (result & 0xFF00) != 0 {
		gb.Reg.F |= FlagC
	}
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
	if ((*target & 0x0F) + (value & 0x0F)) > 0x0F {
		gb.Reg.F |= FlagH
		fmt.Printf("Fixme? halfcarry?")
	}

	gb.Reg.PC++
}

func Do_Or_8(gb *GbCpu, target *uint8, value uint8) {
	*target |= value
	gb.Reg.PC++

	gb.Reg.F &= ^FlagMask // clear all bits
	if *target == 0 {
		gb.Reg.F |= FlagZ
	}
}
