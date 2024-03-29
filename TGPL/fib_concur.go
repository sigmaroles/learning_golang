package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wait...")
	go spinner(500 * time.Millisecond)
	// const n = 70
	// fibN := fib(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	time.Sleep(time.Second * 5)
}

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
