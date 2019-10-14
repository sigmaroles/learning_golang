package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("chamber_of_secrets")
	if err != nil {
		log.Fatal("you are going down!\n" + string(err.Error()))
		log.Fatal(err)
	}
	// do something with the open *File f
	fmt.Println(f.Name())

}
