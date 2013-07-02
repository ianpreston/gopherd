package main

import (
	"net"
)

type Server struct {
	listener net.Listener
	conf *ServerConfig
}

func NewServer(conf *ServerConfig) *Server {
	l, err := net.Listen("tcp", conf.BindTo)
	if err != nil {
		return nil
	}

	return &Server{ l, conf }
}

func (srv *Server) Run() {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			conn.Close()
			continue
		}

		go srv.Handle(conn)
	}
}

func (srv *Server) Handle(conn net.Conn) {
	cli := NewClient(conn, srv.conf)
	cli.Handle()
}