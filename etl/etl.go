package etl

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) (ret map[string]int) {

	ret = map[string]int{}
	for pos, line := range in {
		for _, v := range line {
			fmt.Println(pos, line, v)
			ret[strings.ToLower(v)] = pos
		}
	}
	return
}
