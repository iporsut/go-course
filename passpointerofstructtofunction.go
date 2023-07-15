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

	DisplayCustomer(cust)
	fmt.Println()
	DisplayCustomerWithPointer(&cust)

	fmt.Println("Try to change name")
	changeName(cust)

	fmt.Println("Try to change name by pointer")
	changeNameByPointer(&cust)
}

func DisplayCustomer(cust Customer) {
	fmt.Println("CUSTOMER RECORD:")
	fmt.Println("FIRSTNAME:", cust.FirstName)
	fmt.Println("LASTNAME:", cust.LastName)
	fmt.Println("AGE:", cust.Age)
	fmt.Println("ADDRESS")
	fmt.Println("STREET:", cust.Address.Street)
	fmt.Println("CITY:", cust.Address.City)
	fmt.Println("STATE:", cust.Address.State)
	fmt.Println("POSTAL CODE:", cust.Address.PostalCode)
	fmt.Println("COMPANY")
	fmt.Println("NAME:", cust.Company.Name)
	fmt.Println("ADDRESS")
	fmt.Println("STREET:", cust.Company.Address.Street)
	fmt.Println("CITY:", cust.Company.Address.City)
	fmt.Println("STATE:", cust.Company.Address.State)
	fmt.Println("POSTAL CODE:", cust.Company.Address.PostalCode)
}

func DisplayCustomerWithPointer(cust *Customer) {
	fmt.Println("CUSTOMER RECORD:")
	fmt.Println("FIRSTNAME:", cust.FirstName)
	fmt.Println("LASTNAME:", cust.LastName)
	fmt.Println("AGE:", cust.Age)
	fmt.Println("ADDRESS")
	fmt.Println("STREET:", cust.Address.Street)
	fmt.Println("CITY:", cust.Address.City)
	fmt.Println("STATE:", cust.Address.State)
	fmt.Println("POSTAL CODE:", cust.Address.PostalCode)
	fmt.Println("COMPANY")
	fmt.Println("NAME:", cust.Company.Name)
	fmt.Println("ADDRESS")
	fmt.Println("STREET:", cust.Company.Address.Street)
	fmt.Println("CITY:", cust.Company.Address.City)
	fmt.Println("STATE:", cust.Company.Address.State)
	fmt.Println("POSTAL CODE:", cust.Company.Address.PostalCode)
}

func changeName(cust Customer) {
	cust.FirstName = "Bob"
}

func changeNameByPointer(cust *Customer) {
	cust.FirstName = "Bob"
}
