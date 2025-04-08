package input

import (
	"syscall"
	"unsafe"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	getStdHandle       = kernel32.NewProc("GetStdHandle")
	setConsoleMode     = kernel32.NewProc("SetConsoleMode")
	getConsoleMode     = kernel32.NewProc("GetConsoleMode")
	readConsoleInput   = kernel32.NewProc("ReadConsoleInputW")
)

const (
	stdInputHandle   = uintptr(0xFFFFFFF6)
	enableEchoInput  = 0x0004
	enableLineInput  = 0x0002
	enableProcessed  = 0x0001
	enableWindowInput = 0x0008
	enableMouseInput  = 0x0010

	keyEvent = 0x0001
)

type inputRecord struct {
	EventType uint16
	_         uint16
	KeyEvent  keyEventRecord
}

type keyEventRecord struct {
	KeyDown          int32
	RepeatCount      uint16
	VirtualKeyCode   uint16
	VirtualScanCode  uint16
	UnicodeChar      uint16
	ControlKeyState  uint32
}

var originalMode uint32
var hStdin syscall.Handle

func Init() error {
	h, _, err := getStdHandle.Call(uintptr(uint32(stdInputHandle)))
	if h == 0 {
		return err
	}
	hStdin = syscall.Handle(h)

	_, _, err = getConsoleMode.Call(uintptr(hStdin), uintptr(unsafe.Pointer(&originalMode)))
	if err != nil && err != syscall.Errno(0) {
		return err
	}

	mode := originalMode &^ (enableLineInput | enableEchoInput)
	_, _, err = setConsoleMode.Call(uintptr(hStdin), uintptr(mode))
	if err != nil && err != syscall.Errno(0) {
		return err
	}

	return nil
}

func Restore() {
	setConsoleMode.Call(uintptr(hStdin), uintptr(originalMode))
}

func ReadKey() rune {
	var record inputRecord
	var read uint32
	for {
		readConsoleInput.Call(
			uintptr(hStdin),
			uintptr(unsafe.Pointer(&record)),
			1,
			uintptr(unsafe.Pointer(&read)),
		)

		if record.EventType == keyEvent && record.KeyEvent.KeyDown != 0 {
			r := rune(record.KeyEvent.UnicodeChar)
			if r != 0 {
				return r
			}
		}
	}
}
