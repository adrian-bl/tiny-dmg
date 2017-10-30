package rom

import (
	"io/ioutil"
)

type RomImage struct {
	blob             []byte // whole file
	LogoValid        bool   // 0x104-0x133
	Title            string // 0x134-0x143
	ManufacturerCode string // 0x013f-0x0142
	CGB              bool   // 0x143
	SGB              bool   // 0x146
	RomType          byte   // 0x147
	RomSize          byte   // 0x148
	RamSize          byte   // 0x149
	RomMaskVersion   byte   // 0x14C
}

func NewFromDisk(path string) (r RomImage, err error) {
	r = RomImage{}

	r.blob, err = ioutil.ReadFile(path)
	if err == nil {
		r.CGB = (r.blob[0x143] == 1)
		r.SGB = (r.blob[0x146] == 1)
		r.RomType = r.blob[0x147]
		r.Title = string(r.blob[0x134:0x143]) // fixme: trim junk?
		r.ManufacturerCode = string(r.blob[0x13f:0x142])
	}
	return
}

// GetBytes returns the full ROM as it was read from disk.
func (r *RomImage) GetBytes() []byte {
	return r.blob[0:]
}
