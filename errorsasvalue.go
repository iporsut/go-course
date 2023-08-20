package main

import (
	"errors"
	"fmt"
)

/*
error interface that defined in standard library

	type error interface {
		Error() string
	}
*/

// Simple error with string message by fmt.Errorf function or errors.New function
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by zero")
		// return 0, errors.New("can't divide by zero")
	}
	return x / y, nil
}

func calc(x, y int) (float64, error) {
	v, err := divide(x, y)
	if err != nil {
		return 0, err
	}

	return float64(v), nil
}

func calcWrapErr(x, y int) (float64, error) {
	v, err := divide(x, y)
	if err != nil {
		return 0, fmt.Errorf("calcWrapErr: %w", err)
	}

	return float64(v), nil
}

func main() {
	// result, err := calc(5, 0)
	// if err != nil {
	// 	fmt.Println(err) // fmt.Println will call Error() method of error interface
	// } else {
	// 	fmt.Println(result)
	// }

	result, err := calcWrapErr(5, 0)
	if err != nil {
		fmt.Println(err) // fmt.Println will call Error() method of error interface
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(result)
	}
}
