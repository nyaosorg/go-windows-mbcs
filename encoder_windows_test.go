package mbcs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestEncoder(t *testing.T) {
	srcFd, err := os.Open("testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectCp932, err := os.ReadFile("testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := transform.NewReader(srcFd, mbcs.Encoder{CP: 932})
	resultCp932, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expectCp932, resultCp932) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectCp932, resultCp932)
	}
}
