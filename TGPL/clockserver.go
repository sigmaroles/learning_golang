package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println("Client connected??", conn.RemoteAddr()) // just for fun.. print remote IP to terminal whenever a client connects
		go handleConn(conn)                                  // by creating a goroutine, each client gets own thread. otherwise only one client at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("It is currently 2 Jan 2006, clock 15:04:05 at timezone MST\n"))
		/*side note: the format string _has_ to be a representation of the time Jan 2 15:04:05 2006 MST
		and NOT an arbitrary time. imagine the degree of cool-ness here
		just put any weird combination and the language will work it out!!
		*/
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}
		time.Sleep(1 * time.Second)
	}
}
