package mbcs_test

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nyaosorg/go-windows-mbcs"
)

// ExampleAnsiToUtf8 converts from ANSI-string of STDIN to UTF8 via STDOUT
func ExampleAnsiToUtf8() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		text, err := mbcs.AnsiToUtf8(sc.Bytes(), mbcs.ACP)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Println(text)
	}
}

// ExampleUtf8ToAnsi converts from UTF8-string of STDIN to ANSI via STDOUT
func ExampleUtf8ToAnsi() {
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
}
