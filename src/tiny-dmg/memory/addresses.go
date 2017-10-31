package memory

const (
	RegJoypadInput           = 0xFF00
	RegSerialTransferData    = 0xFF01
	RegSerialTransferControl = 0xFF02
	RegDivider               = 0xFF04
	RegTimerCounter          = 0xFF05
	RegTimerModulo           = 0xFF06
	RegTimerControl          = 0xFF07
)

// LCD related
const (
	RegLcdControl      = 0xFF40 // LCDC
	RegLcdState        = 0xFF41 // STAT
	RegScrollY         = 0xFF42
	RegScrollX         = 0xFF43
	RegCurrentScanline = 0xFF44 // LY
	RegLYCompare       = 0xFF45
	RegDoDMA           = 0xFF46
)

// Interrupts

const (
	RegInterruptFlag   = 0xFF0F
	RegInterruptEnable = 0xFFFF
	BitIrVblank        = (1 << 0)
	BitIrStat          = (1 << 1)
	BitIrTimer         = (1 << 2)
	BitIrSerial        = (1 << 3)
	BitIrJoypad        = (1 << 4)
)

const (
	StartOamRange = 0xFE00
)
