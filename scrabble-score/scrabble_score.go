package scrabble

import "strings"

func value(char rune, letters string, value int) int {
	for _, v := range letters {
		if v == char {
			return value
		}
	}
	return 0
}

func Score(word string) int {

	sum := 0

	for _, v := range strings.ToUpper(word) {
		sum += value(v, "AEIOULNRST", 1)
		sum += value(v, "DG", 2)
		sum += value(v, "BCMP", 3)
		sum += value(v, "FHVWY", 4)
		sum += value(v, "K", 5)
		sum += value(v, "JX", 8)
		sum += value(v, "QZ", 10)
	}
	return sum
}
