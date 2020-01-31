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
	ian.updateName("yanni")
	ian.print()

}
func (p person) print() {
	fmt.Printf("%v", p)
}
func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
