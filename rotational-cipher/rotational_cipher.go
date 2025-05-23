package rotationalcipher

func rot(c byte, n int) string {
	if c >= 'a' && c <= 'z' {
		return string('a' + (c-'a'+byte(n))%26)
	}
	if c >= 'A' && c <= 'Z' {
		return string('A' + (c-'A'+byte(n))%26)
	}
	return string(c)
}
func RotationalCipher(plain string, shiftKey int) string {
	ret := ""
	for _, c := range plain {
		ret += rot(byte(c), shiftKey)
	}
	return ret
}
