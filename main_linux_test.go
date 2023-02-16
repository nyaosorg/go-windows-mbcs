package mbcs_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestAnsiToUtf8onAcpLinux(t *testing.T) {
	ansi, utf8expect := loadTestdata(t)
	origLang, okLang := os.LookupEnv("LC_ALL")
	os.Setenv("LC_ALL", "C.Shift_JIS")
	//os.Setenv("LANG", "C.EUC-JP")
	defer func() {
		if okLang {
			os.Setenv("LC_ALL", origLang)
		} else {
			os.Unsetenv("LC_ALL")
		}
	}()

	utf8result, err := mbcs.AnsiToUtf8(ansi, mbcs.ACP)
	if err != nil {
		t.Fatal(err.Error())
	}
	if utf8expect != utf8result {
		t.Fatalf("expect %#v but %#v", utf8expect, utf8result)
	}
}

func TestUtf8ToAnsiAcpOnLinux(t *testing.T) {
	ansiExpect, utf8 := loadTestdata(t)
	origLang, okLang := os.LookupEnv("LC_ALL")
	os.Setenv("LC_ALL", "C.Shift_JIS")
	//os.Setenv("LANG", "C.EUC-JP")
	defer func() {
		if okLang {
			os.Setenv("LC_ALL", origLang)
		} else {
			os.Unsetenv("LC_ALL")
		}
	}()

	ansiResult, err := mbcs.Utf8ToAnsi(utf8, mbcs.ACP)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(ansiExpect, ansiResult) {
		t.Fatalf("expect %#v but %#v", ansiExpect, ansiResult)
	}
}
