package ui

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image/color"
	"log"
	"runtime"
	"sync"
	"time"
	"tiny-dmg/joypad"
	"tiny-dmg/memory"
)

const (
	width  = 160
	height = 144
)

func Run(m *memory.Memory, j *joypad.Joypad) {
	//go loop(m)
	go mapView(m, 0x9800)
	go mapView(m, 0x9C00)
	go tileView(m)
	go sprites(m)
	go loop(j)
	wde.Run()
	log.Panic("wde run exited!")
}

// mapView displays a tilemap
func mapView(m *memory.Memory, memoff int) {
	dw, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	dw.SetTitle(fmt.Sprintf("MAP:%X - tinydmg!", memoff))
	dw.SetSize(width, height)
	dw.Show()

	s := dw.Screen()
	for {
		yoff := int(m.GetByte(memory.RegScrollY))
		xoff := int(m.GetByte(memory.RegScrollX))
		for i := 0; i < 1024; i++ {
			taddr := 0x8000 + uint16(m.GetByte(uint16(i+memoff)))*16
			x := i % 32
			y := i / 32
			drawTile(s, m, taddr, x*8-xoff, y*8-yoff)
		}
		dw.FlushImage()
		time.Sleep(10 * time.Millisecond)
	}
}

func sprites(m *memory.Memory) {
	dw, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	dw.SetTitle("tiny-dmg! [sprites]")
	dw.SetSize(width, height)
	dw.Show()

	s := dw.Screen()
	base := uint16(0xFE00)
	for {
		for i := uint16(0); i < 40; i++ {
			py := int(m.GetByte(base+i*4+0)) - 16
			px := int(m.GetByte(base+i*4+1)) - 8
			tn := int(m.GetByte(base + i*4 + 2))
			// at := m.GetByte(base + i*4 + 3)
			p := uint16(tn*0x10 + 0x8000)
			drawTile(s, m, p, px, py)
		}
		dw.FlushImage()
		time.Sleep(10 * time.Millisecond)
	}
}

// tileView dumps all tiles
func tileView(m *memory.Memory) {
	dw, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	dw.SetTitle("tiny-dmg! [tiles]")
	dw.SetSize(width, height)
	dw.Show()

	s := dw.Screen()
	// vram is $8000-$97FF, each tile is 8x8 -> 16 bytes
	for {
		for i := 0; i < 360; i++ {
			p := uint16(i*0x10 + 0x8000)
			x := (i % 20) * 8
			y := (i / 20) * 8
			drawTile(s, m, p, x, y)
		}
		dw.FlushImage()
		time.Sleep(10 * time.Millisecond)
	}
}

func loop(j *joypad.Joypad) {
	var wg sync.WaitGroup

	x := func() {
		dw, err := wde.NewWindow(width, height)
		if err != nil {
			fmt.Println(err)
			return
		}
		dw.SetTitle("event input field")
		dw.SetSize(width, height)
		dw.Show()

		events := dw.EventChan()

		done := make(chan bool)

		go func() {
		loop:
			for ei := range events {
				runtime.Gosched()
				switch e := ei.(type) {
				case wde.KeyDownEvent:
					switch e.Key {
					case "right_arrow":
						j.KeyRight = true
					case "left_arrow":
						j.KeyLeft = true
					case "up_arrow":
						j.KeyUp = true
					case "down_arrow":
						j.KeyDown = true
					case "return":
						j.ButtonStart = true
					case "a":
						j.ButtonA = true
					case "s":
						j.ButtonB = true
					case "space":
						j.ButtonSelect = true
					}
				case wde.KeyUpEvent:
					switch e.Key {
					case "right_arrow":
						j.KeyRight = false
					case "left_arrow":
						j.KeyLeft = false
					case "up_arrow":
						j.KeyUp = false
					case "down_arrow":
						j.KeyDown = false
					case "return":
						j.ButtonStart = false
					case "a":
						j.ButtonA = false
					case "s":
						j.ButtonB = false
					case "space":
						j.ButtonSelect = false
					}
				case wde.CloseEvent:
					fmt.Println("close")
					dw.Close()
					break loop
				case wde.ResizeEvent:
					fmt.Println("resize", e.Width, e.Height)
				}
			}
			done <- true
			fmt.Println("end of events")
		}()

	}
	wg.Add(1)
	go x()

	wg.Wait()
	wde.Stop()
}

func drawTile(s wde.Image, m *memory.Memory, taddr uint16, x, y int) {
	for r := 0; r < 8; r++ {
		da := m.GetByte(uint16(r) + taddr)
		taddr++
		db := m.GetByte(uint16(r) + taddr)

		s.Set(x+0, y+r, colorize(da>>7&0x01+(db>>7&0x01)<<1))
		s.Set(x+1, y+r, colorize(da>>6&0x01+(db>>6&0x01)<<1))
		s.Set(x+2, y+r, colorize(da>>5&0x01+(db>>5&0x01)<<1))
		s.Set(x+3, y+r, colorize(da>>4&0x01+(db>>4&0x01)<<1))
		s.Set(x+4, y+r, colorize(da>>3&0x01+(db>>3&0x01)<<1))
		s.Set(x+5, y+r, colorize(da>>2&0x01+(db>>2&0x01)<<1))
		s.Set(x+6, y+r, colorize(da>>1&0x01+(db>>1&0x01)<<1))
		s.Set(x+7, y+r, colorize(da>>0&0x01+(db>>0&0x01)<<1))
	}
}

func colorize(n byte) *color.RGBA {
	switch n {
	case 0:
		return &color.RGBA{0xFA, 0xFA, 0xFA, 0xFF}
	case 1:
		return &color.RGBA{0xF0, 0xA0, 0xAF, 0xFF}
	case 2:
		return &color.RGBA{0xF5, 0x45, 0x45, 0xFF}
	default:
		return &color.RGBA{0xF0, 0x10, 0x10, 0xFF}
	}
}
