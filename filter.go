package mbcs

import (
	"io"
)

var _BOM = []byte{0xEF, 0xBB, 0xBF}

// ForceGuessAlways should be called when you should guess
// the encoding for each line repeadedly
func (f *Filter) ForceGuessAlways() {
	f.forceGuessAlways()
}

// NewFilter is the constructor for Filter
func NewFilter(r io.Reader, codepage uintptr) *Filter {
	return newFilter(r, codepage)
}

// Scan is like bufio.Scanner.Scan for Filter
func (f *Filter) Scan() bool {
	return f.scan()
}

// Text is like "bufio".Scanner.Text for Filter
func (f *Filter) Text() string {
	return f.text
}

// Err is like "bufio".Scanner.Err for Filter
func (f *Filter) Err() error {
	return f.err
}
