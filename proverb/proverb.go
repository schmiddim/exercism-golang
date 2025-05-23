// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	sentence = "For want of a %s the %s was lost."
	last     = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	var ret []string
	if len(rhyme) == 0 {
		return ret
	}
	for i := 0; i < len(rhyme)-1; i++ {
		s := fmt.Sprintf(sentence, rhyme[i], rhyme[i+1])
		ret = append(ret, s)

	}

	st := fmt.Sprintf(last, rhyme[0])
	ret = append(ret, st)
	return ret
}
