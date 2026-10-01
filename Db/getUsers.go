package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetUsers(ctx context.Context, conn *pgx.Conn) ([]Table, error) {

	sqlQuery := `
		SELECT id, full_name, phone_number 
		FROM users
		ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		fmt.Println("err arguments")
		return nil, err
	}

	defer rows.Close()

	tables := make([]Table, 0)

	for rows.Next() {
		var table Table

		err := rows.Scan(
			&table.ID,
			&table.FullName,
			&table.PhoneNumber,
		)

		if err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil

}
