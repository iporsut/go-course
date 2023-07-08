package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int

	fmt.Println("Zero value is nil")
	fmt.Println(nums)        // []
	fmt.Println(nums == nil) // true

	fmt.Println("Initialize slice literal")
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums2) // [1 2 3 4 5]

	fmt.Println("empty slice literal is not nil")
	nums3 := []int{}
	fmt.Println(nums3)        // []
	fmt.Println(nums3 == nil) // false

	fmt.Println("Make slice with make function")
	nums4 := make([]int, 0)
	fmt.Println(nums4) // []

	nums5 := make([]int, 5) // []int{0,0,0,0,0}
	fmt.Println(nums5)      // [0 0 0 0 0]

	nums6 := make([]int, 0, 5)
	fmt.Println(nums6) // []

	fmt.Println("Count length of slice")
	nums7 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(nums7)) // 5

	fmt.Println("Append value to slice")
	nums8 := []int{}
	nums8 = append(nums8, 1)
	nums8 = append(nums8, 2)
	nums8 = append(nums8, 3)
	nums8 = append(nums8, 4)
	nums8 = append(nums8, 5)
	fmt.Println(nums8) // [1 2 3 4 5]

	nums8 = append(nums8, 6, 7, 8, 9, 10)
	fmt.Println(nums8) // [1 2 3 4 5 6 7 8 9 10]

	fmt.Println("Append slice to slice")
	nums9 := []int{}
	nums9 = append(nums9, 1)
	nums9 = append(nums9, 2)
	nums9 = append(nums9, 3)
	nums9 = append(nums9, 4)
	nums9 = append(nums9, 5)
	fmt.Println(nums9) // [1 2 3 4 5]

	othernums := []int{6, 7, 8, 9, 10}
	nums9 = append(nums9, othernums...)
	fmt.Println(nums9) // [1 2 3 4 5 6 7 8 9 10]

	fmt.Println("Loop with for range loop")
	nums10 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums10); i++ {
		fmt.Println(nums10[i])
	}
	// Output
	// 1
	// 2
	// 3
	// 4
	// 5

	nums11 := []int{1, 2, 3, 4, 5}
	for i := range nums11 {
		fmt.Println(nums11[i])
	}
	// Output
	// 1
	// 2
	// 3
	// 4
	// 5

	nums12 := []int{1, 2, 3, 4, 5}
	for _, v := range nums12 {
		fmt.Println(v)
	}
	// Output
	// 1
	// 2
	// 3
	// 4
	// 5

	nums13 := []int{1, 2, 3, 4, 5}
	for i, v := range nums13 {
		fmt.Printf("index %d: %d\n", i, v)
	}
	// Output
	// index 0: 1
	// index 1: 2
	// index 2: 3
	// index 3: 4
	// index 4: 5

	fmt.Println("Slice of slice")
	nums14 := []int{1, 2, 3, 4, 5}
	nums14 = nums14[1:3] // [x:y] => x จนถึง (y-1)
	fmt.Println(nums14)  // [2 3]

	nums15 := []int{1, 2, 3, 4, 5}
	nums15 = nums15[3:]
	fmt.Println(nums15) // [4 5]

	nums16 := []int{1, 2, 3, 4, 5}
	nums16 = nums16[:2]
	fmt.Println(nums16) // [1 2]

	nums16 = nums16[:]  // nums = [5]int{1,2,3,4,5} , slice = nums[:]
	fmt.Println(nums16) // [1 2 3 4 5]

	fmt.Println("Slice of string")
	phoneNumber := "0861234567"
	phoneNumber = phoneNumber[:3]
	fmt.Println(phoneNumber) // 086

	phoneNumber2 := "0861234567"
	phoneNumber2 = phoneNumber2[:3] + "xxx" + phoneNumber2[6:]
	fmt.Println(phoneNumber2) // 086xxx4567

	fmt.Println("convert between slice of byte and string")
	// Assume this data is read from `func ReadFile(name string) ([]byte, error)`
	fileData := []byte{72, 101, 108, 108, 111, 32, 87, 111, 114, 108, 100}
	msg := string(fileData)
	fmt.Println(msg) // Hello World

	msg2 := "Hello World"
	fileData2 := []byte(msg2) // `func Write([]byte) (int, error)`
	fmt.Println(fileData2)    // [72 101 108 108 111 32 87 111 114 108 100]

	fmt.Println("Copy slice")
	srcSlice := []int{1, 2, 3, 4, 5}
	dstSlice := make([]int, 3) // [0 0 0]
	copy(dstSlice, srcSlice)
	fmt.Println(dstSlice) // [1 2 3]

	srcSlice2 := []int{1, 2, 3, 4, 5}
	var dstSlice2 []int
	dstSlice2 = append(dstSlice2, srcSlice2[:3]...)
	fmt.Println(dstSlice2) // [1 2 3]

	fmt.Println("Sorting slice")
	nums17 := []int{5, 1, 2, 4, 3}
	sort.Ints(nums17)
	fmt.Println(nums17) // [1 2 3 4 5]

	nums18 := []int{5, 1, 2, 4, 3}
	sort.Slice(nums18, func(i, j int) bool {
		return nums18[i] > nums18[j]
	})
	fmt.Println(nums18) // [5, 4, 3, 2, 1]

	fmt.Println("Find element in slice")
	nums19 := []int{1, 2, 3, 4, 5}
	found := false
	element := 3
	for _, v := range nums19 {
		if v == element {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Found 3 in nums")
	}

	fmt.Println("Filter slice")
	nums20 := []int{1, 2, 3, 4, 5, 6}
	evenNums := []int{} // or just `var evenNums []int` or `evenNums := make([]int, 0)

	for _, v := range nums20 {
		if v%2 == 0 {
			evenNums = append(evenNums, v)
		}
	}

	fmt.Println(evenNums) // [2 4 6]
}
