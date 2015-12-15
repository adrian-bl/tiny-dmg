package cpu

func Cb_ResetBit(bit uint8, target *uint8) {
	*target &= ^(1 << bit)
}
