package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/greg-beach/httpfromtcp/internal/request"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listeniing for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		reader := bufio.NewReader(conn)

		request, err := RequestFromReader(reader)
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		for line := range linesChan {
			fmt.Println(line)
		}

		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
	}

}
