package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	rate := 0.0
	switch {
	case balance < 0:
		rate = 3.213
	case balance >= 0 && balance < 1000:
		rate = 0.5
	case balance >= 1000 && balance < 5000:
		rate = 1.621
	case balance >= 5000:
		rate = 2.475
	}

	return float32(rate)

}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	rate := InterestRate(balance)
	return balance / 100 * float64(rate)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	i := 0
	for balance < targetBalance {
		balance += Interest(balance)
		i++
	}

	return i
}
