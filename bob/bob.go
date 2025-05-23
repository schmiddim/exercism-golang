// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func isAllUpper(s string) bool {
	re := regexp.MustCompile(`[a-zA-Z]`)
	if !re.MatchString(strings.TrimSpace(s)) {
		return false
	}

	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func isQuestion(s string) bool {
	re := regexp.MustCompile(`.*\?$`)
	if re.MatchString(strings.TrimSpace(s)) {
		return true
	}
	return false
}
func isSilence(s string) bool {
	re := regexp.MustCompile(`\s`)
	if re.ReplaceAllString(s, "") == "" {
		return true
	}
	return false
}

func Hey(remark string) string {
	switch {
	case isQuestion(remark) && isAllUpper(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."

	case isSilence(remark):
		return "Fine. Be that way!"

	case isAllUpper(remark):

		return "Whoa, chill out!"

	default:

		return "Whatever."
	}

}
