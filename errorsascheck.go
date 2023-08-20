package main

import (
	"errors"
	"fmt"
)

type OperatorError struct {
	Type string
	Msg  string
}

func (e *OperatorError) Error() string {
	return e.Type + ": " + e.Msg
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, &OperatorError{
			Type: "ValueError",
			Msg:  "can't divide by zero",
		}
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

		var opErr *OperatorError
		if errors.As(err, &opErr) {
			fmt.Println("Type:", opErr.Type)
			fmt.Println("Msg:", opErr.Msg)
		}
	} else {
		fmt.Println(result)
	}
}
