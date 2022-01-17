package main

import "fmt"

func main() {
	grades := [3]int{70, 80, 90}
	grades2 := [...]int{50, 60, 70, 80, 90, 100} // Skip telling the size
	var students [5]string                       // Empty array
	students[0] = "Malay"
	students[1] = "John"
	students[2] = "Stuart"

	fmt.Printf("Grades:%v\n", grades)
	fmt.Printf("Grades2:%v\n", grades2)
	fmt.Printf("students:%v\n", students)
	fmt.Printf("Student#1-%v\n", students[0])
	fmt.Printf("Number of students: %v\n", len((students)))

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	// Arrays are shallow copied with only values
	// Copied with reference
	c := &a
	c[1] = 5
	fmt.Println(c)

	// Slice
	x := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(x))
	fmt.Printf("Capacity: %v\n", cap(x))

	// Slices are by default reference type
	y := x
	y[1] = 5
	fmt.Println(x)
	fmt.Println(y)

	m := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := m[:]   // Slice all the elements
	o := m[3:]  // Slice from 4th element
	p := m[:6]  // Slices first 6
	q := m[3:6] // Slices from 4th to 6th

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(q)

	r := make([]int, 0, 100) // Use make function(type,length,capacity)
	fmt.Println(r)
	fmt.Printf("Length: %v\n", len(r))
	fmt.Printf("Capacity: %v\n", cap(r))
	r = append(r, 1, 2, 3, 4, 5) // Any number of arguments
	// Use decompose ... to append list on list
	fmt.Println(r)
	fmt.Printf("Length: %v\n", len(r))
	fmt.Printf("Capacity: %v\n", cap(r))
	s := r[:len(r)-1] // Excluding the last element
	fmt.Println(s)

	// Removing the middle element
	t := append(r[:2], r[3:]...)
	fmt.Println(t)

}
