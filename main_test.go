package mbcs_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestConsoleCP(t *testing.T) {
	if runtime.GOOS != "windows" {
		return
	}

	actual := mbcs.ConsoleCP()

	outputBytes, err := exec.Command("chcp").Output()
	if err != nil {
		t.Fatalf("chcp: %s", err.Error())
		return
	}
	outputString, err := mbcs.AnsiToUtf8(outputBytes, actual)
	if err != nil {
		t.Fatalf("mbcs.AtoU: %s", err.Error())
		return
	}
	fields := strings.Fields(outputString)

	if fmt.Sprintf("%d", actual) != fields[len(fields)-1] {
		t.Fatalf("ConsoleCP()==%d, chcp->%s", actual, outputString)
	}
}

func readJugemu(t *testing.T) (cp932, utf8 []byte) {
	var err error
	cp932, err = os.ReadFile("testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	utf8, err = os.ReadFile("testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	return cp932, utf8
}

func TestAnsiToUtf8(t *testing.T) {
	if runtime.GOOS != "windows" {
		println("not windows")
		return
	}
	cp932, utf8 := readJugemu(t)
	actual, err := mbcs.AnsiToUtf8(cp932, 932)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(utf8, []byte(actual)) {
		t.Fatalf("not equal\n%#v\n%#v", utf8, []byte(actual))
	}
}

func TestUtf8ToAnsi(t *testing.T) {
	if runtime.GOOS != "windows" {
		println("not windows")
		return
	}
	cp932, utf8 := readJugemu(t)
	actual, err := mbcs.Utf8ToAnsi(string(utf8), 932)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(cp932, actual) {
		t.Fatalf("not equal\n%#v\n%#v", cp932, actual)
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
