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
	alex := person{
		firstName: "Alex",
		lastName:  "Dunphy",
		contactInfo: contactInfo{
			email:   "alexdunphy@gmail.com",
			zipCode: 10001,
		},
	}
	alex.updateName("Alexandra")
	alex.print()
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
