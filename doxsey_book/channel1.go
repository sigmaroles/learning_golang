package main

import (
	"fmt"
	"strconv"
	"time"
)

/*

modified version of the code in page 89, Doxsey ebook
observations and ideas:
1. sending some useful data from pinger (strconv of i)
2. the variable name of sent and recrived need not match
3. the infinite loop is distracting
4. pressing enter before it has exeuted all the iterations inside pinger leads to program end
5. where would this even be useful?? message passing mechanism, alt to pub-sub?
*/

func pinger(sentmsg chan string) {
	for i := 0; i < 10; i++ {
		sentmsg <- "ping " + strconv.Itoa(i)
	}
}

func printer(recvmsg chan string) {
	for {
		msg := <-recvmsg
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 400)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
