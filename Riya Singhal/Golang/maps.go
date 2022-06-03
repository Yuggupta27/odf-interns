package main

import "fmt"

func main() {
	// creating and initializing empty map
	var map1 map[int]string // key -> int, value -> string

	// map declaration using shorthand property
	map2 := map[int]string{
		1:  "Rose",
		2:  "Sunflower",
		5:  "Lotus",
		10: "Lilly",
	}

	// map declaration using make function
	var map3 = make(map[int]string)

	// adding key-value pairs in map
	map2[31] = "Gold"
	map3[1] = "Riya"

	if map1 == nil {
		fmt.Println("Empty map")
	}

	// iterate over a map (sequence is not fixed)
	for key, value := range map2 {
		fmt.Printf("Key is: %d, value is: %s\n", key, value)
	}

	temp := map3[1]
	fmt.Println(temp)

	delete(map2, 10)
	fmt.Println(map2)

	// maps are of reference type. So changes made in copied map would reflect in original map too.
	map4 := map2
	map4[1] = "Lotus"
	fmt.Println(map2)
	fmt.Print(map4)
}
