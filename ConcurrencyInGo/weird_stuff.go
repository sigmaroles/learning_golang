package main

import "fmt"

func main() {
	var data int
	for data < 60 {

		go func() {
			data++
		}()

		go fmt.Print(data, " : ")
	}
	fmt.Println("\nFinal value is ", data)
}
