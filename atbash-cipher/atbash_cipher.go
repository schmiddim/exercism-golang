package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var builder strings.Builder
	re := regexp.MustCompile(`\s`)

	s = re.ReplaceAllString(s, "")
	for _, r := range strings.ToLower(s) {
		switch {
		case unicode.IsLetter(r):
			r = 'z' - (r - 'a')

			builder.WriteRune(r)
		case unicode.IsDigit(r):
			builder.WriteRune(r)
		case unicode.IsControl(r):
			continue

		}

	}
	build := builder.String()

	ret := ""
	for i, r := range build {
		ret += string(r)
		if (i+1)%5 == 0 && len(build)-1 != i {
			ret += " "
		}
	}
	return ret
}
