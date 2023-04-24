package mbcs

import (
	"bytes"

	"golang.org/x/text/transform"
)

func newDecoder(cp uintptr) transform.Transformer {
	return _Decoder{CP: cp}
}

// _Decoder is a transform.Transformer implementation that converts ANSI strings to UTF8 strings.
type _Decoder struct {
	CP uintptr
}

// Reset does nothing in _Decoder
func (f _Decoder) Reset() {}

// Transform converts the ANSI string in src to a UTF8 string and stores it in dst.
func (f _Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
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
		to, err := ansiToUtf8(from, f.CP)
		if err != nil {
			return nDst, nSrc, err
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
