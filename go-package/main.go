package main

import (
	"fmt"
	su "go-package/strutil" // <- package path
)

func main() {
	var telephone = "123-456-7890"
	tel := su.RemoveDash(telephone) // packagename.FunctionName
	fmt.Println(tel)
}
