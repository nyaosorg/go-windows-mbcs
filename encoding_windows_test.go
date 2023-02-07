package mbcs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func TestAutoDecoderFromCP932ToUTF8(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-cp932.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.AutoDecoder{CP: 932})
}

func TestAutoDecoderFromUTF8ToUTF8(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.AutoDecoder{CP: 932})
}

func TestDecoderByReader(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-cp932.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.Decoder{CP: 932})
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
	w := transform.NewWriter(&buffer, mbcs.Decoder{CP: 932})
	io.Copy(w, srcFd)
	w.Close()
	resultUtf8 := buffer.Bytes()

	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}

func TestEncoder(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-cp932.txt",
		mbcs.Encoder{CP: 932})
}
