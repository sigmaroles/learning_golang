package main

import "fmt"

/* write a program to find the smallest number from a list
an exercise (from [Caleb_Doxsey]_Introducing_Go.pdf) that I actually did, after about 10 syntax errors
and one logical error on 9Feb2019 5:48pm
*/

func main() {

	var s int
	x := []int{
		23, 11, 78, 32, 12, 7, 33, 45, 13,
	}

	s = x[0]

	for _, elem := range x {
		if elem < s {
			s = elem
			//fmt.Println(s, "!!!")
		}
	}

	fmt.Println(s)
}
