package main

import "fmt"

type Student struct {
	age     int
	name    string
	friends []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal    // Embedding
	SpeedKMPH float64
	CanFly    bool
}

func main() {
	population := map[string]int{
		"Odisha":      100000,
		"West Bengal": 20000,
		"Karnataka":   30000,
	}
	fmt.Println(population)

	// Use make function
	grades := make(map[string]int) // without initializing
	grades = map[string]int{
		"Malay":  90,
		"Swagat": 80,
		"Rittik": 70,
		"Aitik":  60,
	}
	fmt.Println(grades["Malay"])
	grades["Siba"] = 50
	fmt.Println(grades["Siba"])
	delete(grades, "Malay")
	fmt.Println(grades)
	// We get 0 for non existent keys
	pop, ok := grades["Malay"]
	fmt.Println(pop, ok)
	fmt.Println(len(grades))
	// Maps are copied by reference

	// Struct
	aStudent := Student{
		age:     22,
		name:    "Malay Kumar Parida",
		friends: []string{"Siba", "Ghose"},
	}
	fmt.Println(aStudent)
	fmt.Println(aStudent.name)

	// Anonymous struct
	aDoctor := struct{ name string }{name: "Malay"}
	fmt.Println(aDoctor)

	// Structs are copied shallowly like array unlike Slice & maps

	// Embedding
	// This is the model of inheritance that Go follows

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKMPH = 48
	b.CanFly = false
	fmt.Println(b)
	// In go we are not inheriting, we are embedding
}
