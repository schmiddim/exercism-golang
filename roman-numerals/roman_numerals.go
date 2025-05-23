package romannumerals

import (
	"errors"
	"strings"
)

var lookup = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("out of range")
	}

	var romanBuilder strings.Builder

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(val); i++ {
		for input >= val[i] {
			romanBuilder.WriteString(syms[i])
			input -= val[i]
		}
	}

	return romanBuilder.String(), nil
}
