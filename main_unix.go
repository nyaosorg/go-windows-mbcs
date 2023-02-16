//go:build !windows
// +build !windows

package mbcs

import (
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

const cpUTF8 = 65001

var cpToEncoding = map[uintptr]encoding.Encoding{
	932:   japanese.ShiftJIS,
	936:   simplifiedchinese.GBK,
	949:   korean.EUCKR, // Unified Hangul Code
	950:   traditionalchinese.Big5,
	951:   traditionalchinese.Big5, // Big5-HKSCS
	50222: japanese.ISO2022JP,
	51932: japanese.EUCJP,
	51949: korean.EUCKR,
	52936: simplifiedchinese.HZGB2312,
}

var currentOsEncoding encoding.Encoding

func getCurrentOsEncoding() (e encoding.Encoding) {
	if currentOsEncoding != nil {
		return currentOsEncoding
	}
	defer func() {
		currentOsEncoding = e
	}()
	envLang, ok := os.LookupEnv("LC_ALL")
	if !ok {
		envLang, ok = os.LookupEnv("LANG")
		if !ok {
			return encoding.Nop
		}
	}
	periodPos := strings.IndexByte(envLang, '.')
	if periodPos >= 0 {
		envLang = envLang[periodPos+1:]
	}
	enc, err := ianaindex.IANA.Encoding(envLang)
	if err != nil {
		return encoding.Nop
	}
	return enc
}

// AtoU Converts Ansi-bytes to UTF8-String
func ansiToUtf8(mbcs []byte, codepage uintptr) (string, error) {
	var enc encoding.Encoding
	switch codepage {
	case ACP, THREAD_ACP:
		enc = getCurrentOsEncoding()
	case cpUTF8:
		return string(mbcs), nil
	default:
		var ok bool
		enc, ok = cpToEncoding[codepage]
		if !ok {
			return string(mbcs), ErrUnsupportedOs
		}
	}
	b, err := enc.NewDecoder().Bytes(mbcs)
	return string(b), err
}

// UtoA Converts UTF8-String to Ansi-bytes
func utf8ToAnsi(utf8 string, codepage uintptr) ([]byte, error) {
	var enc encoding.Encoding
	switch codepage {
	case ACP, THREAD_ACP:
		enc = getCurrentOsEncoding()
	case cpUTF8:
		return []byte(utf8), nil
	default:
		var ok bool
		enc, ok = cpToEncoding[codepage]
		if !ok {
			return []byte(utf8), ErrUnsupportedOs
		}
	}
	return enc.NewEncoder().Bytes([]byte(utf8))
}

func consoleCP() uintptr {
	return ACP
}
