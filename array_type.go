package main

import "fmt"

func main() {
	var nums [5]int
	fmt.Println(nums) // [0 0 0 0 0]
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5
	// nums[5] = 0 // compile error
	fmt.Println(nums)

	nums2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums2)

	nums3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(nums3)

	nums4 := [5]int{1, 2, 3}
	fmt.Println(nums4)

	nums5 := [5]int{1: 10, 3: 30}
	fmt.Println(nums5)

	nums6 := [...]int{1: 10, 3: 30}
	fmt.Println(nums6)

	fmt.Println("Loop with for loop")
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println("Loop with for range loop")
	for i, v := range nums {
		fmt.Println(i, v)
	}

	fmt.Println("Loop with for range loop only value")
	for _, v := range nums {
		fmt.Println(v)
	}

	fmt.Println("Loop with for range loop only index")
	for i := range nums {
		fmt.Println(i)
	}
}
