package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	for ch := 'a'; ch <= 'z'; ch++ {
		if !strings.ContainsRune(strings.ToLower(input), ch) {
			return false
		}
	}
	return true
}
