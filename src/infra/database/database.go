package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

func NewDatabase(username string, password string, hostname string, db string) (*Database, error) {
	cfg, _ := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=disable", username, password, hostname, db))
	cfg.MaxConns = 8
	cfg.MinConns = 4

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	return &Database{
		db: pool,
	}, err
}
