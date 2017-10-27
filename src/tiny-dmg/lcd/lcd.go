package lcd

import (
	"fmt"
	"tiny-dmg/memory"
)

const (
	FirstScanLine        = 0
	FirstVisibleScanline = 8
	LastScanLine         = 153
	LastVisibleScanline  = 144
)

const (
	CyclesPerScanline  = 456
	CyclesHblank       = 204
	CyclesSrchSprites  = 80
	CyclesTransToLCD   = 172
)

const (
	FlagLcdcBgDisplay          = 1 << 0
	FlagLcdcObjEnable          = 1 << 1
	FlagLcdcObjSize            = 1 << 2
	FlagLcdcBgTileMap          = 1 << 3
	FlagLcdcBgWindowTileSelect = 1 << 4
	FlagLcdcWindowEnable       = 1 << 5
	FlagLcdcWindowTileEnable   = 1 << 6
	FlagLcdcEnable             = 1 << 7
)

const (
	GpuModeHblank      = 0
	GpuModeVblank      = 1
	GpuModeSrchSprites = 2 // oam
	GpuModeTransToLCD  = 3 // vram
)

type Lcd struct {
	m             *memory.Memory
	cyclesCounter int16
}

func New(m *memory.Memory) (l *Lcd, err error) {
	l = new(Lcd)
	l.m = m
	return
}

func (l *Lcd) PowerOn() {
	l.m.WriteByte(memory.RegLcdControl, FlagLcdcBgDisplay|FlagLcdcBgWindowTileSelect|FlagLcdcEnable)
	fmt.Printf("# turning LCD on via %08X -> 0x%02X \n", memory.RegLcdControl, l.m.GetByte(memory.RegLcdControl))

	l.cyclesCounter = CyclesSrchSprites // gnuboy sets this to 40? (we use double cycles)
	l.m.WriteByte(memory.RegCurrentScanline, 0x00)

}

func (l *Lcd) Update(opCycles uint8) {
	// FIXME: Should check if LCD is enabled and reinit + return if not!

	/* Was GNUboy macht:
	 * Init mit       STAT=0
	 * Bei 80 cycles: STAT=2, currentScanline=1
	 * Bei 172 cycle: STAT=3, immer noch cs=1
	 * Bei 344 cycle: STAT=0,
	 * Bei 548        STAT=2, currentScanline++
	 * Bei 632        STAT=3
	 */

	//	fmt.Printf("IST: cpu.lcdc=%d, cnt=%d, --> NOW=%d\n", l.cyclesCounter, opCycles, l.cyclesCounter-int16(opCycles))

	l.cyclesCounter += int16(opCycles)

	state := l.m.GetByte(memory.RegLcdState) & 0xF
	switch state {
	case GpuModeHblank:
		if l.cyclesCounter >= CyclesHblank {
			l.cyclesCounter = 0
			thisScanline := l.m.GetByte(memory.RegCurrentScanline)
			thisScanline++
			l.m.WriteRaw(memory.RegCurrentScanline, thisScanline)
			if thisScanline == LastVisibleScanline-1 {
				state = GpuModeVblank
			} else {
				state = GpuModeSrchSprites
			}
		}
	case GpuModeVblank:
		if l.cyclesCounter >= CyclesPerScanline {
			thisScanline := l.m.GetByte(memory.RegCurrentScanline)
			thisScanline++
			state = GpuModeHblank
			if thisScanline > LastScanLine {
				//???
				thisScanline = 0
				state = GpuModeHblank
			}
			l.m.WriteRaw(memory.RegCurrentScanline, thisScanline)
		}
	case GpuModeSrchSprites:
		if l.cyclesCounter >= CyclesSrchSprites {
			l.cyclesCounter = 0
			state = GpuModeTransToLCD
		}
	case GpuModeTransToLCD:
		{
			if l.cyclesCounter >= CyclesTransToLCD {
				l.cyclesCounter = 0
				state = GpuModeHblank
				fmt.Printf(">>> RENDER SCANLINE?\n")
			}
		}
	default:
		panic(nil)
	}
	l.m.WriteRaw(memory.RegLcdState, state)
	l.m.Dump()
	/*

		mode := uint8(0)
		if (l.m.GetByte(RegCurrentScanline) >= LastVisibleScanline) {
			mode = GpuModeVblank
		} else if (l.cyclesCounter >= CyclesPerScanline-CyclesSrchSprites) {
			mode = GpuModeSrchSprites
		} else if (l.cyclesCounter >= CyclesPerScanline-CyclesSrchSprites-CyclesTransToLCD) {
			mode = GpuModeTransToLCD
		} else {
			mode = GpuModeHblank
		}
		l.m.WriteRaw(RegLcdStat, mode)



		// if display on:
		l.cyclesCounter -= uint16(opCycles)

	*/

	/*
	   	l.scanlineCounter -= int16(opCycles)

	   	if (l.scanlineCounter <= 0) {
	   		l.scanlineCounter = CyclesPerScanline

	   		currentScanline := l.m.GetByte(RegCurrentScanline)
	   		currentScanline++

	   		if (currentScanline == LastVisibleScanline) {
	   			fmt.Printf("LCD: Should request Interupt 0\n")
	   		} else if (currentScanline > LastScanLine) {
	   			fmt.Printf("LCD: Resetting current scanline to ZERO\n")
	   			currentScanline = 0
	   		} else if (currentScanline < LastVisibleScanline) {
	   			fmt.Printf("LCD Could DRAW at %d\n", currentScanline)
	   		}
	   fmt.Printf("LCD WRITE:: %X -> %X\n", RegCurrentScanline, currentScanline)
	   		l.m.WriteRaw(RegCurrentScanline, currentScanline)


	   	}
	*/
}
