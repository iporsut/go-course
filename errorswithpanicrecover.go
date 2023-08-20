package main

import "fmt"

func divide(x, y int) int {
	if y == 0 {
		panic("can't divide by zero")
	}
	return x / y
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()
	fmt.Println("5/0", divide(5, 0))
	fmt.Println("5/2", divide(5, 2))
}
