package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error converting to int", err)
		os.Exit(-1)
	}
	fmt.Println(n)
}
