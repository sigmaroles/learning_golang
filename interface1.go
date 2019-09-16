package main

import (
	"fmt"
	"log"
)

type FruitType interface {
	Wash() FruitType
	Eat() string
}

type Fruit struct {
	name  string
	dirty bool
	fruit FruitType
}

func (f *Fruit) Wash() FruitType {
	f.dirty = false
	if f.fruit != nil {
		return f.fruit
	}
	return f
}
func (f *Fruit) Eat() string {
	if f.dirty {
		return fmt.Sprintf("The %s is dirty, wash it first!", f.name)
	}
	return fmt.Sprintf("The %s is so delicious!", f.name)
}

type Orange struct {
	*Fruit
}

func NewOrange() *Orange {
	ft := &Orange{&Fruit{"Orange", true, nil}}
	ft.fruit = ft
	return ft
}
func NewApple() *Fruit {
	ft := &Fruit{"apple", true, nil}
	return ft
}

func (o *Orange) Eat() string {
	return "The orange is just ... orange"
}

func main() {
	apple1 := NewApple()
	log.Println(apple1.Eat())
	apple1.Wash()
	log.Println(apple1.Eat())

	orange1 := NewOrange()
	log.Println(orange1.Eat())
	orange1.Wash()
	log.Println(orange1.Eat())
}
