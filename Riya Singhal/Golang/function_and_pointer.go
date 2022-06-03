package main

import "fmt"

func main() {
	var x int = 31
	var p *int = &x
	fmt.Println("Value of x is: ", x)
	fmt.Println("Address of x is: ", &x)
	fmt.Println("Value at pointer is: ", p)
	fmt.Println("Address of pointer is: ", &p)

	// the default value is <nil>
	var temp *int
	fmt.Println("Value of uninitialized pointer is: ", temp)

	// if we dont want to specify the type of pointers
	var temp2 = &x
	fmt.Println("Value at pointer is: ", temp2)

	// shorthand declarations
	z := 12
	y := &z
	fmt.Println("Value of z is: ", z)
	fmt.Println("Value of y is: ", y)

	// dereferencing
	fmt.Println("Value stored at p is", *p)

	*p = 24
	fmt.Println("x is also modified to: ", x)
	fmt.Println("Changed value pointed by p is: ", *p)

	// functions
	// pass by value
	val1 := 10
	val2 := 20
	incrementValues(val1, val2)
	fmt.Printf("Value1: %d, Value2: %d\n", val1, val2)

	// pass by reference
	value1 := 10
	value2 := 20
	decrementValues(&value1, &value2)
	fmt.Printf("Value1: %d, Value2: %d\n", value1, value2)

	// with return type
	result := add(val1, val2)
	fmt.Print("Result is: ", result)
}

func incrementValues(val1 int, val2 int) {
	val1 += 10
	val2 += 10
}

func decrementValues(value1 *int, value2 *int) {
	*value1 -= 10
	*value2 -= 10
}

func add(val1 int, val2 int) int {
	return val1 + val2
}
