/*
Example 1 basic interface
*/

package main // main package

import (
	"fmt"
	"math"
)

type geometry interface {
	// Area of the geometric shape
	area() float64
	// Perimeter of the geometric shape
	perim() float64
}

type rect struct {
	// Define attributes for rectangle
	width, height float64
}

type circle struct {
	// define attributes for circle
	radius float64
}

func (r rect) area() float64 {
	// Define function area that implements
	// the rect stucture
	return r.width * r.height
}

func (r rect) perim() float64 {
	// calculate perimeter of a rectangle
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	// calculate radius of circle
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	// calulate the perimeter of a circle
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	// print measurements
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
