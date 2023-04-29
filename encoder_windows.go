package mbcs

import (
	"bytes"
	"unicode/utf8"

	"golang.org/x/text/transform"
)

func newEncoder(cp uintptr) transform.Transformer {
	return _Encoder{CP: cp}
}

// _Encoder is a transformer implementation that converts UTF8 strings to ANSI strings.
type _Encoder struct {
	CP uintptr
}

// Reset does nothing in _Encoder
func (f _Encoder) Reset() {}

// Transform converts the UTF8 string in src to an ANSI string and stores it in dst.
func (f _Encoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for len(src) > 0 {
		// println("called Transform")
		n := bytes.IndexByte(src, '\n')
		var from []byte
		if n < 0 {
			if atEOF {
				n = len(src)
				from = src
			} else {
				n = 0
				for n < len(src) {
					r,size := utf8.DecodeRune(src[n:])
					if r == utf8.RuneError {
						break
					}
					n += size
				}
				if n <= 0 {
					return nDst, nSrc, transform.ErrShortSrc
				}
				from = src[:n]
			}
		} else {
			n++
			from = src[:n]
		}
		to, err := utf8ToAnsi(string(from), f.CP)
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
