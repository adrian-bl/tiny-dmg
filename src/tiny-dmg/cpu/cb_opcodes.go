package cpu

func Cb_ResetBit(bit uint8, target *uint8) {
	*target &= ^(1 << bit)
}

func Cb_SwapReg(flags *uint8, target *uint8) {
	*target = ((*target & 0xF) << 4) | ((*target & 0xF0) >> 4)

	*flags &= ^FlagMask // clear all bits

	if *target == 0 {
		*flags |= FlagZ
	}
}

func Cb_RES(r uint8, target *uint8) {
	*target &= ^(1 << r)
}
