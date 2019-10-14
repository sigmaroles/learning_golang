package main

import "fmt"

func main() {

	x := [4]float64{32, 1, 44, 5}

	var total float64

	for _, value := range x {
		total += value

	}

	fmt.Println(total / float64(len(x)))

}
