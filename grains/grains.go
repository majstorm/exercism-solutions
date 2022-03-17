package grains

import (
	"errors"
)

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Square(number int) (uint64, error) {

	if number <= 0 || number > 64 {
		return 0, errors.New("incorrect field")
	}
	return uint64(IntPow(2, number-1)), nil
}

func Total() uint64 {
	return uint64(IntPow(2, 65) - 1)
}
