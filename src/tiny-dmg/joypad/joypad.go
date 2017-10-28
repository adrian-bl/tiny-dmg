package joypad


type Joypad struct {
	KeyUp bool
	KeyDown bool
	KeyLeft bool
	KeyRight bool
	ButtonA bool
	ButtonB bool
	ButtonStart bool
	ButtonSelect bool
}


func New() (*Joypad, error)  {
	return new(Joypad), nil
}

// GetJoypadByte returns how 0xFF00 should look like
// given the current keypad situation
func (j *Joypad) GetJoypadByte(b uint8) uint8 {
	b |= 0x0F // Set all 4 bits to 1 = not pressed

	if b & 0x20 != 0 {
		if j.KeyRight {
			b &^= 1<<0
		}
		if j.KeyLeft {
			b &^= 1<<1
		}
		if j.KeyUp {
			b &^= 1<<2
		}
		if j.KeyDown {
			b &^= 1<<3
		}
	}

	if b & 0x10 != 0 {
		if j.ButtonA {
			b &^= 1<<0
		}
		if j.ButtonB {
			b &^= 1<<1
		}
		if j.ButtonSelect {
			b &^= 1<<2
		}
		if j.ButtonStart {
			b &^= 1<<3
		}
	}

	return b
}
