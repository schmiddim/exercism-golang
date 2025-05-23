package luhn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.TrimSpace(id)

	trimmed := strings.Replace(id, " ", "", -1)
	re := regexp.MustCompile("\\D")

	if re.MatchString(trimmed) {
		return false
	}

	if len(id) <= 1 {
		return false
	}

	fmt.Println(id)
	ctr := 1
	checkSum := 0
	for i := len(trimmed) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(trimmed[i]))
		fmt.Println(string(trimmed[i]))
		if ctr%2 == 0 {

			if err != nil {

				fmt.Println("ERROR")
			}
			n *= 2
			if n > 9 {
				n -= 9
			}

		}
		checkSum += n
		ctr++
	}
	return checkSum%10 == 0
}
