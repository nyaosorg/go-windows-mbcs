package mbcs

import (
	"bufio"
	"io"
	"unicode/utf8"
)

type Filter struct {
	sc       *bufio.Scanner
	text     string
	ansi     bool
	err      error
	codepage uintptr
}

func NewFilter(r io.Reader, codepage uintptr) *Filter {
	return &Filter{
		sc:       bufio.NewScanner(r),
		codepage: codepage,
	}
}

func (f *Filter) Scan() bool {
	if !f.sc.Scan() {
		f.err = f.sc.Err()
		return false
	}
	line := f.sc.Bytes()
	if !f.ansi && utf8.Valid(line) {
		f.text = f.sc.Text()
	} else {
		f.text, f.err = AtoU(line, f.codepage)
		if f.err != nil {
			return false
		}
		f.ansi = true
	}
	return true
}
func (f *Filter) Text() string {
	return f.text
}

func (f *Filter) Err() error {
	return f.err
}
