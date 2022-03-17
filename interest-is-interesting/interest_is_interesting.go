package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return float32(3.213)
	case balance < 1000:
		return float32(0.5)
	case balance < 5000:
		return float32(1.621)
	case balance >= 5000:
		return float32(2.475)
	}
	return 0
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	switch {
	case balance < 0:
		return float64(balance * 0.03213)
	case balance < 1000:
		return float64(balance * 0.005)
	case balance < 5000:
		return float64(balance * 0.01621)
	case balance >= 5000:
		return float64(balance * 0.02475)
	}
	return 0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yr := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		yr++
	}
	return yr
}
