// Package diffsquares implements simple calculation of different square sums
package diffsquares

func Square(n int) int {
	return n * n
}

func Abs(n int) int {
	if n < 0 {
		return 0 - n
	}
	return n
}

// SquareOfSum implements the square of the sum of the positive numbers up to (including) n
func SquareOfSum(n int) int {
	return Square((n * (n + 1) / 2))
}

// SumOfSquares implements the sum of squares of the positive numbers up to (including) n
func SumOfSquares(n int) int {
	return (n * (n + 1) * ((n * 2) + 1)) / 6
}

func Difference(n int) int {
	return Abs(SumOfSquares(n) - SquareOfSum(n))
}
