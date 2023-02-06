//go:build !windows
// +build !windows

package core

// AtoU Converts Ansi-bytes to UTF8-String
func Atou(mbcs []byte, codepage uintptr) (string, error) {
	return string(mbcs), ErrUnsupportedOs
}

// UtoA Converts UTF8-String to Ansi-bytes
func Utoa(utf8 string, codepage uintptr) ([]byte, error) {
	return []byte(utf8), ErrUnsupportedOs
}

func ConsoleCP() uintptr {
	return ACP
}
