package main

import (
	"fmt"
	"math/big"
)

func main() {
	aa := big.NewInt(2)
	fmt.Println("The int is ", aa)
	fmt.Println("The sum is ", testf(a))
}

func testf(a int) big.Int {
	a1 := big.NewInt(a)
	return big.NewInt(a1 + a1*3)
}
