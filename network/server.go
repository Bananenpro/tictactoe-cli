package network

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type ServerConnection struct {
	con net.Conn
}

func Connect(address string) (*ServerConnection, error) {
	con, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &ServerConnection{
		con: con,
	}, nil
}

// waits for 'match-found' message from server and returns the assigned sign ('x'/'o')
func (c *ServerConnection) WaitForMatch() (string, error) {
	scanner := bufio.NewScanner(c.con)
	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		if strings.HasPrefix(text, "match-found:") {
			parts := strings.Split(text, ":")
			if len(parts) == 2 {
				sign := parts[1]
				if sign != "x" && sign != "o" {
					fmt.Println("Invalid sign:", sign)
				} else {
					return sign, nil
				}
			}
		} else {
			fmt.Println("Unexpected message from server:", text)
		}
	}
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	return "", errors.New("disconnected")
}

func (c *ServerConnection) Close() error {
	return c.con.Close()
}
