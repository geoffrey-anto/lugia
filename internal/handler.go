package server

import (
	"fmt"
	"log/slog"
	"net"
)

func HandleRequest(conn net.Conn, s *Server) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			conn.Close()
			return
		}

		message := string(buf)[:n]

		fmt.Println(s.Rf.Add("1", 0, 2))
		fmt.Printf("%+v\n", s.Rf)
		fmt.Println(s.Rf.Update("1", 1))
		fmt.Printf("%+v\n", s.Rf)
		fmt.Println(s.Rf.Update("1", 2))
		fmt.Printf("%+v\n", s.Rf)
		fmt.Println(s.Rf.End("1"))

		if message == "END" {
			fmt.Printf("Closing connection with %+v\n", conn.RemoteAddr().String())
			err = conn.Close()
			if err != nil {
				slog.Error("Failed to disconnect!")
			}
			return
		}

		queryType := ParseQuery(message)

		_, err = conn.Write([]byte(queryType))

		if err != nil {
			slog.Error("Failed to send response!")
			conn.Close()
			return
		}
	}
}
