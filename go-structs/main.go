package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// alex := person{"Alex", "Anderson"}
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Bob",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 90210,
		},
	}

	// Turn value into address
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// Go shortcut even though type mismatch person to *person
	jim.updateName("jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// type description *, not an operator
func (pointerToPerson *person) updateName(newFirstName string) {
	// Turn address into value
	(*pointerToPerson).firstName = newFirstName
}
