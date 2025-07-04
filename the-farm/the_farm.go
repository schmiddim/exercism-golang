package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	amount, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	f, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount / float64(cows) * f, err
}

func ValidateInputAndDivideFood(calculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(calculator, cows)
	}

	return 0.0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	count         int
	customMessage string
}

func (e InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.count, e.customMessage)
}

func ValidateNumberOfCows(cows int) error {

	if cows < 0 {
		return &InvalidCowsError{
			count:         cows,
			customMessage: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			count:         cows,
			customMessage: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
