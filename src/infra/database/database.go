package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabase(username string, password string, hostname string, db string) (*pgxpool.Pool, error) {
	cfg, _ := pgxpool.ParseConfig(fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s sslmode=disabled", username, password, hostname, db))
	cfg.MaxConns = 8
	cfg.MinConns = 4

	return pgxpool.NewWithConfig(context.Background(), cfg)
}
