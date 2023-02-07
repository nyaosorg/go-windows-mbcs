//go:build !windows
// +build !windows

package mbcs

// AtoU Converts Ansi-bytes to UTF8-String
func ansiToUtf8(mbcs []byte, codepage uintptr) (string, error) {
	switch codepage {
	case ACP, THREAD_ACP, 65001:
		return string(mbcs), nil
	default:
		return string(mbcs), ErrUnsupportedOs
	}
}

// UtoA Converts UTF8-String to Ansi-bytes
func utf8ToAnsi(utf8 string, codepage uintptr) ([]byte, error) {
	switch codepage {
	case ACP, THREAD_ACP, 65001:
		return []byte(utf8), nil
	default:
		return []byte(utf8), ErrUnsupportedOs
	}
}

func consoleCP() uintptr {
	return ACP
}
