package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	// explicit type
	var number uint16 = 31
	fmt.Println("Number is:", number)

	// implicit type
	var number2 = 31
	fmt.Println("Number is:", number2)

	// determining datatype
	fmt.Printf("%T\n", number2)

	// assignment expression (fastest one)
	number3 := 24
	fmt.Printf("%T\n", number3)

	// variable direct access
	fmt.Printf("Random %T %v\n", 31, 100)

	fmt.Printf("Binary representation of 10 is: %b\n", 10)

	// boolean (%t)
	x := 31
	y := x > 5
	fmt.Printf("%t\n", y)

	// if-else
	if number2 == 31 {
		fmt.Print("True\n")
	}

	// for loop
	for x := 1; x < 5; x++ {
		fmt.Println(x)
	}

	// for loop as while loop
	num := 1
	for num < 10 {
		if num%2 == 0 {
			fmt.Println(num)
		}
		num++
	}

	// switch case
	switch num {
	case 10:
		fmt.Print("1st case")
		break
	case 11:
		fmt.Print("2nd case")
		break
	default:
		fmt.Print("Default case")
	}
}
