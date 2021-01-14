package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}
	fmt.Println(alex)

	// third way of declaring a struct
	var alex2 person
	fmt.Println(alex2)
	fmt.Printf("%+v", alex2)

	alex2.firstName = "Alex-2"
	alex2.lastName = "Anderson-2"
	fmt.Printf("%+v", alex2)

	// contactinfo

	jim := person{
		firstName: "Jim",
		lastName:  "Booo",
		contact: contactInfo{
			email:   "jim.booo@gmail.com",
			zipCode: "6070",
		},
	}
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateFirstName(firstName string) {
	p.firstName = firstName
}