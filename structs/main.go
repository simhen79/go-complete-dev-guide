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
		firstName: "Jim",
		lastName:  "Hardy",
		contactInfo: contactInfo{
			email:   "jims@example.com",
			zipCode: 12345,
		},
	}

	// jimPointer := &jim

	jim.updateName("Jimmy")
	jim.print()

	// var alex person

	// george := person{firstName: "George", lastName: "Lucas"}

	//fmt.Println(alex, george)

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
