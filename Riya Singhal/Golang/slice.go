package main

import (
	"fmt"
	"sort"
)

func main() {
	//slice is declared same as array but without size
	var slice = []int{1, 2, 3}
	fmt.Println(slice)

	//shorthand declaration for slice
	slice2 := []string{"Abc", "Xyz"}
	fmt.Println(slice2)

	//slice is the reference of array. example:
	arr := [5]int{1, 2, 3, 4, 5}

	//creating slices from the given array
	var slice1 = arr[1:4] //index 1 till index 3(incl.)
	var slice3 = arr[:]   //index 0 till last (all elements)
	var slice4 = arr[0:]  //index 0 till last (all elements)
	var slice5 = arr[:4]  //index 0 till index 3(incl.)

	fmt.Println(slice1)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)

	//creating slice from already existing slice
	var slice6 = slice1[1:2]
	fmt.Println(slice6)

	//make function
	var slice7 = make([]int, 3, 6) //type-length-capacity
	fmt.Printf("Slice: %v, Length of slice: %d, Capacity of slice: %d\n", slice7, len(slice7), cap(slice7))

	var slice8 = make([]int, 5)
	fmt.Printf("Slice: %v, Length of slice: %d, Capacity of slice: %d\n", slice8, len(slice8), cap(slice8))

	// iteration
	for ele := 0; ele < len(slice2); ele++ {
		fmt.Print(slice2[ele], " ")
	}
	fmt.Print("\n")

	// range in for loop
	for index, element := range slice {
		fmt.Printf("Index: %d, element: %d\n", index, element)
	}

	// Blank identifier: if we don't want index in range for loop, we can use "_" in place of index
	for _, element := range slice {
		fmt.Printf("element: %d\n", element)
	}

	// Creating empty slice or zero value slice (capacity = 0, length = 0)
	// Nil slice does not contain an array reference
	var nilSlice []int
	fmt.Printf("Slice: %v, length: %d, Capacity: %d", nilSlice, len(nilSlice), cap(nilSlice))

	// Slice is referencing array, so if we modify slice then array should also get modified.
	fmt.Printf("Original array: %v", arr)
	fmt.Println()
	fmt.Printf("Original slice: %v", slice1)
	fmt.Println()

	slice1[2] = 31
	arr[1] = 100

	fmt.Printf("New array: %v", arr)
	fmt.Println()
	fmt.Printf("New slice: %v", slice1)
	fmt.Println()

	// operator(==) in slice
	fmt.Println(nilSlice == nil)

	// sorting in slice
	sort.Ints(slice1) // import sort library for this
	fmt.Print(slice1)
}
