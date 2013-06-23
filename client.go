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
}

func NewClient(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)

	return &Client { conn, reader }
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
	fmt.Fprintf(cli.conn, "iSpecified Selector: '" + selector + "'\t\tlocalhost\t7070\r\n")
	fmt.Fprintf(cli.conn, "1Not a real directory (/dev/null)\t/dev/null\tlocalhost\t7070\r\n")
	fmt.Fprintf(cli.conn, "1Not a real directory (root)\t\tlocalhost\t7070\r\n")
	fmt.Fprintf(cli.conn, "0Using web browsers in Gopherspace\tgopher/wbgopher\tgopher.floodgap.com\t70\r\n")
	fmt.Fprintf(cli.conn, ".")
}