package main

import (
	"flag"
)

func main() {
	flag.Parse()
	confPath := flag.Arg(0)
	if confPath == "" {
		confPath = "config.json"
	}
	conf := LoadJsonConfig(confPath)

	server := NewServer(conf)
	server.Run()
}