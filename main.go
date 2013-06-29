package main

func main() {
	sc := &ServerConfig{
		":7070",
		"localhost",
		7070,
		"/Users/ian/gopherroot",
	}
	server := NewServer(sc)
	server.Run()
}