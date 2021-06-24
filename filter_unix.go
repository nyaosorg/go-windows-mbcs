//+build !windows

package mbcs

import (
	"bufio"
	"io"
)

type Filter = bufio.Scanner

func NewFilter(r io.Reader, _ uintptr) *Filter {
	return bufio.NewScanner(r)
}
