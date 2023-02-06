//go:build !windows
// +build !windows

package core

// AtoU Converts Ansi-bytes to UTF8-String
func AnsiToUtf8(mbcs []byte, codepage uintptr) (string, error) {
	return string(mbcs), ErrUnsupportedOs
}

// UtoA Converts UTF8-String to Ansi-bytes
func Utf8ToAnsi(utf8 string, codepage uintptr) ([]byte, error) {
	return []byte(utf8), ErrUnsupportedOs
}

func ConsoleCP() uintptr {
	return ACP
}
