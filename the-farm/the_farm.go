package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	amount int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.amount)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0.0, &SillyNephewError{amount: cows}
	}

	amount, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		amount = amount * 2
	} else if err != nil {
		return 0.0, err
	}

	if amount < 0 {
		return 0.0, errors.New("negative fodder")
	}

	return amount / float64(cows), nil
}
