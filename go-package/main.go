package main

import (
	"fmt"
	"go-package/strutil"
)

func main() {
	var telephone = "123-456-7890"
	fmt.Println(strutil.RemoveDash(telephone))
}
