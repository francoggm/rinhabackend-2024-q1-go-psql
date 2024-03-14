package main

import (
	"context"
	"log"
	"time"

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

	for i := 3; i > 0; i-- {
		if err := db.Ping(context.Background()); err == nil {
			break
		}

		time.Sleep(time.Duration(i) * time.Second)
	}

	if err := server.Run(cfg.Port, db); err != nil {
		log.Fatal(err)
	}
}
