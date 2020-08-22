package ui

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image/color"
	"log"
	"runtime"
	"time"

	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/joypad"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/lcd"
	"github.com/adrian-bl/tiny-dmg/lib/tiny-dmg/memory"
)

const (
	visibleWidth  = 160
	visibleHeight = 144
	width         = 256
	height        = 256
	shiftWidth    = (width - visibleWidth) / 2
	shiftHeight   = (height - visibleHeight) / 2
)

func Run(m *memory.Memory, j *joypad.Joypad) {
	//go loop(m)
	go mapView(m, 0x9800, j)
	go mapView(m, 0x9C00, j)
	go tileView(m)
	go sprites(m)
	wde.Run()
	log.Panic("wde run exited!")
}

// mapView displays a tilemap
func mapView(m *memory.Memory, memoff int, j *joypad.Joypad) {
	dw, err := wde.NewWindow(width, height)
	if err != nil {
		fmt.Println(err)
		return
	}
	dw.SetTitle(fmt.Sprintf("MAP:%X - tinydmg!", memoff))
	dw.SetSize(width, height)
	dw.Show()

	go captureInput(dw, j)

	s := dw.Screen()
	for {
		yoff := int(m.GetByte(memory.RegScrollY))
		xoff := int(m.GetByte(memory.RegScrollX))
		lcdc := m.GetByte(memory.RegLcdControl)

		basePointer := uint16(0x9000)

		tileDataOne := lcdc&(lcd.FlagLcdcBgWindowTileSelect) != 0
		if tileDataOne {
			basePointer = 0x8000
		}

		for i := 0; i < 1024; i++ {
			x := i % 32
			y := i / 32

			taddr := basePointer
			baddr := uint8(m.GetByte(uint16(i + memoff)))
			if tileDataOne {
				taddr += uint16(baddr) * 16
			} else {
				taddr += uint16(int8(baddr)) * 16
			}

			drawTile(s, m, taddr, shiftWidth+x*8-xoff, shiftHeight+y*8-yoff)
		}

		doSprites(s, m, shiftWidth, shiftHeight)

		markPosition(s, shiftWidth, shiftHeight)
		markPosition(s, shiftWidth+visibleWidth, shiftHeight+visibleHeight)
		dw.FlushImage()
		time.Sleep(10 * time.Millisecond)
	}
}

func doSprites(s wde.Image, m *memory.Memory, xoff, yoff int) {
	base := uint16(0xFE00)
	for i := uint16(0); i < 40; i++ {
		py := int(m.GetByte(base+i*4+0)) - 16
		px := int(m.GetByte(base+i*4+1)) - 8
		tn := int(m.GetByte(base + i*4 + 2))
		// at := m.GetByte(base + i*4 + 3)
		p := uint16(tn*0x10 + 0x8000)
		drawTile(s, m, p, xoff+px, yoff+py)
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
	for {
		doSprites(s, m, 0, 0)
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

func captureInput(dw wde.Window, j *joypad.Joypad) {
	events := dw.EventChan()

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
			break
		case wde.ResizeEvent:
			fmt.Println("resize", e.Width, e.Height)
		}
	}
}

func drawTile(s wde.Image, m *memory.Memory, taddr uint16, x, y int) {
	for r := 0; r < 8; r++ {
		da := m.GetByte(uint16(r) + taddr)
		taddr++
		db := m.GetByte(uint16(r) + taddr)

		set(s, x+0, y+r, colorize(da>>7&0x01+(db>>7&0x01)<<1))
		set(s, x+1, y+r, colorize(da>>6&0x01+(db>>6&0x01)<<1))
		set(s, x+2, y+r, colorize(da>>5&0x01+(db>>5&0x01)<<1))
		set(s, x+3, y+r, colorize(da>>4&0x01+(db>>4&0x01)<<1))
		set(s, x+4, y+r, colorize(da>>3&0x01+(db>>3&0x01)<<1))
		set(s, x+5, y+r, colorize(da>>2&0x01+(db>>2&0x01)<<1))
		set(s, x+6, y+r, colorize(da>>1&0x01+(db>>1&0x01)<<1))
		set(s, x+7, y+r, colorize(da>>0&0x01+(db>>0&0x01)<<1))
	}
}

func set(s wde.Image, x, y int, color *color.RGBA) {
	if x > width {
		x -= width
	} else if x < 0 {
		x += width
	}

	if y > height {
		y -= height
	} else if y < 0 {
		y += height
	}

	s.Set(x, y, color)
}

func markPosition(s wde.Image, x, y int) {
	len := 200
	for i := -1 * len; i < len; i++ {
		set(s, x, y+i, &color.RGBA{0x00, 0xFF, 0x80, 0xFF})
	}
	for i := -1 * len; i < len; i++ {
		set(s, x+i, y, &color.RGBA{0x00, 0xFF, 0x00, 0xFF})
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
