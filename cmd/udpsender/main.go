package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "localhost:42069"

func main() {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("error resolving UDP address: %s\n", err.Error())
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("error connecting to UDP address %s\n", err.Error())
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading string: %s\n", err.Error())
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatalf("error writing line: %s\n", err.Error())
		}
	}
}
