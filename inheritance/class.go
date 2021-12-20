package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTypeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTypeEmployee{}
	ftEmployee.id = 22
	ftEmployee.name = "William"
	ftEmployee.age = 42
	fmt.Printf("%v\n", ftEmployee)
}
