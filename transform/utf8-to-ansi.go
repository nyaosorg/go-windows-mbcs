package mbcstrans

import (
	"bytes"

	"golang.org/x/text/transform"

	"github.com/nyaosorg/go-windows-mbcs/internal/core"
)

type Utf8ToAnsiTransformer struct {
	CodePage uintptr
}

func (f Utf8ToAnsiTransformer) Reset() {}

func (f Utf8ToAnsiTransformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
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
		to, err := core.Utf8ToAnsi(string(from), f.CodePage)
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
