package mbcs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func testFiles(t *testing.T, aFilePath, uFilePath string, tr transform.Transformer) {
	srcFd, err := os.Open(aFilePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectUtf8, err := os.ReadFile(uFilePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := transform.NewReader(srcFd, tr)
	resultUtf8, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}

func TestAutoDecoderFromUTF8ToUTF8On65001(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.AutoDecoder{CP: 65001})
}

func TestDecoderFromUTF8ToUTF8On65001(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.Decoder{CP: 65001})
}
