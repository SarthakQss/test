package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	Name := person{
		firstName: "Pruthi",
		lastName:  "Qss",
		contactInfo: contactInfo{
			email:   "Sarthak@gmail.com",
			zipCode: 136118,
		},
	}

	Name.updateName("Sarthak")
	Name.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
