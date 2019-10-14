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


func main() {

	a := []int32{1, 2, 3, 4}

	fmt.Println(sum2(a, 0))

	fmt.Println(product(a, 1))

}
