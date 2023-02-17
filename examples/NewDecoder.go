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
	sc := bufio.NewScanner(transform.NewReader(os.Stdin, mbcs.NewDecoder(mbcs.ACP)))
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
