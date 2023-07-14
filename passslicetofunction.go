package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("everything is pass by value")
	fmt.Println("Try to reassign slice parameter")
	reassignSlice(nums)
	for _, n := range nums {
		fmt.Println(n)
	}

	fmt.Println("Inplace double")
	inplaceDouble(nums)
	for _, n := range nums {
		fmt.Println(n)
	}
}

func reassignSlice(nums []int) {
	nums = []int{5, 4, 3, 2, 1}
}

func inplaceDouble(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}
