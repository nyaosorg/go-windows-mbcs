package encoding

import (
	"bytes"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs/internal/core"
)

// Decoder is a transform.Transformer implementation that converts ANSI strings to UTF8 strings.
type Decoder struct {
	CodePage uintptr
}

// Reset does nothing in Decoder
func (f Decoder) Reset() {}

// Transform converts the ANSI string in src to a UTF8 string and stores it in dst.
func (f Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
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
		to, err := core.AnsiToUtf8(from, f.CodePage)
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
