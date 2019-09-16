package main

import "fmt"

func main() {

	//var x map[string]int

	days := make(map[string]string)

	days["M"] = "Monday"
	days["T"] = "Tuesday"
	days["W"] = "Wednesday"
	days["Th"] = "Thursday"

	fmt.Println(days["W"])

	delete(days, "go")

	//	if name, ok := x["go"]; ok != nil {
	fmt.Println(days)
	//	}

}
