package main

import "fmt"

func raceTest() {
	data := 100
	iter := 0
	for data > 80 {

		go func() {
			//fmt.Println("Plus func; got data = ", data)
			data++
		}()

		go func() {
			//fmt.Println("Minus func; got data = ", data)
			data--
		}()

		iter++
		//fmt.Println("iter = ", iter, " data : ", data)
	}
	fmt.Println("\tFinal data = ", data, " ; No. of iterations = ", iter)
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Print("Race test iter ", i)
		raceTest()
	}

}
