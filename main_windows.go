package mbcs

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var kernel32 = windows.NewLazySystemDLL("kernel32")
var multiByteToWideChar = kernel32.NewProc("MultiByteToWideChar")
var wideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")
var getConsoleCP = kernel32.NewProc("GetConsoleCP")

func ansiToUtf8(mbcs []byte, codepage uintptr) (string, error) {
	if len(mbcs) == 0 {
		return "", nil
	}
	size, _, err := multiByteToWideChar.Call(
		codepage, 0,
		uintptr(unsafe.Pointer(&mbcs[0])),
		uintptr(len(mbcs)),
		uintptr(0), 0)
	if size == 0 {
		return "", err
	}
	utf16 := make([]uint16, size)
	rc, _, err := multiByteToWideChar.Call(
		codepage, 0,
		uintptr(unsafe.Pointer(&mbcs[0])), uintptr(len(mbcs)),
		uintptr(unsafe.Pointer(&utf16[0])), size)
	if rc == 0 {
		return "", err
	}
	return windows.UTF16ToString(utf16), nil
}

func utf8ToAnsi(utf8 string, codepage uintptr) ([]byte, error) {
	utf16, err := windows.UTF16FromString(utf8)
	if err != nil {
		return nil, err
	}
	size, _, err := wideCharToMultiByte.Call(
		codepage, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(0), 0, uintptr(0), uintptr(0))
	if size == 0 {
		return nil, err
	}
	mbcs := make([]byte, size)
	rc, _, err := wideCharToMultiByte.Call(
		codepage, 0,
		uintptr(unsafe.Pointer(&utf16[0])),
		uintptr(len(utf16)),
		uintptr(unsafe.Pointer(&mbcs[0])), size, uintptr(0), uintptr(0))
	if rc == 0 {
		return nil, err
	}
	if mbcs[size-1] == 0 {
		mbcs = mbcs[:size-1]
	}
	return mbcs, nil
}

func consoleCP() uintptr {
	cp, _, _ := getConsoleCP.Call()
	return cp
}
