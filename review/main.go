package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// explicit and inpliit variables
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// error handling
	myInt, err := strconv.ParseInt("Nan", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myInt)
	}

	// maps
	m := make(map[string]int)
	m["key"] = 79
	fmt.Println(m["key"])

	// slices
	s := []int{1, 2, 3}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index, value)
	}

	// channel and go rutines call
	c := make(chan int)
	go doSomething(c)
	fmt.Println(<-c)

	// pointers
	g := 25
	fmt.Println(g)
	h := &g
	// print memory pointer
	fmt.Println(h)
	// print value
	fmt.Println(*h)
}

// go rutines
func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
