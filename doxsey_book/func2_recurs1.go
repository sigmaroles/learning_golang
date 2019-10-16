package main

import "fmt"

/*
Vague flashbacks of 2008 (or was it 2009?), when I first learnt how to use recursion.
It's been a long time, and probably never too late...
A new language, a platform (cloud), a new set of goals!
*/

func product(inparr []int32, prod int32) int32 {
	if len(inparr) == 1 {
		return prod * inparr[0]
	}

	return product(inparr[1:], prod*inparr[0])
}

func sum2(inparr []int32, isum int32) int32 {

	if len(inparr) == 1 {
		return isum + inparr[0]
	}

	return sum2(inparr[1:], isum+inparr[0])
}

func fibonacci(n int) int64 {

	// nested function syntax: https://stackoverflow.com/a/43555474
	// memoization: https://forum.golangbridge.org/t/memoization-in-go/7285/10

	cache := make(map[int]int64)

	var recurseFib func(int) int64

	recurseFib = func(n int) int64 {
		if n == 1 || n == 2 {
			return 1
		}
		if _, ok := cache[n]; !ok {
			cache[n] = recurseFib(n-1) + recurseFib(n-2)
		}
		return cache[n]
	}

	// LEARNING TODO: what's the difference between returning a value (as done below)
	// as opposed to recurseFib, i.e the function itself (as done in the forum linked above) ??
	return recurseFib(n)

}

func main() {

	// a := []int32{1, 2, 3, 4}

	// fmt.Println(sum2(a, 0))

	// fmt.Println(product(a, 1))
	for i := 40; i < 50; i++ {
		fmt.Println("i = ", i, "fibonacci number = ", fibonacci(i))
	}

}
