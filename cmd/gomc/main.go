package main

import "gomc/internal/pkg/server"

func main() {
	s := server.NewServer()
	s.Start()
}
