package main

import (
	"context"
	"log"

	"github.com/francoggm/rinhabackend-2024-q1-go-psql/config"
	"github.com/francoggm/rinhabackend-2024-q1-go-psql/infra/database"
	"github.com/francoggm/rinhabackend-2024-q1-go-psql/infra/server"
)

func main() {
	cfg := config.NewConfig()

	db, err := database.NewDatabase(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(cfg.Port, db); err != nil {
		log.Fatal(err)
	}
}
