package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("MENU")
	fmt.Println("G: Greeting")
	fmt.Println("S: Calc Score")
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
	case "S":
		fmt.Print("Please input your score: ")
		var score int
		var err error
		if scanner.Scan() {
			text := scanner.Text()
			score, err = strconv.Atoi(text)
			if err != nil {
				fmt.Println("Invalid score")
				return
			}
		}
		switch {
		case score < 60:
			fmt.Println("F")
		case score < 70:
			fmt.Println("D")
		case score < 80:
			fmt.Println("C")
		case score < 90:
			fmt.Println("B")
		case score <= 100:
			fmt.Println("A")
		default:
			fmt.Println("Invalid score")
		}
	}
}
