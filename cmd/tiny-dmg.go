package main

import (
	"flag"
	"log"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/cpu"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/joypad"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/lcd"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/machine"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/memory"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/rom"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/ui"
)

var (
	flagBios = flag.String("bios", "", "path to DMG boot bios")
	flagRom  = flag.String("rom", "", "path to DMG rom")
)

func main() {
	flag.Parse()

	j, err := joypad.New()
	if err != nil {
		panic(err)
	}

	var b *rom.RomImage
	if *flagBios != "" {
		log.Printf("Loading bios from %s\n", *flagBios)
		if b, err = rom.NewFromDisk(*flagBios); err != nil {
			log.Fatalf("failed to read bios: %v\n", err)
		}
	} else {
		log.Printf("Using builtin bios\n")
		b = rom.NewBuiltinBios()
	}

	var r *rom.RomImage
	if *flagRom != "" {
		log.Printf("Loading ROM from %s\n", *flagRom)
		if r, err = rom.NewFromDisk(*flagRom); err != nil {
			log.Fatalf("failed to read ROM: %v\n", err)
		}
	} else {
		log.Fatalf("-rom flag must point to a DMG rom\n")
	}
	log.Printf("Loaded ROM with title %s\n", r.Title)

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
