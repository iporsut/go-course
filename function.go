package main

import (
	"fmt"
	"io"
	"os"
)

func add(x int, y int) int {
	return x + y
}

func divide(x, y int) (int, int) {
	return x / y, x % y
}

func divide2(x, y int) (quotient int, remainder int) {
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

func exampleOfDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func exampleOfDefer2() (string, error) {
	// Read file and close it by defer
	f, err := os.Open("filename.ext")
	if err != nil {
		return "", err
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func functionAsValue() {
	f := func() {
		fmt.Println("Hello")
	}
	f()
}

func MapInts(f func(int) int, xs []int) []int {
	ys := make([]int, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func Map[T any, R any](f func(T) R, xs []T) []R {
	ys := make([]R, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func main() {

}
