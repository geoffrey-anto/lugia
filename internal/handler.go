package server

import (
	"fmt"
	"log/slog"
	"net"
)

func HandleRequest(conn net.Conn, s *Server) {
	// TODO: Handle the request from client
	for {
		buf := make([]byte, 2048)

		n, err := conn.Read(buf)

		if err != nil {
			slog.Error("Failed to get data!")
		}

		fmt.Printf("Message From Client %s", string(buf[:n]))

		fmt.Printf("%+v %+v", s.Rf.Graph, s.Rf.NameToNodes)
	}
}
