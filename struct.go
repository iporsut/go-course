package main

import "fmt"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Customer struct {
	FirsName string
	LastName string
	Age      int
	Address  Address
}

type Person struct {
	FirsName string
	LastName string
}

type Employee struct {
	Person
}

func main() {
	c := Customer{
		FirsName: "John",
		LastName: "Smith",
		Age:      30,
		Address: Address{
			Street:     "1234 Main Street",
			City:       "Columbus",
			State:      "Ohio",
			PostalCode: "43210",
		},
	}
	fmt.Println(c)
	fmt.Println(c.FirsName)
	fmt.Println(c.LastName)
	fmt.Println(c.Age)
	fmt.Println(c.Address)
	fmt.Println(c.Address.Street)
	fmt.Println(c.Address.City)
	fmt.Println(c.Address.State)
	fmt.Println(c.Address.PostalCode)

	e := Employee{
		Person: Person{
			FirsName: "John",
			LastName: "Smith",
		},
	}
	fmt.Println(e)
	fmt.Println(e.FirsName)
	fmt.Println(e.LastName)
}
