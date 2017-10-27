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
	CyclesPerScanline = CyclesHblank + CyclesSrchSprites + CyclesTransToLCD
	CyclesHblank      = 204 * 2 // Mode0
	CyclesSrchSprites = 80 * 2  // Mode2
	CyclesTransToLCD  = 172 * 2 // Mode3
)

// LCDC FF40 flags
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

// LCD STATUS FF41 status
const (
	BitmaskLcdsGpuMode           = 0x03
	FlagLcdsCoincidenceSignal    = 1 << 2
	FlagLcdsModeZeroHblankEnable = 1 << 3
	FlagLcdsModeOneVblankEnable  = 1 << 4
	FlagLcdsModeTwoOamEnable     = 1 << 5
	FlagLcdsCoincidenceEnable    = 1 << 6
	FlagLcdsUnusedStatus         = 1 << 7
)

const (
	GpuModeHblank     = 0
	GpuModeVblank     = 1
	GpuModeReadOAM    = 2 // oam
	GpuModeTransToLCD = 3 // vram
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
	l.m.WriteRaw(memory.RegLcdControl, FlagLcdcBgDisplay|FlagLcdcBgWindowTileSelect|FlagLcdcEnable)
	fmt.Printf("# turning LCD on via %08X -> 0x%02X \n", memory.RegLcdControl, l.m.GetByte(memory.RegLcdControl))

	//l.cyclesCounter = CyclesSrchSprites
	l.m.WriteRaw(memory.RegCurrentScanline, 0x00)

}

func (l *Lcd) Update(cycles uint8) {
	if (l.m.GetByte(memory.RegLcdControl) & FlagLcdcEnable) == 0 {
		l.cyclesCounter = 0
		l.m.WriteRaw(memory.RegCurrentScanline, 0)
		fmt.Printf("FIXME: Should move lcd status to state 2?\n")
		return
	}

	l.setStatus()

	l.cyclesCounter += int16(cycles)
	if l.cyclesCounter > CyclesPerScanline {
		currentScanline := l.m.GetByte(memory.RegCurrentScanline) + 1
		l.cyclesCounter = 0

		if currentScanline == LastVisibleScanline {
			// FIXME: RequestInterrupt
		}
		if currentScanline > LastScanLine {
			currentScanline = 0
		} else if currentScanline < LastVisibleScanline {
			// draw
		}

		l.m.WriteRaw(memory.RegCurrentScanline, currentScanline)
	}
}

func (l *Lcd) setStatus() {

	status := l.m.GetByte(memory.RegLcdState)

	currentScanline := l.m.GetByte(memory.RegCurrentScanline)
	//	currentMode := status & BitmaskLcdsGpuMode
	newMode := byte(0)

	if currentScanline > LastVisibleScanline {
		newMode = GpuModeVblank
		// FIXME: Test reqInt 4
	} else {
		if l.cyclesCounter < CyclesSrchSprites { // X->2
			newMode = GpuModeReadOAM
			// FIXME: Test reqInt 5
		} else if l.cyclesCounter < CyclesTransToLCD { // 2->3
			newMode = GpuModeTransToLCD
		} else { // 3->0
			newMode = GpuModeHblank
			// FIXME: Test reqInt 3
		}
	}

	// Update status with what we want it to be
	// by setting bit 0 and 1
	status &= 0xFC
	status |= newMode

	//if (reqInt && (newMode != currentMode)) {
	// RequestInterrupt
	//}

	// FIXME: Check coincidence:
	// http://www.codeslinger.co.uk/pages/projects/gameboy/lcd.html
	l.m.WriteRaw(memory.RegLcdState, status)

}
