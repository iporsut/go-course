package main

import "fmt"

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
	Company   *Company
}

type Company struct {
	Name    string
	Address Address
}

func (c *Customer) Name() string {
	return c.FirstName + " " + c.LastName
}

func (c *Customer) ChangeName(newName string) {
	c.FirstName = newName
}

type MyString string

func (s MyString) Length() int {
	return len(s)
}

// type MyT a.T

// func (m MyT) Method() {

// }

func main() {
	comp := Company{
		Name: "ABC Inc.",
		Address: Address{
			Street:     "1234 Main Street",
			City:       "Columbus",
			State:      "Ohio",
			PostalCode: "43210",
		},
	}

	cust := Customer{
		FirstName: "John",
		LastName:  "Smith",
		Age:       30,
		Address: Address{
			Street:     "1234 Main Street",
			City:       "Columbus",
			State:      "Ohio",
			PostalCode: "43210",
		},
		Company: &comp,
	}

	cust.ChangeName("Michael") // (&cust).ChangeName("Michael")
	fmt.Println(cust.Name())

	s := MyString("Hello")
	fmt.Println(s.Length())
}
