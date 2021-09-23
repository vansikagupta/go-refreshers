/*
This Clock server listens and accepts client requests at port 8000
and sends them the current time every second.

Represents concurrency
*/

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	ls, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue       // accept other requests
		}
		handleconn(conn)
	}
}

//
func handleconn(conn net.Conn) {
	log.Print("Accepted Client connection")
	// blocking for a client connection
	for {
		// conn implements io.Writer interface
		_, err := io.WriteString(conn, "Hello World!")
		if err != nil {
			log.Print(err)
			return // client disconnected
		}
		time.Sleep(time.Second * 1)
	}
}
