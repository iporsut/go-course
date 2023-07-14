package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("everything is pass by value")
	fmt.Println("Try to reassign map parameter")
	reassignMap(m)
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Inplace double")
	inplaceMapDouble(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func reassignMap(m map[string]int) {
	m = map[string]int{
		"five": 5,
		"four": 4,
		"one":  1,
	}
}

func inplaceMapDouble(m map[string]int) {
	for k, v := range m {
		m[k] = v * 2
	}
}
