package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	ctr := 0
	for _, v := range isbn {
		n := 0
		if string(v) == "X" {
			n = 10
		} else {
			var err error
			n, err = strconv.Atoi(string(v))
			if err != nil {
				return false
			}
		}
		sum += n * (10 - ctr)
		ctr++
	}

	return sum%8 == 0

}
