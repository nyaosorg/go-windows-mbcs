package mbcs_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/nyaosorg/go-windows-mbcs"
)

func loadTestdata(t *testing.T) ([]byte, string) {
	ansi, err := os.ReadFile("testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	utf8, err := os.ReadFile("testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	return ansi, string(utf8)
}

func TestAnsiToUtf8(t *testing.T) {
	ansi, utf8expect := loadTestdata(t)
	utf8result, err := mbcs.AnsiToUtf8(ansi, 932)
	if err != nil {
		t.Fatal(err.Error())
	}
	if utf8expect != utf8result {
		t.Fatalf("expect %#v but %#v", utf8expect, utf8result)
	}
}

func TestUtf8ToAnsi(t *testing.T) {
	ansiExpect, utf8 := loadTestdata(t)
	ansiResult, err := mbcs.Utf8ToAnsi(utf8, 932)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(ansiExpect, ansiResult) {
		t.Fatalf("expect %#v but %#v", ansiExpect, ansiResult)
	}
}

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
