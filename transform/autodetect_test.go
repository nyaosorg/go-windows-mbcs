package mbcstrans_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs/transform"
)

func testFiles(t *testing.T, aFilePath, uFilePath string) {
	srcFd, err := os.Open(aFilePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer srcFd.Close()

	expectUtf8, err := os.ReadFile(uFilePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	r := transform.NewReader(srcFd, mbcstrans.AutoDetectTransformer{CodePage: 932})
	resultUtf8, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err.Error())
	}
	if !bytes.Equal(expectUtf8, resultUtf8) {
		t.Fatalf("not equal:\n%#v\nand\n%#v", expectUtf8, resultUtf8)
	}
}

func TestAutoDetectFromCP932ToUTF8(t *testing.T) {
	testFiles(t, "../testdata/jugemu-cp932.txt", "../testdata/jugemu-utf8.txt")
}

func TestAutoDetectFromUTF8ToUTF8(t *testing.T) {
	testFiles(t, "../testdata/jugemu-utf8.txt", "../testdata/jugemu-utf8.txt")
}
