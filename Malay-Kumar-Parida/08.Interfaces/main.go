package main

import (
	"fmt"
	"math"
)

type shape1 interface {
	area() float64 // Any Struct that have all this method is of type shape
}
type shape2 interface {
	parameters() float64 // Any struct that have all the methods here is of this type
}

type shape interface {
	shape1
	shape2
}
type circle struct {
	radius float64
}

type rect struct { // Interfaces composed of other interfaces
	width  float64
	height float64
}

func (r rect) parameters() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) parameters() float64 {
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape1) float64 {
	return s.area()
}

func main() {
	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape1{c1, r1}
	for _, v := range shapes {
		fmt.Println(v.area()) // Can only use the method of the interface not the individual member functions
	}
	for _, v := range shapes {
		fmt.Println(getArea(v)) // We can pass the shapes around to function
	}

	shapes2 := []shape2{c1, r1}
	for _, v := range shapes2 {
		fmt.Println(v.parameters())
	}

	// Incrementing integer types
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// Interfaces composed of other interfaces
	polygons := []shape{c1, r1}
	for _, v := range polygons {
		fmt.Println(v.area(), v.parameters())
	}

}

// Not necessarily struct, any type that can have methods can implement an interface
type Incrementer interface {
	Increment() int
}
type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
