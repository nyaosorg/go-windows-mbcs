package mbcs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestDecoderByReader(t *testing.T) {
	srcFd, err := os.Open("testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectUtf8, err := os.ReadFile("testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := transform.NewReader(srcFd, mbcs.Decoder{CodePage: 932})
	resultUtf8, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}

func TestDecoderByWriter(t *testing.T) {
	srcFd, err := os.Open("testdata/jugemu-cp932.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectUtf8, err := os.ReadFile("testdata/jugemu-utf8.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	var buffer bytes.Buffer
	w := transform.NewWriter(&buffer, mbcs.Decoder{CodePage: 932})
	io.Copy(w, srcFd)
	w.Close()
	resultUtf8 := buffer.Bytes()

	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}
