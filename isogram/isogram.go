package isogram

import (
	"regexp"
	"strings"
)

func match(c rune, matches string) bool {
	for _, m := range matches {
		if m == c {
			return true
		}

	}
	return false
}

func IsIsogram(word string) bool {
	matches := ""
	word = strings.ToLower(word)
	re := regexp.MustCompile("[^a-z]")
	word = re.ReplaceAllString(word, "")

	for _, c := range word {
		if match(c, matches) {
			return false
		}
		matches += string(c)

	}

	return true
}
