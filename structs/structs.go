package main

import "fmt"

type contact struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contact: contact{
			email:   "jim@email.com",
			zipcode: 12345,
		},
	}

	jim.updateName("James")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
