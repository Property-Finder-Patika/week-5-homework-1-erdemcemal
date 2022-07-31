package main

import (
	"sync"
)

const NumberOfGoroutines = 5 // Total number of goroutines to run concurrently
const NumberOfRequests = 50  // Total number of requests to make

func main() {
	var ch = make(chan int, NumberOfRequests)
	var wg sync.WaitGroup

	// starts NumberOfGoroutines goroutines that wait for something to do
	wg.Add(NumberOfGoroutines)
	for i := 0; i < NumberOfGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			a := NewAdobeProxy(&License{Key: "valid"})
			runAdobe(a, ch)
		}(i)
	}

	// wait for all goroutines to finish
	wg.Wait()
}
