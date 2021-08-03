package main

import "fmt"

type contactinfo struct{
	email string 
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact contactinfo
}



func main() {
jim := person{
	firstName: "Jimmy",
	lastName: "Gin",
	contact: contactinfo{
		email: "jim@gmail.com",
		zipcode: 94112,
	},
}

jim.updateName("jimanji")
jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string){
	(*p).firstName = newFirstName
}