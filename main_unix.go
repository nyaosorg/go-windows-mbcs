// +build !windows

package mbcs

// AtoU Converts Ansi-bytes to UTF8-String
func atou(mbcs []byte, codepage uintptr) (string, error) {
	return string(mbcs), ErrUnsupportedOs
}

// UtoA Converts UTF8-String to Ansi-bytes
func utoa(utf8 string, codepage uintptr) ([]byte, error) {
	return []byte(utf8), ErrUnsupportedOs
}

func consoleCP() uintptr {
	return ACP
}
