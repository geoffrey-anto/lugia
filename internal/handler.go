package server

import (
	"fmt"
	"log/slog"
	"net"
)

func HandleRequest(conn net.Conn) {
	// TODO: Handle the request from client
	for {
		buf := make([]byte, 1)

		n, err := conn.Read(buf)

		if err != nil {
			slog.Error("Failed to get data!")
		}

		fmt.Printf("Message From Client %s", string(buf[:n]))
	}
}
