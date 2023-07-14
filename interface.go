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

type CustomerNotifier interface {
	Notify(cust *Customer)
}

type CustomerRenderer interface {
	Render(cust *Customer)
}

type Service struct {
	Notifier CustomerNotifier
	Renderer CustomerRenderer
	Customer []*Customer
}

func (s *Service) AddCustomer(cust *Customer) {
	s.Customer = append(s.Customer, cust)
}

func (s *Service) NotifyCustomer() {
	for _, cust := range s.Customer {
		s.Notifier.Notify(cust)
	}
}

func (s *Service) RenderCustomer() {
	for _, cust := range s.Customer {
		s.Renderer.Render(cust)
	}
}

type EmailNotifier struct{}

func (e *EmailNotifier) Notify(cust *Customer) {
	fmt.Println("Sending email to", cust.FirsName, cust.LastName)
}

type SMSNotifier struct{}

func (s *SMSNotifier) Notify(cust *Customer) {
	fmt.Println("Sending SMS to", cust.FirsName, cust.LastName)
}

type CommandlineRenderer struct{}

func (c *CommandlineRenderer) Render(cust *Customer) {
	fmt.Println("CUSTOMER RECORD:")
	fmt.Println("FIRSTNAME:", cust.FirsName)
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

type WebRenderer struct{}

func (w *WebRenderer) Render(cust *Customer) {
	// render to web
	fmt.Println("<h1>CUSTOMER RECORD:</h1>")
	fmt.Println("<h2>FIRSTNAME:", cust.FirsName, "</h2>")
	fmt.Println("<h2>LASTNAME:", cust.LastName, "</h2>")
	fmt.Println("<h2>AGE:", cust.Age, "</h2>")
	fmt.Println("<h2>ADDRESS</h2>")
	fmt.Println("<h2>STREET:", cust.Address.Street, "</h2>")
	fmt.Println("<h2>CITY:", cust.Address.City, "</h2>")
	fmt.Println("<h2>STATE:", cust.Address.State, "</h2>")
	fmt.Println("<h2>POSTAL CODE:", cust.Address.PostalCode, "</h2>")
	fmt.Println("<h2>COMPANY</h2>")
	fmt.Println("<h2>NAME:", cust.Company.Name, "</h2>")
	fmt.Println("<h2>ADDRESS</h2>")
	fmt.Println("<h2>STREET:", cust.Company.Address.Street, "</h2>")
	fmt.Println("<h2>CITY:", cust.Company.Address.City, "</h2>")
	fmt.Println("<h2>STATE:", cust.Company.Address.State, "</h2>")
	fmt.Println("<h2>POSTAL CODE:", cust.Company.Address.PostalCode, "</h2>")
}

func main() {
	// make service
	svc := Service{
		Notifier: &SMSNotifier{},
		Renderer: &WebRenderer{},
	}

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

	svc.AddCustomer(&cust)

	svc.NotifyCustomer()
	fmt.Println()
	svc.RenderCustomer()
}
