package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Request struct {
	cli *Client
	gopherRoot string
}

func NewRequest(cli *Client, gopherRoot string) *Request {
	gopherRoot = path.Clean(gopherRoot)
	return &Request{ cli, gopherRoot };
}

func (req *Request) Handle(selector string) {
	selector = path.Clean(selector)
	physPath := path.Join(req.gopherRoot, selector)

	if strings.Index(physPath, req.gopherRoot) != 0 {
		req.serveError("Invalid")
		return
	}

	fi, err := os.Stat(physPath)
	if err != nil {
		req.serveError("Invalid")
		return
	}

	if fi.IsDir() {
		req.HandleDirectory(physPath)
	} else {
		req.HandleFile(physPath)
	}
}

func (req *Request) HandleDirectory(physPath string) {
	dir, err := os.Open(physPath)
	if err != nil {
		req.serveError("Invalid")
		return
	}

	children, _ := dir.Readdir(-1)
	if err != nil {
		req.serveError("Invalid")
		return
	}

	for _, fi := range children {
		fmt.Fprintf(req.cli.conn, "0" + fi.Name() + "\t" + fi.Name() + "\tlocalhost\t7070\r\n")
	}

	fmt.Fprintf(req.cli.conn, ".")
}

func (req *Request) HandleFile(physPath string) {
	fmt.Fprintf(req.cli.conn, "This will be the content of '" + physPath + "'")
	fmt.Fprintf(req.cli.conn, "\r\n.")
}

func (req *Request) serveError(err string) {
	fmt.Fprintf(req.cli.conn, "3" + err + "\terror\terror\t0\r\n.")
}