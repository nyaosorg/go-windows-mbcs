package mbcs

const THREAD_ACP = 3
const ACP = 0

// AtoU Converts Ansi-bytes to UTF8-String
func AtoU(mbcs []byte, codepage uintptr) (string, error) {
	return atou(mbcs, codepage)
}

// UtoA Converts UTF8-String to Ansi-bytes
func UtoA(utf8 string, codepage uintptr) ([]byte, error) {
	return utoa(utf8, codepage)
}
