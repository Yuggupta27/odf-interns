package main

import "fmt"

func main() {
	//if statement
	if true {
		fmt.Println("This is a if statement")
	}
	//Alwasy use curly braces even if single statement block
	//If initializer run a statement for running the if block
	target := 10
	num := 11

	//Demonstarting all compariosion operators & Logical operators

	if num < 0 || num > 100 {
		fmt.Println("Enter a valid number in 0-100 range")
	}
	if num >= 0 && num <= 100 {
		if num == target {
			fmt.Println("Number is 10")
		}
		if num != target {
			if num < target {
				fmt.Println("Number is less than 10")
			}
			if num > target {
				fmt.Println("Number is greater than 10")
			}
		}

	}

	//Doing the same thing using else

	if num < 0 || num > 100 {
		fmt.Println("Enter a valid number in 0-100 range")
	} else {
		if num == target {
			fmt.Println("Number is 10")
		} else {
			if num < target {
				fmt.Println("Number is less than 10")
			} else {
				fmt.Println("Number is greater than 10")
			}
		}

	}
	//comparing float values are tricky

	//switch case
	//Using Initializer
	switch s := 5; s {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	//using multiple cases
	case 4, 5:
		fmt.Println("Four or Five")
	default:
		fmt.Println("Not expected number")
	}

	//switch 2nd type syntax
	//break statement is implicit fall throug can be given
	//cases can overlap,
	s2 := 10
	switch {
	case s2 <= 10:
		fmt.Println("Less than or equal to 10")
	case s2 <= 20:
		fmt.Println("Less than or equal to 20")
	default:
		fmt.Println("Greater than 20")
	}

	//Lopps

	//For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//multiple varibales, Can't be separted with commas

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
		if i+j == 6 {
			fmt.Println("Hurray")
		}
	}

	//There is no do & do while
	//These things can be done with skipping initialization or skipping increment

	//Similar to while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//Similar to Do While
	i = 0
Loop:
	for {
		i++
		if i == 7 {
			continue
		}
		fmt.Println(i)
		if i == 9 {
			break Loop //Can break to a specified label
		}
	}

	//Range based Loop

	s4 := []int{1, 2, 3}
	for k, v := range s4 {
		fmt.Println(k, v)
	}

	s5 := "Malay"
	for k, v := range s5 {
		fmt.Println(k, string(v))
	}

	//use _ to skip key part
}
