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
	Company  *Company
}

type Company struct {
	Name    string
	Address Address
}

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
		FirsName: "John",
		LastName: "Smith",
		Age:      30,
		Address: Address{
			Street:     "1234 Main Street",
			City:       "Columbus",
			State:      "Ohio",
			PostalCode: "43210",
		},
		Company: &comp,
	}

	fmt.Println(cust)
	fmt.Println(cust.FirsName)
	fmt.Println(cust.LastName)
	fmt.Println(cust.Age)
	fmt.Println(cust.Address)
	fmt.Println(cust.Address.Street)
	fmt.Println(cust.Address.City)
	fmt.Println(cust.Address.State)
	fmt.Println(cust.Address.PostalCode)
	fmt.Println(cust.Company)
	fmt.Println(cust.Company.Name)
	fmt.Println(cust.Company.Address)
	fmt.Println(cust.Company.Address.Street)
	fmt.Println(cust.Company.Address.City)
	fmt.Println(cust.Company.Address.State)
	fmt.Println(cust.Company.Address.PostalCode)

}
