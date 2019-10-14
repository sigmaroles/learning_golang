package main

import "fmt"

func sum(inparr []float64) float64 {
	total := 0.0
	for _, val := range inparr {
		total += val
	}
	return total
}

func avg(inparr []float64) float64 {
	return sum(inparr) / float64(len(inparr))
}

func main() {
	zaa := []float64{22.4, 2.9, 1, 2, 4}

	fmt.Println(avg(zaa))

}
