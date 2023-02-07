package mbcs_test

import (
	"os"
	"runtime"
	"testing"

	"github.com/nyaosorg/go-windows-mbcs"
)

var expect1 = []string{
	`ShiftJIS`,
	`で書いた。サンプルテキストです。`,
	`判定できるかな`,
}

var expect2 = []string{
	`UTF8`,
	`で書いた。サンプルテキストです。`,
	`判定できるかな`,
}

var expect3 = []string{
	`UTF8-BOM`,
	`で書いた。サンプルテキストです。`,
	`判定できるかな`,
}

func compareFileAndArray(t *testing.T, fname string, expect []string) bool {
	fd, err := os.Open(fname)
	if err != nil {
		t.Fatal(err.Error())
		return false
	}
	defer fd.Close()

	sc := mbcs.NewFilter(fd, 932)
	i := 0
	for sc.Scan() {
		if sc.Text() != expect[i] {
			t.Fatalf("not match with [%d] %s", i, expect[i])
			return false
		}
		i++
	}
	return true
}

func TestFilter(t *testing.T) {
	if !compareFileAndArray(t, "testdata/testdata-utf8.txt", expect2) {
		return
	}
	if !compareFileAndArray(t, "testdata/testdata-bom.txt", expect3) {
		return
	}
	if runtime.GOOS != "windows" {
		return
	}
	if !compareFileAndArray(t, "testdata/testdata-cp932.txt", expect1) {
		return
	}
}
