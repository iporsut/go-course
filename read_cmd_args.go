package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of string
	fmt.Println("os.Args:", os.Args)
	// ./calc odd 1 2 3 4 5
	// os.Args: [./calc odd 1 2 3 4 5]
}
