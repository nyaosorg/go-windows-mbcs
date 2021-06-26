package mbcs

import (
	"bufio"
	"bytes"
	"io"
	"unicode/utf8"
)

type _Status int

const (
	_GuessMinimum _Status = iota
	_GuessAlways
	_AnsiFixed
	_Utf8Fixed
)

type Filter struct {
	sc       *bufio.Scanner
	text     string
	status   _Status
	err      error
	codepage uintptr
}

func (f *Filter) ForceGuessAlways() {
	f.status = _GuessAlways
}

func NewFilter(r io.Reader, codepage uintptr) *Filter {
	return &Filter{
		sc:       bufio.NewScanner(r),
		codepage: codepage,
	}
}

var _BOM = []byte{0xEF, 0xBB, 0xBF}

func (f *Filter) Scan() bool {
	if !f.sc.Scan() {
		f.err = f.sc.Err()
		return false
	}
	line := f.sc.Bytes()
	if f.status == _Utf8Fixed {
		f.text = f.sc.Text()
		return true
	}
	if f.status != _AnsiFixed {
		if bytes.HasPrefix(line, _BOM) {
			f.text = string(line[len(_BOM):])
			if f.status != _GuessAlways {
				f.status = _Utf8Fixed
			}
			return true
		}
		if utf8.Valid(line) {
			f.text = f.sc.Text()
			// f.status should not be fixed yet,
			// because it does not work expected
			// when the first line uses ascii-code only
			// and the second line uses CP932.
			return true
		}
	}
	f.text, f.err = AtoU(line, f.codepage)
	if f.err != nil {
		return false
	}
	if f.status != _GuessAlways {
		f.status = _AnsiFixed
	}
	return true
}
func (f *Filter) Text() string {
	return f.text
}

func (f *Filter) Err() error {
	return f.err
}
