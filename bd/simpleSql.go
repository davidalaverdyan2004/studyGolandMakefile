package bd

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS users (
	id  SERIAL PRIMARY KEY,
	full_name VARCHAR(100) NOT NULL,
	phone_number VARCHAR
	)
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}
