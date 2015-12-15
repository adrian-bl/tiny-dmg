package rom

import (
	"io/ioutil"
)

type RomImage struct {
	Blob             []byte // whole file
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

	r.Blob, err = ioutil.ReadFile(path)
	if err == nil {
		r.CGB = (r.Blob[0x143] == 1)
		r.SGB = (r.Blob[0x146] == 1)
		r.Title = string(r.Blob[0x134:0x143]) // fixme: trim junk?
		r.ManufacturerCode = string(r.Blob[0x13f:0x142])
	}
	return
}
