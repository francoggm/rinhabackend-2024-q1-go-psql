package database

const (
	Transaction = `SELECT * FROM make_transaction($1, $2, $3, $4)`
)
