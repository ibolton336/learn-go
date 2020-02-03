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
	// ianPointer := &ian

	ian.updateName("yanni")
	ian.print()

}
func (p person) print() {
	fmt.Printf("%v", p)
}

// type of pointer indicated by *person

// go will auto coax the value from the address here by using *person shortcut
func (pointerToPerson *person) updateName(newFirstName string) {

	//turn address to value
	(*pointerToPerson).firstName = newFirstName
}

//slices arent pass by value?
