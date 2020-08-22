package rom

import (
	"fmt"
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
	IsBios           bool
}

func NewFromDisk(path string) (*RomImage, error) {
	r := &RomImage{}
	var err error

	r.blob, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(r.blob) == 256 {
		r.IsBios = true
		return r, nil
	}

	if len(r.blob) < 0x148 {
		return nil, fmt.Errorf("File too small / not a valid rom")
	}
	r.CGB = (r.blob[0x143] == 1)
	r.SGB = (r.blob[0x146] == 1)
	r.RomType = r.blob[0x147]
	r.Title = string(r.blob[0x134:0x143]) // fixme: trim junk?
	r.ManufacturerCode = string(r.blob[0x13f:0x142])
	return r, nil
}

// GetBytes returns the full ROM as it was read from disk.
func (r *RomImage) GetBytes() []byte {
	return r.blob[0:]
}

// Returns a single byte from the raw rom.
func (r *RomImage) GetByte(addr uint32) byte {
	return r.blob[addr]
}
