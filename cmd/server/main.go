package main

import (
	"log"

	"github.com/yourusername/redis-server/internal/server"
)

func main() {
	srv := server.New(":6379")
	log.Fatal(srv.Start())
}
