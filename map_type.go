package main

import (
	"fmt"
)

func main() {
	var info map[string]string

	fmt.Println("Zero value is nil")
	fmt.Println(info)        // map[]
	fmt.Println(info == nil) // true

	fmt.Println("Initialize map literal")
	info2 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	fmt.Println(info2) // map[Country:Canada Name:John]

	fmt.Println("empty map literal is not nil")
	info3 := map[string]string{}
	fmt.Println(info3)        // map[]
	fmt.Println(info3 == nil) // false

	fmt.Println("Make map with make function")
	info4 := make(map[string]string)
	fmt.Println(info4) // map[]

	fmt.Println("Count entries of map")
	info5 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	fmt.Println(len(info5)) // 2

	fmt.Println("Append value to slice")
	info6 := map[string]string{}
	info6["Name"] = "John"
	info6["Country"] = "Canada"
	fmt.Println(info6) // map[Country:Canada Name:John]

	fmt.Println("Loop with for range loop")
	info7 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	for k, v := range info7 {
		fmt.Println(k, v)
	}

	info8 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	for _, v := range info8 {
		fmt.Println(v)
	}

	info9 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	for k := range info9 {
		fmt.Println(info9[k])
	}

	fmt.Println("Delete entry from map")
	info10 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	delete(info10, "Name")
	fmt.Println(info10) // map[Country:Canada]

	fmt.Println("Check if key exists in map")
	info11 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	_, ok := info11["Name"]
	fmt.Println(ok) // true

	_, ok = info11["Age"]
	fmt.Println(ok) // false

	fmt.Println("Get not exists key from map")
	info12 := map[string]string{
		"Name":    "John",
		"Country": "Canada",
	}
	v := info12["Age"]
	fmt.Println(v) // ""
}
