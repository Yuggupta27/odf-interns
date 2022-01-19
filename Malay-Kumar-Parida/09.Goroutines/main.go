package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0

var wg = sync.WaitGroup{} // Waitgroup helps in synchronization of goroutines
var m = sync.RWMutex{}

func main() {

	wg.Add(3) // Telling to wait for how many done signals

	go sayHello() // Spin off a green thread(light & cheap weight thread)
	// Main function is running in a goroutine
	// The go keyword creates another goroutine and puts the function there

	var msg = "Hello"
	go func() {
		fmt.Println(msg)
		wg.Done() // Telling the Goroutine is done
	}()
	msg = "Bye"
	// Due to closure property ths go routine has access to the main function variable msg
	// But this is creating race condition

	msg = "Hello"
	go func(msg string) { // This way we have decoupled the variables
		fmt.Println(msg)
		wg.Done()
	}(msg)

	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // Obtaining read lock
		go sayHi()
		m.Lock() // Obtaining write lock
		go increment()
	}
	// This example runs worse by using Goroutines because of us forcing it to behave in a single threaded manner

	// Without mutex this results in complete random order as our Groutines are not being synchronized with each other

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	//Gives number of processor threads available
	runtime.GOMAXPROCS(1)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1)) // Sets number of processor threads available
	wg.Wait()                                           // Waiting at this point for the waitgroup to be finished
	// During running use  -race in go run command to check concurrency
}

func sayHi() {
	fmt.Printf("Hello #%v\n", counter) // Protecting the counter variable
	m.RUnlock()
	wg.Done() // Releasing read lock
}

func increment() {
	counter++
	m.Unlock() //Releasing write lock
	wg.Done()
}

func sayHello() {
	fmt.Println("Hello")
	wg.Done()
}
