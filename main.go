package main

import (
	"github.com/willboland/website/internal/server"
	"log"
)

func main() {
	s := server.NewServer()
	err := s.ListenAndServe(":9000")
	if err != nil {
		log.Fatalf("server failed to run: %s", err)
	}
}
