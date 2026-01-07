package server

import (
	"fmt"
	"net"
)

type Server struct {
	closed bool
}

func Serve(port int) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, err
	}
	defer listener.Close()

	server := &Server{
		closed: false,
	}
	go server.Listen(listener)

	return server, nil
}

func (s *Server) Close() {
	s.closed = true
}

func (s *Server) Listen(listener net.Listener) {
	for {
		conn, err := listener.Accept()

		if s.closed {
			return
		}

		if err != nil {
			return
		}

		s.handle(conn)

	}
}

func (s *Server) handle(conn net.Conn) {
	out := []byte("HTTP/1.1 200 OK \r\nContent-Type: text/plain\r\n\r\nHello World!")
	conn.Write(out)
	conn.Close()
}
