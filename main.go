package main

import (
	"fmt"
	"tiny-dmg/cpu"
	"tiny-dmg/lcd"
	"tiny-dmg/memory"
	"tiny-dmg/rom"
	"tiny-dmg/ui"
)

func main() {

	r, err := rom.NewFromDisk("/tmp/hello-world.gb")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded rom with name: *%s*\n", r.Title)

	m, err := memory.New(r)
	if err != nil {
		panic(err)
	}

	l, err := lcd.New(m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Assembling machine...")
	gb, err := cpu.New(m, l)
	if err != nil {
		panic(err)
	}

	go ui.Run()

	fmt.Printf("done! machine dump: %v\n", gb)
	gb.Mem.PowerOn()
	gb.Lcd.PowerOn()
	gb.Boot()

}
