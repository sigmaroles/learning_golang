package main

import "fmt"

// source: https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/

type FibFunc func() int

func makeFibGen() FibFunc {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f2 + f1), f2 //clever, no third variable needed
		return f1
	}
	/* returns an anonymous function that "traps" f1 and f2 within it (essence of closure).
	this means that in future, whenever the returned anonymous function is called, it has its
	own copy of f1 and f2, which it mutates.
	also, the anonymous function has to return a value, so that callers are able to
	print the fibonnaci numbers

	অর্থাৎ, অতি গভীর রহস্য। এর থেকে পাইথন অনেক সোজা :D
	*/

}

func main() {

	fg := makeFibGen()
	fg2 := makeFibGen()
	fg()

	for i := 0; i < 15; i++ {
		fmt.Println(fg(), fg2())
	}
}
