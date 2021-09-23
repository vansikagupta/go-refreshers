/*
* Representing how channels can act as pipeline
* counter -> Squares -> Printer
 */

package main

import "fmt"

// Counter function
func counter(naturals chan int) {
	for x := 0; x < 20; x++ {
		// sending values throgh naturals channel
		naturals <- x
	}
	// Not neccessary to close every channel
	// it is required when we want to inform the receiver that no more values are coming you way
	close(naturals)
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)

	// function literal; anonymous function
	// Squarer function
	go func() {
		// receiver on naturals channel
		for x := range naturals {
			// sending values to the squares channel
			squares <- x * x
		}
		close(squares)
	}()

	fmt.Println("Execution of main goroutine blocked by squares channel")
	// Printer (in main goroutine)
	for val := range squares {
		fmt.Println(val)
	}
}
