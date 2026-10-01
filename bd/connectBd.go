package bd

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectBd(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, os.Getenv("CONN_STRING"))
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Успешное подключение")
	return conn
}
