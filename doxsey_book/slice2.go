package main

import "fmt"

func main() {

	sl1 := []int{3, 2, 1, 12, 87, 33, 67}
	sl2 := append(sl1, 112, 113, 114)

	fmt.Println(sl1)
	fmt.Println(sl2)

	sl3 := make([]int, 3)
	copy(sl3, sl2)
	fmt.Println(sl3)
}

/*
the book does not explain slices well enough. read the following for in depth, and bookmark for recap:

https://www.calhoun.io/how-to-use-slice-capacity-and-length-in-go/
https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
https://blog.golang.org/go-slices-usage-and-internals

key takeaways: slices are:
* less bug prone
* more "flexible and powerful"
* reference types...passed as reference to functions (arrays are copied)
*

*/
