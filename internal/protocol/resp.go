package protocol

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

func ReadCommand(reader *bufio.Reader) (Command, error) {
	peek, _ := reader.Peek(32) // Peek at first 32 bytes
	fmt.Printf("Received data: %q\n", peek)

	// Read the first byte to determine the type
	firstByte, err := reader.ReadByte()
	if err != nil {
		return Command{}, err
	}

	// Check if it's an array (command with arguments)
	if firstByte != '*' {
		return Command{}, fmt.Errorf("expected '*', got '%c'", firstByte)
	}

	// Read the number of elements in the array
	countStr, err := reader.ReadString('\n')
	if err != nil {
		return Command{}, err
	}
	countStr = strings.TrimSpace(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return Command{}, err
	}

	// Initialize command with capacity for all arguments
	cmd := Command{
		Args: make([]string, 0, count-1),
	}

	// Read each element
	for i := 0; i < count; i++ {
		// Read the $ character
		b, err := reader.ReadByte()
		if err != nil {
			return Command{}, err
		}
		if b != '$' {
			return Command{}, fmt.Errorf("expected '$', got '%c'", b)
		}

		// Read the length
		lenStr, err := reader.ReadString('\n')
		if err != nil {
			return Command{}, err
		}
		lenStr = strings.TrimSpace(lenStr)
		strLen, err := strconv.Atoi(lenStr)
		if err != nil {
			return Command{}, err
		}

		// Read the actual string
		str := make([]byte, strLen+2) // +2 for \r\n
		_, err = io.ReadFull(reader, str)
		if err != nil {
			return Command{}, err
		}

		// Remove \r\n and store the string
		value := string(str[:strLen])

		if i == 0 {
			// First element is the command name
			cmd.Name = strings.ToUpper(value)
		} else {
			// Rest are arguments
			cmd.Args = append(cmd.Args, value)
		}
	}

	return cmd, nil
}

func WritePong(writer io.Writer) error {
	_, err := writer.Write([]byte("+PONG\r\n"))
	return err
}

func WriteString(writer io.Writer, s string) error {
	_, err := writer.Write([]byte("+" + s + "\r\n"))
	return err
}
