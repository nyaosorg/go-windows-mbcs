package mbcs

import (
	"errors"
)

// THREAD_ACP is the constant meaning the active codepage for thread
const THREAD_ACP = 3

// ACP is the constant meaning the active codepage for OS
const ACP = 0

var _BOM = []byte{0xEF, 0xBB, 0xBF}

// ErrUnsupportedOs is return value when AtoU,UtoA is called on not Windows
var ErrUnsupportedOs = errors.New("Unsupported OS")

// AtoU Converts Ansi-bytes to UTF8-String
func AtoU(mbcs []byte, codepage uintptr) (string, error) {
	return atou(mbcs, codepage)
}

// UtoA Converts UTF8-String to Ansi-bytes
func UtoA(utf8 string, codepage uintptr) ([]byte, error) {
	return utoa(utf8, codepage)
}

// ConsoleCP returns Codepage number of Console.
func ConsoleCP() uintptr {
	return consoleCP()
}
