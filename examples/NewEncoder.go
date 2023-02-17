//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func main() {
	sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.NewEncoder(mbcs.ACP)))
	for sc.Scan() {
		os.Stdout.Write(sc.Bytes())
		os.Stdout.Write([]byte{'\n'})
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
