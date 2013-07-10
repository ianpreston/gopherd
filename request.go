package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"mime"
	"bufio"
	"strconv"
)

type Request struct {
	cli *Client
}

func NewRequest(cli *Client) *Request {
	return &Request{ cli };
}

func (req *Request) Handle(selector string) {
	selector = path.Clean(selector)
	physPath := path.Join(req.cli.conf.Root, selector)

	if strings.Index(physPath, req.cli.conf.Root) != 0 {
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
		if strings.Index(fi.Name(), ".") == 0 {
			continue
		}
		fmt.Fprintf(req.cli.conn, string(req.getPathByte(fi)) + fi.Name() + "\t" + fi.Name() + "\t" + req.cli.conf.Host + "\t" + strconv.Itoa(req.cli.conf.Port) + "\r\n")
	}

	fmt.Fprintf(req.cli.conn, ".")
}

func (req *Request) HandleFile(physPath string) {
	if strings.Index(path.Base(physPath), ".") == 0 {
		req.serveError("Failed to open file")
		return
	}

	file, err := os.Open(physPath)
	if (err != nil) {
		req.serveError("Failed to open file")
		return
	}

	buf := make([]byte, 4096)
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(req.cli.conn)

	for {
		_, err := reader.Read(buf)
		if (err != nil) {
			break
		}

		writer.Write(buf)
	}

	writer.Flush()
	file.Close()

	// If the file is text, we must send a '.' on its own line
	fi, _ := os.Stat(physPath)
	if (req.getPathByte(fi) == '0') {
		fmt.Fprintf(req.cli.conn, "\r\n.")
	}
}

func (req *Request) getPathByte(fi os.FileInfo) byte {
	ext := path.Ext(fi.Name())
	mimeType := mime.TypeByExtension(ext)

	if (fi.IsDir()) {
		return '1'
	}

	if (mimeType == "image/gif") {
		return 'g'
	}
	if (strings.Index(mimeType, "image/") == 0) {
		return 'I'
	}
	if (strings.Index(mimeType, "text/") == 0 ||
		strings.Index(mimeType, "json") != -1 ||
		strings.Index(mimeType, "xml") != -1) {
		return '0'
	}

	return '9'
}

func (req *Request) serveError(err string) {
	fmt.Fprintf(req.cli.conn, "3" + err + "\terror\terror\t0\r\n.")
}