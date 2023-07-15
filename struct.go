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
	FirstName string
	LastName  string
}

type Employee struct {
	Person    // embedded struct type
	FirstName string
}

func main() {
	// c := Customer{
	// 	FirsName: "John",
	// 	LastName: "Smith",
	// 	Age:      30,
	// 	Address: Address{
	// 		Street:     "1234 Main Street",
	// 		City:       "Columbus",
	// 		State:      "Ohio",
	// 		PostalCode: "43210",
	// 	},
	// }
	// fmt.Println(c)
	// fmt.Println(c.FirsName)
	// fmt.Println(c.LastName)
	// fmt.Println(c.Age)
	// fmt.Println(c.Address)
	// fmt.Println(c.Address.Street)
	// fmt.Println(c.Address.City)
	// fmt.Println(c.Address.State)
	// fmt.Println(c.Address.PostalCode)

	e := Employee{
		Person: Person{
			FirstName: "John",
			LastName:  "Smith",
		},
		FirstName: "Jack",
	}
	fmt.Println(e)
	fmt.Println(e.FirstName) // e.Person.FirstName
	fmt.Println(e.Person.FirstName)
	fmt.Println(e.LastName) // e.Person.LastName
}
