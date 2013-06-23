package main

func main() {
	server := NewServer(":7070")
	server.Run()
}