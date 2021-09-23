package main

import (
	"fmt"
	"time"
)

//main goroutine
func main() {
	//new goroutine gets created; non-blocking
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) //main goroutine is blocked until the function returns
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	// at this point, main returns, all goroutines are killed and program exits.
}

// There is no programmatic way to stop one goroutine from another.
// spinner goroutine keeps executing until abrubtly terminated before program exit
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
