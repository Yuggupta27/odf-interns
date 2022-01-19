package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // Using the wait group to synchronize our routines with main function

func main() {
	// Channels can be used to pass data between different goroutines safely
	ch := make(chan int, 50) // Using make function to create a channel with the type of data to be moved
	// 2nd parameters shows the buffer size

	wg.Add(2)
	go func(ch <-chan int) { // We are setting this as a receive only channel
		// i := <-ch // Receiving data from channel
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)

		for i := range ch { // The loop stops when the channel is closed
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // We are setting this as a send only channel
		i := 100
		ch <- i  // Sending message to a channel
		ch <- 50 // Stored in the buffer
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
