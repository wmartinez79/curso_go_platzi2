package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "Will"
	fmt.Printf("%v\n", e)
	e.setId(3)
	e.setName("Donald")
	fmt.Println(e.getId())
	fmt.Println(e.getName())
}
