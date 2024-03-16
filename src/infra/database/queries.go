package database

const (
	ExtractQuery = `
		SELECT 
			c.balance, 
			c.limit, 
			t.value, 
			t.description, 
			t.type,
			t.created_at
		FROM clients c
		LEFT JOIN transactions t ON t.client_id = c.id 
		WHERE c.id = $1
		ORDER BY t.created_at DESC
		LIMIT 10;
	`

	TransactionQuery = `SELECT * FROM make_transaction($1, $2, $3, $4)`
)
