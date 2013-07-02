package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type ServerConfig struct {
	BindTo string
	Host string
	Port int
	Root string
}

func LoadJsonConfig(path string) *ServerConfig {
	fc, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	conf := &ServerConfig{}
	err = json.Unmarshal(fc, conf)
	if err != nil {
		panic(err)
	}

	CleanConfig(conf)

	return conf
}

func CleanConfig(sc *ServerConfig) {
	sc.Root = path.Clean(sc.Root)
}