package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	if scanner.Scan() {
		input = scanner.Text()
	}

	fmt.Println(input)
}
