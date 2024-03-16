package database

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/francoggm/rinhabackend-2024-q1-go-psql/domain/client"
)

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *Database) GetExtract(ctx context.Context, id int) (*client.Extract, error) {
	rows, err := d.db.Query(ctx, ExtractQuery, id)

	if err != nil {
		return nil, err
	}

	var info client.ExtractClientInfo
	var transactions = make([]client.ExtractTransaction, 0)

	for rows.Next() {
		var (
			value       sql.NullInt32
			description sql.NullString
			ttype       sql.NullString
			createdAt   sql.NullTime
		)

		err := rows.Scan(&info.Balance, &info.Limit, &value, &description, &ttype, &createdAt)
		if err != nil || !value.Valid || !description.Valid || !ttype.Valid || !createdAt.Valid {
			continue
		}

		transactions = append(transactions, client.ExtractTransaction{
			Value:       int(value.Int32),
			Description: description.String,
			Type:        ttype.String,
			CreatedAt:   createdAt.Time,
		})
	}

	if info.Balance == 0 && info.Limit == 0 {
		return nil, client.ErrNotFound
	}

	info.Date = time.Now()

	return &client.Extract{
		Info:         info,
		Transactions: transactions,
	}, nil
}

func (d *Database) MakeTransaction(ctx context.Context, id int, value int, description string, ttype string) (*client.TransactionRes, error) {
	var transaction client.TransactionRes

	row := d.db.QueryRow(ctx, TransactionQuery, id, value, description, ttype)
	err := row.Scan(&transaction.Balance, &transaction.Limit)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "client not found"):
			return nil, client.ErrNotFound
		case strings.Contains(err.Error(), "violates check constraint"):
			return nil, client.ErrInsufficientLimit
		default:
			return nil, err
		}
	}

	return &transaction, nil
}
