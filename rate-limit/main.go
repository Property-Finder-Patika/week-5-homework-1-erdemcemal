package main

import "fmt"

// Create a sample program that simulates a race condition. So it should run fine when run sequentially but cause trouble when run in two or more goroutines.

func main() {
	//runSequentially() // this will run fine

	runWithRaceCondition() // output will not be in order

	fmt.Scanln()
}

// runSequentially is a function that runs the program sequentially.
func runSequentially() {
	for i := 1; i < 6; i++ {
		fmt.Printf("%d", i)
	}
}

// runWithRaceCondition is a function that runs the program
func runWithRaceCondition() {
	for i := 1; i < 6; i++ {

		go func(i int) {
			fmt.Printf("%d; ", i)
		}(i)

	}
}
