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
	jim := person{
		firstName: "Jimmy",
		lastName:  "Jamerson",
		contactInfo: contactInfo{
			email:   "a@mail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Bondan")
	jim.print()

}

func (pointertoPerson *person) updateName(newFirstName string) {
	(*pointertoPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
