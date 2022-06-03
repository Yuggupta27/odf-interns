package main

import "fmt"

type circle struct {
	radius int
	colour string
}

type rectangle struct {
	length  int
	breadth int
	colour  string

	// nested struct
	geometry struct {
		area      int
		perimeter int
	}
}

type student struct {
	rollNo int
	age    int
}

func main() {
	// using struct without instantiate a new instance of that type.
	fmt.Println(circle{10, "Red"})

	// creating instances
	var rect rectangle
	rect.length = 20
	rect.breadth = 30
	rect.colour = "Blue"
	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Printf("Area is: %d, Perimeter is: %d\n", rect.geometry.area, rect.geometry.perimeter)
	fmt.Println(rect)

	// creating a struct instance using a struct literal
	var circle1 = circle{10, "Blue"}
	fmt.Println(circle1)

	// mostly used
	circle2 := circle{20, "Pink"}
	fmt.Println(circle2)

	var circle3 = circle{radius: 15, colour: "Maroon"}
	fmt.Println(circle3)

	circle4 := circle{radius: 15, colour: "Maroon"}
	fmt.Println(circle4)

	// skipping radius so by default 0
	circle5 := circle{colour: "Green"}
	fmt.Println(circle5)

	// skipped color, by default empty string
	circle6 := circle{radius: 10}
	fmt.Println(circle6)

	// struct instantiation using new keyword
	cir := new(circle)
	cir.radius = 10
	cir.colour = "Yellow"
	fmt.Println(cir)

	// methods
	s1 := student{10, 20}
	res := s1.getAge()
	fmt.Println("Age is: ", res)
	s1.setAge(22)
	fmt.Print(s1.age)
}

// a function for object of type student, accepting no parameter and returning int.
func (S student) getAge() int {
	return S.age
}

// as pointers because I want the changes to be reflected
func (S *student) setAge(val int) {
	S.age = val
}
