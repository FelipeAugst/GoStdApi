package main

import (
	"api/src/server"
)

func main() {
	server := server.NewServer("localhost", ":5000")
	server.Start()

}
