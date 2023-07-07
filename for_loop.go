package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 8 {
			continue
		}

		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	j := 10
	for j >= 0 {
		fmt.Println(j)
		j -= 2
	}

	m := 10
	for {
		fmt.Println(m)
		m -= 2
		if m < 0 {
			break
		}
	}
}
