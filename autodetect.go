package mbcs

import (
	"bytes"
	"unicode/utf8"

	"golang.org/x/text/transform"
)

// AutoDecoder is an implementation of transform.Transformer that converts strings that are unknown to ANSI or UTF8 to UTF8.
type AutoDecoder struct {
	CodePage uintptr
}

// Reset does nothing in AutoDecoder
func (f AutoDecoder) Reset() {}

// Transform converts a strings that are unknown to ANSI or UTF8 in src to a UTF8 string and stores it in dst.
func (f AutoDecoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
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
			to, err = ansiToUtf8(from, f.CodePage)
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
