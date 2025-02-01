package handler

import (
	"bufio"
	"fmt"
	"net"

	"github.com/Vaivaswat2244/redis-server/internal/protocol"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		fmt.Println("Waiting for command...")
		cmd, err := protocol.ReadCommand(reader)
		if err != nil {
			fmt.Printf("Error reading command: %v\n", err)
			return
		}
		fmt.Printf("Received command: %+v\n", cmd)

		switch cmd.Name {
		case "PING":
			protocol.WritePong(conn)
		case "ECHO":
			if len(cmd.Args) > 0 {
				protocol.WriteString(conn, cmd.Args[0])
			}
		}
	}
}
