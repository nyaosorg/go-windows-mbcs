package mbcstrans_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs/transform"
)

func TestAnsiToUtf8Transformer(t *testing.T) {
	srcFd, err := os.Open("../testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectUtf8, err := os.ReadFile("../testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := transform.NewReader(srcFd, mbcstrans.AnsiToUtf8Transformer{CodePage: 932})
	resultUtf8, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}
