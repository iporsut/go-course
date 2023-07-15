package main

import "fmt"

func main() {
	v1 := 10
	v2 := v1

	fmt.Println("v1", v1)
	fmt.Println("v2", v2)

	v1 = 20
	fmt.Println("Change value of v1 to 20")
	fmt.Println("v1", v1)
	fmt.Println("v2", v2)

	var p1 *int
	p1 = &v1
	fmt.Println("p1 is a pointer to v1")
	fmt.Println("p1", p1)

	*p1 = 30
	fmt.Println("Change value of v1 to 30 by pointer")
	fmt.Println("v1", v1)
	fmt.Println("*p1", *p1)

	fmt.Println("Change value of v1 to 50 by function")
	changeTo50(v1) // v = v1
	fmt.Println("v1", v1)

	fmt.Println("Change value of v1 to 50 by pointer to function")
	changeTo50ByPointer(&v1)
	fmt.Println("v1", v1)

}

func changeTo50(v int) {
	v = 50
}

func changeTo50ByPointer(p *int) {
	*p = 50
}
