package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("can't divide by zero")

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivideByZero
	}
	return x / y, nil
}

func calcWrapErr(x, y int) (float64, error) {
	v, err := divide(x, y)
	if err != nil {
		return 0, fmt.Errorf("calcWrapErr: %w", err)
	}

	return float64(v), nil
}

func main() {
	result, err := calcWrapErr(5, 0)
	if err != nil {
		fmt.Println(err) // fmt.Println will call Error() method of error interface
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("Error is ErrDivideByZero")
		}
	} else {
		fmt.Println(result)
	}
}
