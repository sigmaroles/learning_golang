package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// a "method" without any classes. in fact go does not have classes
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) change_radius(newradius float64) {
	c.r = newradius
}

func main() {

	c := Circle{x: 1, r: 3, y: 0}
	fmt.Println(c)

	fmt.Println(c.area())
	c.change_radius(2)
	fmt.Println(c.area())

}
