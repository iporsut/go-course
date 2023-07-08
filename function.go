package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func divide(x, y int) (int, int) {
	return x / y, x % y
}

func divide2(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

func divide3(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by zero")
	}
	return x / y, nil
}

func main() {

}
