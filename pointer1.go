package main

import "fmt"

func add_n(xp *int, val int) {
	*xp = *xp + val
}

func square(a *int) {
	*a = *a * *a
}

func swap(a *int, b *int) {
	c := *b
	*b = *a
	*a = c

}

func main() {
	x := 22
	pp := new(int)

	fmt.Println(x, *pp)

	add_n(&x, 20)
	add_n(pp, 10)
	fmt.Println(x, *pp)

	swap(&x, pp)
	fmt.Println(x, *pp)

	square(&x)
	square(pp)
	fmt.Println(x, *pp)
}
