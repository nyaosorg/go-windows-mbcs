package mbcs

import (
	"golang.org/x/text/transform"
)

var procIsDBCSLeadByteEx = kernel32.NewProc("IsDBCSLeadByteEx")

func isDBCSLeadByteEx(cp uintptr, b byte) int {
	rc, _, _ := procIsDBCSLeadByteEx.Call(cp, uintptr(b))
	return int(rc)
}

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
func (f _Decoder) transform(dst, src []byte) (nDst, nSrc int, err error) {
	if f.CP == 65001 {
		n := copy(dst, src)
		return n, n, nil
	}
	n := 0
	for n < len(src) {
		if isDBCSLeadByteEx(f.CP, src[n]) != 0 {
			if n+2 > len(src) {
				err = transform.ErrShortSrc
				break
			}
			n += 2
		} else {
			n++
		}
	}
	utf8, _err := ansiToUtf8(src[:n], f.CP)
	if _err != nil {
		return 0, 0, _err
	}

	if len(dst) >= len(utf8) {
		return copy(dst, utf8), n, err
	}
	return 0, 0, transform.ErrShortDst
}

func (f _Decoder) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	const step = 80
	for {
		var _ndst, _nsrc int
		if len(src) > step {
			_ndst, _nsrc, err = f.transform(dst, src[:step])
		} else {
			_ndst, _nsrc, err = f.transform(dst, src[:])
		}
		if err != nil && err != transform.ErrShortSrc {
			return nDst, nSrc, err
		}
		nDst += _ndst
		dst = dst[_ndst:]
		nSrc += _nsrc
		src = src[_nsrc:]

		if len(src) <= 0 {
			return nDst, nSrc, nil
		}
		if len(src) < 2 {
			return nDst, nSrc, transform.ErrShortSrc
		}
	}
}
