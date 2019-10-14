package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(strings.Count("sohhom", "h"))
	fmt.Println(strings.HasSuffix("concatenation", "ion"))
	fmt.Println(strings.Split("a,b,c,d,e", ",")[2:])

	//var buf bytes.Buffer
	//buf.Write([]byte("wtf"))

	file, err := os.Open("txtt")
	if err != nil {
		// wtf happened?
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}
	// initialize a byte slice of same size as the size of file
	bs := make([]byte, stat.Size())
	//read the entire file into the slice
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	fmt.Println(string(bs))
}
