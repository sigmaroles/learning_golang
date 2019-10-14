package main

import "fmt"

func main() {
	//var x []float64
	yx := make([]float64, 6, 15)
	eed := yx[0:7]

	fmt.Println(eed)

	yx[7] := 443

	fmt.Println(yx)
}
