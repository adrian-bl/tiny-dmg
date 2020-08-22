package main

import (
	"fmt"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/cpu"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/joypad"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/lcd"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/machine"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/memory"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/ui"
)

func main() {

	j, err := joypad.New()
	if err != nil {
		panic(err)
	}

	b, err := rom.NewFromDisk("/tmp/bios.bin")
	if err != nil {
		panic(err)
	}

	r, err := rom.NewFromDisk("/tmp/hello-world.gb")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded rom with name: *%s*\n", r.Title)

	m, err := memory.New(b, r, j)
	if err != nil {
		panic(err)
	}

	l, err := lcd.New(m)
	if err != nil {
		panic(err)
	}

	c, err := cpu.New(m)
	if err != nil {
		panic(err)
	}

	// Launch our simple UI
	go ui.Run(m, j)

	mach, err := machine.New(c, m, l)
	if err != nil {
		panic(err)
	}

	mach.PowerOn()
	mach.Run()
}
