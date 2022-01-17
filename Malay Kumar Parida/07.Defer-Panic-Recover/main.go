package main

import (
	"fmt"
	"log"
)

func main() {
	//defer
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	// Defered function are called after the function just before return
	// Executed in lifo order

	a := "start"
	defer fmt.Println(a) // Variables take the value of variable at the time of defering
	a = "end"

	// Panic
	// m,n:=1,0
	// ans:=m/n
	// fmt.Println(ans)

	fmt.Println("start")
	// panic("Something wrong happened")
	// panics happen after defer
	panicker()
}
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error", err)
		}
	}()
}
