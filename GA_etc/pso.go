package main

import "fmt"

type Particle struct {
	x, y float64
}

func main() {
	c := Particle{x: 1, y: 0}
	fmt.Println(c)
}
