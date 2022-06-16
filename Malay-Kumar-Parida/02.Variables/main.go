package main

import "fmt"

var l float32 = 100 // Lower case visible thorughout this package
var L int = 50      // Visible everywhere
// Cant use 3rd type of declarration here

// Variables can be grouped like this
var (
	name    string = "Malay"
	college string = "VSSUT"
	regd    int    = 100
)

func main() {
	var i int // 1st style of variable declaration without initialization
	i = 10
	fmt.Printf("%v,%T\n", i, i)

	var j int = 20 // 2nd Style of variable declaration with initialization
	fmt.Printf("%v,%T\n", j, j)

	k := 30 // 3rd style of Variable declaration without telling type
	fmt.Printf("%v,%T\n", k, k)

	// All declared variables have to be used

	fmt.Printf("%v,%T\n", l, l)

	fmt.Println("My name is ", name, ",I am from ", college, ". And my regd number is ", regd)

	var l int = 300 // Variables can be shadowed but not redeclared within same socpe
	fmt.Println(l)

	l = 500 // Variables reasigned
	fmt.Println(l)

	// Small variable name should be for immediate use
	// Good variable name for longer use times

	// Changing variables types
	num1 := 10.5
	var num2 int = int(num1)
	fmt.Printf("%v,%T,%v,%T\n", num1, num1, num2, num2)
	// Have to do explicitly

	// Primitive Data Types

	// Three types mainly
	// Boolean,Numeric,Text

	// Boolean Types
	// All variables are initialised to zero i.e false

	var b bool = true
	fmt.Printf("%v, %T\n", b, b)

	// Comparision
	n := 1 == 1
	m := 1 == 2

	fmt.Printf("%v, %T, %v, %T\n", m, m, n, n)

	// Integer Types
	// Initialised to numeric zero

	// int types
	var integer int = 1
	fmt.Printf("%v, %T\n", integer, integer)
	// Can be minimum of 32 bit, or 64bit or 128 bit depending on system.
	// We can have int8,int 16,int 32,int 64

	// We also have unsigned int
	// Can not add two different types implicitly
	var num3 int8 = 10
	var num4 uint = 20
	fmt.Println(num3 + int8(num4))

	// Floating point numbers
	// Default float 64, Cant operate with float32 & float64
	num5 := 3.14
	num6 := 2.5

	fmt.Println(num5 / num6)

	// Complex Numbers
	// Two types complex 64 or 128
	var num7 complex64 = 1 + 2i
	fmt.Printf("%v,%T\n", num7, num7)

	// Text Types
	// String
	s := "This is a string"
	fmt.Printf("%v,%T\n", s, s)

	// We can treat strings as arrays
	fmt.Printf("%v,%T\n", s[2], s[2])
	// We get the ascii values

	fmt.Printf("%v,%T\n", string(s[2]), s[2])

	// Strings are immuatable

	// String concatenation
	s1 := "This is also a string"

	fmt.Println(s, " ", s1)

	// Convert string to collection of bytes

	r := 'a'
	fmt.Printf("%v,%T\n", r, r)

	// Constants

	// Typed constants
	const myConst int = 42
	fmt.Printf("%v,%T\n", myConst, myConst)
	myVar := 10

	// Constants can not be modified
	// Constants have to be set at compile time

	// Constant declaration can be shadowed
	fmt.Printf("%v,%T\n", myConst+myVar, myConst+myVar)
}
