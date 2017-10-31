package lcd

import (
	"fmt"
	"tiny-dmg/memory"
)

const (
	FirstScanLine        = 0
	FirstVisibleScanline = 8
	LastScanLine         = CyclesFullRefresh / CyclesPerScanline
	LastVisibleScanline  = 144
)

// Drawing a full line takes about 456 cycles
const (
	CyclesFullRefresh = 70224 // How many cycles a full screen refresh takes (aka LCDC_PERIOD)
	CyclesPerScanline = CyclesHblank + CyclesOamTransfer + CyclesRendering
	CyclesHblank      = 204 // Mode0
	CyclesOamTransfer = 80  // Mode2
	CyclesRendering   = 172 // Mode3
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
	cyclesCounter uint32
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
	l.cyclesCounter = 66084
	l.m.WriteRaw(memory.RegCurrentScanline, LastVisibleScanline)

}

// Updates the LCD status, that is:
// Zero bit 0 and 1 and set a new mask
func (l *Lcd) updateLcdState(set uint8) {
	state := l.m.GetByte(memory.RegLcdState)
	state &^= BitmaskLcdsGpuMode
	state |= set
	l.m.WriteRaw(memory.RegLcdState, state)
}

// vblankInterrupt gets called by the Update loop when
// a vblank interrupt fired. Sets bit 0 in IF.
func (l *Lcd) vblankInterrupt() {
	l.m.WriteRawSet(memory.RegInterruptFlag, memory.BitIrVblank)
}

func (l *Lcd) Update(cycles uint8) {

	if (l.m.GetByte(memory.RegLcdControl) & FlagLcdcEnable) == 0 {
		// LCD is disabled, values are reset
		l.m.WriteRaw(memory.RegCurrentScanline, 0x00)

		l.updateLcdState(FlagLcdsCoincidenceSignal)

		// Keep counting cycles and fire vblanks
		l.cyclesCounter += uint32(cycles)

		if l.cyclesCounter >= CyclesFullRefresh {
			l.cyclesCounter -= CyclesFullRefresh
			l.vblankInterrupt()
		}
		fmt.Printf("LCD IS STILL OFF...\n")
		return
	}

	// 4 CpuCycles = 1 MachineCycle, so make sure to update the LCD
	// on every machine cycle
	cycleStep := uint8(4)
	statDelay := uint32(4) // ?!
	scxDelay := uint32(4)  // ?!

	for ; cycles != 0; cycles -= cycleStep {
		l.cyclesCounter += uint32(cycleStep)
		l.m.WriteRaw(memory.RegCurrentScanline, uint8(l.cyclesCounter/CyclesPerScanline))

		if l.cyclesCounter == CyclesFullRefresh {
			// Full refresh cycle completed, reset display to initial values.
			l.m.WriteRaw(memory.RegCurrentScanline, 0)
			l.updateLcdState(GpuModeHblank)
			l.cyclesCounter = 0
		} else if l.cyclesCounter == LastVisibleScanline*CyclesPerScanline+statDelay {
			// -> VBLANK state entered
			l.updateLcdState(GpuModeVblank)
			l.vblankInterrupt()
		} else if l.cyclesCounter < LastVisibleScanline*CyclesPerScanline {
			posInScanline := l.cyclesCounter % CyclesPerScanline

			if posInScanline == statDelay {
				l.updateLcdState(GpuModeReadOAM)
			} else if posInScanline == CyclesOamTransfer+statDelay {
				l.updateLcdState(GpuModeTransToLCD)
			} else if posInScanline == CyclesOamTransfer+CyclesRendering+statDelay+scxDelay {
				l.updateLcdState(GpuModeHblank)
			}
		} else if l.cyclesCounter >= ((LastScanLine - 1) * CyclesPerScanline) {
			// >= Line 153 is a special snowflake..
			// This is mostly stolen from SameBoy.
			pp := l.cyclesCounter - (LastScanLine-1)*CyclesPerScanline
			switch pp {
			case 0:
			// should_compare_ly = false
			case 4:
				l.m.WriteRaw(memory.RegCurrentScanline, 0)
			case 8:
				l.m.WriteRaw(memory.RegCurrentScanline, 0)
			default:
				l.m.WriteRaw(memory.RegCurrentScanline, 0)
			}
		} else {
			// 144 - 152
			// if statDelay && l.cyclesCount % CyclesPerScanline == 0
			// ly_for_comparsion--;
		}

	}
}
