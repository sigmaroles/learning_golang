package main

import "fmt"

type Mammal interface {
	sayHello()
}

type Human struct {
	name string
}

type Cat struct {
	name string
}

func (c *Cat) sayHello() {
	fmt.Println("Hello, my name is", c.name)
}

func (h *Human) sayHello() {
	fmt.Println("My name is", h.name)
}

/*
The rest of the code is boilerplate. Notice how greet works.
It assumes each element in entities 'implements' the Mammal interface. I
In golang that means each element in entities must be a struct that has all the methods
of Mammal interface defined. In this case, that means Human and Cat must have
a sayHello() func (with exact same signature)
*/
func greet(entities ...Mammal) {
	for _, e := range entities {
		e.sayHello()
	}
}

func main() {
	h1 := new(Human)
	h1.name = "Python"
	c1 := new(Cat)
	c1.name = "Whiskey"

	greet(c1, h1)

	/*
		output:
		Hello, my name is Whiskey
		My name is Python

	*/
}

/*
from the book:

Interfaces are particularly useful as software projects grow and become more com‐
plex. They allow us to hide the incidental details of implementation (e.g., the fields of
our struct), which makes it easier to reason about software components in isolation.
In our example, as long as the area methods we defined continue to produce the
same results, we’re free to change how a Circle or Rectangle is structured without
having to worry about whether or not the totalArea function will continue to work.
*/
