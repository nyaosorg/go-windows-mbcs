//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nyaosorg/go-windows-mbcs"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		bytes, err := mbcs.Utf8ToAnsi(sc.Text(), mbcs.ACP)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Stdout.Write(bytes)
		os.Stdout.Write([]byte{'\n'})
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
