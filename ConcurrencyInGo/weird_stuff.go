package main

import "fmt"

func main() {
	data := 0
	iter := 0
	for data < 10 {

		go func() {
			fmt.Println("Plus func; got data = ", data)
			data++
		}()

		go func() {
			fmt.Println("Minus func; got data = ", data)
			data--
		}()

		iter++
		fmt.Println("iter = ", iter, " data : ", data)
	}
	fmt.Println("\nFinal value is ", data)
}
