package server

import (
	"fmt"
	"net"
	"os"
)

type Server struct {
	Id string

	Lis net.Listener
}

func NewServer(addr string, port int, id string) *Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))

	if err != nil {
		fmt.Printf("Failed to access port %d!", port)
		os.Exit(1)
	}

	return &Server{
		Id: id,

		Lis: lis,
	}
}

func (s *Server) Start() {
	fmt.Printf("Server ready to accept connections @ %+v\n", s.Lis.Addr().String())
	for {
		if conn, err := s.Lis.Accept(); err == nil {
			fmt.Printf("New Connection from %v\n", conn.RemoteAddr().String())

			go HandleRequest(conn)
		}
	}
}
