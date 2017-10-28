package main

import (
	"fmt"
	"tiny-dmg/cpu"
	"tiny-dmg/joypad"
	"tiny-dmg/lcd"
	"tiny-dmg/machine"
	"tiny-dmg/memory"
	"tiny-dmg/rom"
	"tiny-dmg/ui"
)

func main() {

	j, err := joypad.New()
	if err != nil {
		panic(err)
	}

	r, err := rom.NewFromDisk("/tmp/hello-world.gb")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded rom with name: *%s*\n", r.Title)

	m, err := memory.New(r, j)
	if err != nil {
		panic(err)
	}

	l, err := lcd.New(m)
	if err != nil {
		panic(err)
	}

	c, err := cpu.New(m, l)
	if err != nil {
		panic(err)
	}

	// Launch our simple UI
	go ui.Run(m, j)

	mach, err := machine.New(c, l, m)
	if err != nil {
		panic(err)
	}

	mach.PowerOn()
	mach.Run()
}
