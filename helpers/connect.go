package helpers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DBConnect() (context.Context, *pgx.Conn) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://neocipher:kaleab@localhost:5432/neocipher")
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected sucessfully!")
	return ctx, conn
}
