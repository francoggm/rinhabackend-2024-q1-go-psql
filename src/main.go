package main

import (
	"log"

	"github.com/francoggm/rinhabackend-2024-q1-go-psql/server"
)

func main() {
	if err := server.Run("8081"); err != nil {
		log.Panic(err)
	}
}
