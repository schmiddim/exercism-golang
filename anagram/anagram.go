package anagram

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
func Detect(subject string, candidates []string) []string {
	var ret []string
	subject = strings.ToLower(subject)
	for i, candidate := range candidates {
		candidateLower := strings.ToLower(candidate)
		if candidateLower == subject {
			continue
		}
		if SortString(candidateLower) == SortString(strings.ToLower(subject)) {
			ret = append(ret, candidates[i])
		}
	}
	return ret
}
