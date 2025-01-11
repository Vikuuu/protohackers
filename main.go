package main

import (
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen on TCP post 8080 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Make a buffer to hold incoming data.
			buf := make([]byte, 1024)
			// Read the incoming connection into the buffer.
			_, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			// Send a response back to person contacting us.
			conn.Write([]byte(buf))
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
