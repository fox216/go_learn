// factory example
// factory_1

package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

func PersonFactory(name string, age int, height float64) *Person {
	// Return person
	return &Person{
		Name:   name,
		Age:    age,
		Height: height,
	}
}

func main() {
	// get a person
	emp1 := PersonFactory("Fred", 20, 6.2)
	fmt.Printf("Hi im %s, im %d years old and %f tall\n", emp1.Name, emp1.Age, emp1.Height)
}
