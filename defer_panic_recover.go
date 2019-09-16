package main

import "fmt"

func first() {
	fmt.Println("First function, huh")

}

func second() {
	fmt.Println("second")

}

func third() {
	fmt.Println("third??")
}

func main() {
	defer third()
	defer third()
	defer second()
	defer third()
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("whaaa -- this is a panic")

	first()
	first()
}
