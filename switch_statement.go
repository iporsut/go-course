package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("MENU")
	fmt.Println("G: Greeting")
	fmt.Println("Q: Quit")
	var menu string
	if scanner.Scan() {
		menu = scanner.Text()
	}

	switch menu {
	case "Q":
		fmt.Println("Bye!")
	case "G":
		fmt.Println("Hello, World!")
	}
}
