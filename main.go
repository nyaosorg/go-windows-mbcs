package mbcs

import (
	"io"

	"github.com/nyaosorg/go-windows-mbcs/filter"
	"github.com/nyaosorg/go-windows-mbcs/internal/core"
)

// THREAD_ACP is the constant meaning the active codepage for thread
const THREAD_ACP = 3

// ACP is the constant meaning the active codepage for OS
const ACP = core.ACP

// ErrUnsupportedOs is return value when AtoU,UtoA is called on not Windows
var ErrUnsupportedOs = core.ErrUnsupportedOs

// Deprecated: use AnsiToUtf8
func AtoU(ansi []byte, codepage uintptr) (string, error) {
	return core.Atou(ansi, codepage)
}

// AnsiToUtf8 Converts Ansi-bytes to UTF8-String
func AnsiToUtf8(ansi []byte, codepage uintptr) (utf8 string, err error) {
	return core.Atou(ansi, codepage)
}

// Deprecated: use Utf8ToAnsi
func UtoA(utf8 string, codepage uintptr) (ansi []byte, err error) {
	return core.Utoa(utf8, codepage)
}

// Utf8ToAnsi Converts UTF8-String to Ansi-bytes
func Utf8ToAnsi(utf8 string, codepage uintptr) (ansi []byte, err error) {
	return core.Utoa(utf8, codepage)
}

// ConsoleCP returns Codepage number of Console.
func ConsoleCP() uintptr {
	return core.ConsoleCP()
}

// Deprecated: use "github.com/nyaosrorg/go-windows-mbcs/filter".New() instead
func NewFilter(r io.Reader, codepage uintptr) *mbcsfilter.Filter {
	return mbcsfilter.New(r, codepage)
}
