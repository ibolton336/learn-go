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

	var yanni person
	yanni.firstName = "Yanni"
	yanni.lastName = "Laurel"

	ian := person{
		firstName: "Ian",
		lastName:  "Bolton",
		contactInfo: contactInfo{
			email:   "i@s.com",
			zipCode: 27609,
		}}

	//turn value into address
	ianPointer := &ian

	ianPointer.updateName("yanni")
	ian.print()

}
func (p person) print() {
	fmt.Printf("%v", p)
}

// type of pointer indicated by *person
func (pointerToPerson *person) updateName(newFirstName string) {

	//turn address to value
	(*pointerToPerson).firstName = newFirstName
}
