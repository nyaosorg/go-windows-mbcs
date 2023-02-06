package mbcstrans

import (
	"bytes"
	"unicode/utf8"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs/internal/core"
)

type AutoDetectTransformer struct {
	CodePage uintptr
}

func (f AutoDetectTransformer) Reset() {}

func (f AutoDetectTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for len(src) > 0 {
		// println("called Transform")
		n := bytes.IndexByte(src, '\n')
		var from []byte
		if n < 0 {
			n = len(src)
			from = src
			if !atEOF {
				return nDst, nSrc, transform.ErrShortSrc
			}
		} else {
			n++
			from = src[:n]
		}
		var to string
		if utf8.Valid(from) {
			to = string(from)
		} else {
			var err error
			to, err = core.AnsiToUtf8(from, f.CodePage)
			if err != nil {
				return nDst, nSrc, err
			}
		}
		if len(dst) < len(to) {
			return nDst, nSrc, transform.ErrShortDst
		}
		for i, iEnd := 0, len(to); i < iEnd; i++ {
			dst[i] = to[i]
		}
		nSrc += n
		nDst += len(to)
		src = src[n:]
		dst = dst[len(to):]
	}
	return nDst, nSrc, nil
}
