package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)           // what is this syntax?? why does a string need to be "made"?
	for _, url := range os.Args[1:] { // note: this is how you get cmdline args
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //the magic of channels.. will print only when the fetch function returns
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) { //directional channel
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// my guess: workaround to calculate the nbytes by "copying" the response body to something like /dev/null

	resp.Body.Close() //?? "don't leak resources"
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
