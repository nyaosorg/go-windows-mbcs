package mbcsfilter

import (
	"bufio"
	"bytes"
	"io"
	"unicode/utf8"

	core "github.com/nyaosorg/go-windows-mbcs/internal/core"
)

type _Status int

const (
	_GuessMinimum _Status = iota
	_GuessAlways
	_AnsiFixed
	_Utf8Fixed
)

// Filter is the class like bufio.Scanner but detects the encoding-type
// and converts to utf8 on Windows. On other OSes, it works like bufio.Scanner
type Filter struct {
	sc       *bufio.Scanner
	text     string
	status   _Status
	err      error
	codepage uintptr
}

func (f *Filter) forceGuessAlways() {
	f.status = _GuessAlways
}

func newFilter(r io.Reader, codepage uintptr) *Filter {
	return &Filter{
		sc:       bufio.NewScanner(r),
		codepage: codepage,
	}
}

func (f *Filter) scan() bool {
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
	f.text, f.err = core.Atou(line, f.codepage)
	if f.err != nil {
		return false
	}
	if f.status != _GuessAlways {
		f.status = _AnsiFixed
	}
	return true
}
