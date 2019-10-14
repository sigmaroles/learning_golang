package main

import "fmt"

func main() {
	/* 		Nested maps; {string : {string: string}} */
	// the python equivalent is a dict of dicts

	elem := map[string]map[string]string{
		"H":  map[string]string{"name": "Hydrogen", "state": "gas"},
		"C":  map[string]string{"name": "Carbon", "state": "solid"},
		"N":  map[string]string{"name": "Nitrogen", "state": "gas"},
		"Hg": map[string]string{"name": "Mercury", "state": "liquid"},
	}
	// syms := []string{"H", "C"}

	//for _, sym := range syms {

	for key, val := range elem {
		// note that val is itself a map
		fmt.Println(key + " stands for " + val["name"] + " and is in state " + val["state"])
	}
}
