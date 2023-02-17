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
		mbcs.NewAutoDecoder(65001))
}

func TestDecoderFromUTF8ToUTF8On65001(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.NewDecoder(65001))
}

func TestDecoderFromCP932ToUTF8(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-cp932.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.NewDecoder(932))
}

func TestAutoDecoderFromUTF8ToUTF8(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.NewAutoDecoder(932))
}

func TestDecoderByReader(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-cp932.txt",
		"testdata/jugemu-utf8.txt",
		mbcs.NewDecoder(932))
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
	w := transform.NewWriter(&buffer, mbcs.NewDecoder(932))
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
		mbcs.NewEncoder(932))
}
