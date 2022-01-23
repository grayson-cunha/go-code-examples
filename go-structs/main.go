package main

import "fmt"

type contact struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact
}

func main() {
	jim := person{firstName:  "Jim", lastName: "Party", contact: contact{
		email: "jim@gmail.com",
		zipCode: 94000,
	}}

	jim.updateFirstName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName; 
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
