package resistorcolorduo

import (
	"fmt"
	"strconv"
)

func Colors() []string {
	return []string{

		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}

}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	sumString := ""
	for n, color := range colors {
		for i, c := range Colors() {
			if color == c && n < 2 {
				sumString += fmt.Sprintf("%d", i)
			}
		}
	}
	sum, err := strconv.Atoi(sumString)
	if err != nil {
		fmt.Println(err)
	}
	return sum
}
