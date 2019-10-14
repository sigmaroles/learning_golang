package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name string
	age  int
}

type ByName []Student
type ByAge []Student

// ------------------------

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].age < ps[j].age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// ------------------------

func (ps ByName) Len() int {
	return len(ps)
}
func (ps ByName) Less(i, j int) bool {
	return ps[i].name < ps[j].name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// ------------------------

func main() {
	aclass := []Student{
		{"Rambo", 12},
		{"JovwwsxAzxxsw", 8},
		{"Lebowski", 23},
	}

	fmt.Println(aclass)
	/*
	   unusual way to sort data - potentially very powerful too.
	   sort.Sort is an interface. requires Len, Less and Swap to be defined
	   on the custom slice type (here, ByName and ByAge)
	*/
	sort.Sort(ByName(aclass))
	fmt.Println(aclass)

	sort.Stable(ByAge(aclass))
	fmt.Println(aclass)

}

/*
func main() {
	var x list.List

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	fmt.Println(x.Front().Value)

	x.PushBack(1)
	x.PushBack(23)
	x.PushFront(122)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}*/
