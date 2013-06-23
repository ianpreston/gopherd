package main

import (
	"net"
)

type Server struct {
	listener net.Listener	
}

func NewServer(bindTo string) *Server {
	l, err := net.Listen("tcp", bindTo)
	if err != nil {
		return nil
	}

	return &Server{ l }
}

func (srv *Server) Run() {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			panic(err)
			conn.Close()
			continue
		}

		go srv.Handle(conn)
	}
}

func (srv *Server) Handle(conn net.Conn) {
	cli := NewClient(conn)
	cli.Handle()
}