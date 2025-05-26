package prime

import (
	"errors"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, error) {

	if n == 0 {
		return 0, errors.New("zero")
	}
	primeCounter := 0
	for i := 1; i > 0; i++ {

		if isPrime(i) {
			primeCounter++
		}
		if primeCounter == n {
			return i, nil
		}
	}
	return n, nil
}
