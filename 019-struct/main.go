package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type employee struct {
	person
	id        int
	firstName string
}

func main() {
	x := person{
		firstName: "rahul",
		lastName:  "kumar",
	}
	fmt.Println(x)
	fmt.Println(x.firstName)
	fmt.Println(x.lastName)

	x.firstName = "rahul2"
	x.lastName = "kumar2"
	fmt.Println(x)
	fmt.Println(x.firstName)
	fmt.Println(x.lastName)

	emp := employee{
		person: person{
			firstName: "rahul",
			lastName:  "kumar",
		},
		id:        12,
		firstName: "rk",
	}
	fmt.Println(emp)
	fmt.Println(emp.firstName, emp.lastName, emp.id)
	fmt.Println(emp.firstName, emp.person.firstName, emp.person.lastName, emp.id)

	anonymous := struct {
		name string
		age  int
	}{
		name: "rahul",
		age:  27,
	}
	fmt.Println(anonymous)
}
