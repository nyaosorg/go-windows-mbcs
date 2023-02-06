package mbcs_test

import (
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
	outputString, err := mbcs.AtoU(outputBytes, actual)
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
