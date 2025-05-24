package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n <= 0 {
		return 0, errors.New("error")
	}

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		steps++
	}

	return steps, nil
}
