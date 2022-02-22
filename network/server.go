package network

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

var (
	ErrDisconnected = errors.New("disconnected")
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

type ReadCallback func(con *ServerConnection, text string)

func (c *ServerConnection) Read(callback ReadCallback) {
	scanner := bufio.NewScanner(c.con)
	for scanner.Scan() {
		callback(c, strings.ToLower(scanner.Text()))
	}
	if scanner.Err() != nil {
		callback(c, "error: "+scanner.Err().Error())
	} else {
		callback(c, "disconnect")
	}
}

func (c *ServerConnection) ReadLine() (string, error) {
	scanner := bufio.NewScanner(c.con)
	if scanner.Scan() {
		fmt.Println(scanner.Text())
		return strings.ToLower(scanner.Text()), nil
	}
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	return "", ErrDisconnected
}

func (c *ServerConnection) ClickField(fieldIndex int) error {
	return c.Send(fmt.Sprintf("click:%d", fieldIndex))
}

func (c *ServerConnection) Close() error {
	return c.con.Close()
}

func (c *ServerConnection) Send(text string) error {
	_, err := fmt.Fprintf(c.con, text+"\n")
	if err != nil {
		fmt.Printf("Failed to send '%s': '%s'\n", text, err)
		return err
	}
	return nil
}
