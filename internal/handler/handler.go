package handler

import (
	"net"

	"github.com/yourusername/redis-server/internal/protocol"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		cmd, err := protocol.ReadCommand(conn)
		if err != nil {
			return
		}

		switch cmd.Name {
		case "PING":
			protocol.WritePong(conn)
		case "ECHO":
			protocol.WriteString(conn, cmd.Args[0])
		}
	}
}
