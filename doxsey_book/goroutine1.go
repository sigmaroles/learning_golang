package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 1; i < 5; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	}
	fmt.Println("####")
}

func main() {
	for i := 23; i < 29; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
