package mbcs_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs"
)

func insertTop(source []byte, b byte) []byte {
	output := make([]byte, len(source)+1)
	output[0] = b
	copy(output[1:], source)
	return output
}

func testFiles(t *testing.T, aFilePath, uFilePath string, tr transform.Transformer) {
	source, err := os.ReadFile(aFilePath)
	if err != nil {
		t.Helper()
		t.Fatal(err.Error())
	}
	expect, err := os.ReadFile(uFilePath)
	if err != nil {
		t.Helper()
		t.Fatal(err.Error())
	}

	for i := 0; i < 3; i++ {
		// println("try:", i)
		r := transform.NewReader(bytes.NewReader(source), tr)
		result, err := io.ReadAll(r)
		if err != nil {
			t.Helper()
			t.Fatal(err.Error())
		}
		if !bytes.Equal(expect, result) {
			t.Helper()
			os.WriteFile("error-result.bin", result, 0666)
			t.Fatalf("not equal:\nExpect(size:%d)\nand\nResult(size:%d)\n", len(expect), len(result))
		}
		// Shift source and expect to test boundary at reading
		source = insertTop(source, 'x')
		expect = insertTop(expect, 'x')
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

// target
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
		os.WriteFile("TestDecodeByWriter.bin", resultUtf8, 0666)
		t.Fatalf("not equal: expect(%d bytes) and result(%d)\n", len(expectUtf8), len(resultUtf8))
	}
}

func TestEncoder(t *testing.T) {
	testFiles(t,
		"testdata/jugemu-utf8.txt",
		"testdata/jugemu-cp932.txt",
		mbcs.NewEncoder(932))
}

func TestDecoder431(t *testing.T) {
	testFiles(t,
		"testdata/nyagos_issue_431.txt",
		"testdata/nyagos_issue_431.txt",
		mbcs.NewDecoder(932))
}

func TestEecoder431(t *testing.T) {
	testFiles(t,
		"testdata/nyagos_issue_431.txt",
		"testdata/nyagos_issue_431.txt",
		mbcs.NewEncoder(932))
}
