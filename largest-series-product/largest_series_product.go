package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		return 0, errors.New("negative int ")
	}
	if span > len(digits) {
		return 0, errors.New("negative int ")

	}
	var maxN int64
	for i := 0; i <= len(digits)-span; i++ {
		series := digits[i : i+span]
		var sum int64 = 1
		for _, c := range series {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, err
			}
			sum *= int64(n)
		}
		if maxN < sum {
			maxN = sum
		}
	}
	return maxN, nil
}
