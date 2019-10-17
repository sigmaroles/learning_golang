package main

import (
	"fmt"
	"math/rand"
	"time"
)

// also look at https://www.calhoun.io/what-is-a-closure/

/*
The book's example was too trivial, so used this page:
https://tleyden.github.io/blog/2016/12/20/understanding-function-closures-in-go/
*/

// similar to typedef in c++
// SenderF is now a type that defines a function with no arguments and returns a bool
type SenderF func() bool

/*
	One thing to note here is the clean separation of concerns —
	all sendLoop() knows is that it gets a SenderF which it should call and it will return
	a boolean indicator of whether or not it worked or not. It knows absolutely nothing about
	the inner workings of the SenderF, nor does it care.
*/

/*
chanegd the sendloop to better reflect the strategy:
keep trying until receive a success message from the SenderF function object (sender)
sleep in between retries for specific amount of time
*/
func sendloop(sender SenderF) {
	for {
		success := sender()
		if success {
			break
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {

	counter := 0
	maxAttempts := 20000

	/*
		take no arguments, "sends" predefined message, returns boolean (status)
		keeps track of how many times it has been called -- crucial aspect
	*/
	mySender := func() bool {
		success := rand.Intn(10) // message "sent" or not
		// fmt.Print("\nrandInt = ", success, "\n")
		if success == 0 {
			return true
		}
		fmt.Println("was not able to send the message yet")
		counter += 1
		return counter > maxAttempts
	}

	/*
	   The sendLoop() doesn’t know anything about the internals of the SenderF
	   in terms of how it tracks whether or not it should retry or not, it just treats
	   it as a black box. Different SenderF implementations could use vastly different
	   rules and/or states for deciding whether the sendLoop() should retry a failed send.
	*/

	sendloop(mySender)
}
