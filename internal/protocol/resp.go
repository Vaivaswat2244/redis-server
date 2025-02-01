package protocol

import (
	"net"
)

type Command struct {
	Name string
	Args []string
}

func ReadCommand(conn net.Conn) (*Command, error) {
	// RESP protocol parsing implementation
}

func WritePong(conn net.Conn) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	return err
}

func WriteString(conn net.Conn, s string) error {
	_, err := conn.Write([]byte("+" + s + "\r\n"))
	return err
}
