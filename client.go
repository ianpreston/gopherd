package main

import (
	"fmt"
	"strings"
	"net"
	"bufio"
)

type Client struct {
	conn net.Conn
	reader *bufio.Reader
	conf *ServerConfig
}

func NewClient(conn net.Conn, conf *ServerConfig) *Client {
	reader := bufio.NewReader(conn)

	return &Client { conn, reader, conf }
}

func (cli *Client) Handle() {
	line, err := cli.reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		cli.conn.Close()
		return
	}

	selector := strings.Trim(line, "\r\n ")

	cli.HandleRequest(selector)

	cli.conn.Close()
}

func (cli *Client) HandleRequest(selector string) {
	fmt.Println("Handling request for: " + selector)
	req := NewRequest(cli)
	req.Handle(selector)
}