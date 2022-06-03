package main

import "fmt"

func main() {
	// 1-D arrays

	// default values are set to 0
	var arr [3]int
	fmt.Println(arr) // [0 0 0]

	arr[0] = 31
	arr[1] = 24
	arr[2] = 5
	fmt.Println(arr) // [31 24 5]

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)      // [1 2 3]
	fmt.Println(arr2[0])   // 1
	fmt.Println(len(arr2)) // 3

	prod := 1
	for i := 0; i < len(arr); i++ {
		prod *= arr[i]
	}
	fmt.Println(prod) // 3720

	// 2-D arrays
	arr2D := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(arr2D)
	fmt.Print(arr2D[0][3], arr2D[2][3])
}
