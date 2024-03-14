package database

import "context"

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) Ping(ctx context.Context) error {
	return d.db.Ping(ctx)
}

func (d *Database) GetExtract(id int) {}

func (d *Database) MakeTransaction() {}
