package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}
func main() {
	// option 1: default values
	e := Employee{}
	fmt.Printf("%v\n", e)

	// option 2: initial values
	e2 := Employee{id: 2, name: "William", vacation: true}
	fmt.Printf("%v\n", e2)

	// option 3: using new keyword
	e3 := new(Employee)
	fmt.Printf("%v\n", *e3)
	e3.id = 3
	e3.name = "Leonardo"
	e3.vacation = true
	fmt.Printf("%v\n", *e3)

	// option 4: using a custom function
	e4 := NewEmployee(5, "Xochilt", false)
	fmt.Printf("%v\n", *e4)
}
