// +build !windows

package mbcs

import (
	"errors"
)

// AtoU Converts Ansi-bytes to UTF8-String
func atou(mbcs []byte, codepage uintptr) (string, error) {
	return "", errors.New("AtoU: not supported in this OS")
}

// UtoA Converts UTF8-String to Ansi-bytes
func utoa(utf8 string, codepage uintptr) ([]byte, error) {
	return nil, errors.New("UtoA: not supported in this OS")
}
