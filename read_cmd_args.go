package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	subCmd := os.Args[1]
	inputs := os.Args[2:]

	// Convert []string to []int
	var nums []int
	for _, v := range inputs {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Invalid input")
			os.Exit(-1)
		}
		nums = append(nums, num)
	}

	var result int
	var filterFunc func(int) bool
	switch subCmd {
	case "even":
		filterFunc = func(n int) bool {
			return n%2 == 0
		}
	case "odd":
		filterFunc = func(n int) bool {
			return n%2 == 1
		}
	case "all":
		filterFunc = func(n int) bool {
			return true
		}
	}

	for _, num := range nums {
		if filterFunc(num) {
			result += num
		}
	}

	fmt.Println(result)

	// ./calc odd 1 2 3 4 5
	// os.Args: [./calc odd 1 2 3 4 5]
}
