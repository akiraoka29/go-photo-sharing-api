package main

import (
	"log"

	"github.com/akiraoka29/go-photo-sharing-api/pkg/server"
)

func main() {
	srv := server.NewServer()
	log.Fatal(srv.Run(":7070"))
}