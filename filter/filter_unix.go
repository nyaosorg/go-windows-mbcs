//go:build !windows
// +build !windows

package mbcsfilter

import (
	"bufio"
	"io"
	"strings"
)

// Filter is the class like bufio.Scanner but detects the encoding-type
// and converts to utf8 on Windows. On other OSes, it works like bufio.Scanner
type Filter struct {
	sc   *bufio.Scanner
	text string
	err  error
}

func (f *Filter) forceGuessAlways() {
}

func newFilter(r io.Reader, _ uintptr) *Filter {
	return &Filter{
		sc: bufio.NewScanner(r),
	}
}

func (f *Filter) scan() bool {
	if !f.sc.Scan() {
		f.err = f.sc.Err()
		return false
	}
	f.text = f.sc.Text()
	if strings.HasPrefix(f.text, string(_BOM)) {
		f.text = f.text[len(_BOM):]
	}
	return true
}
