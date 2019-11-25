package main

import "fmt"

func main() {
	data := 100
	for data > 90 {

		go func() {
			fmt.Println("Plus func; got data = ", data)
			data++
		}()

		go func() {
			fmt.Println("Minus func; got data = ", data)
			data--
		}()

		fmt.Println("State of data : ", data)
	}
	fmt.Println("\nFinal value is ", data)
}
