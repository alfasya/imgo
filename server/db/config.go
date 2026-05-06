package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool
var Ctx = context.Background()

func Connect() {
	var err error

	Pool, err = pgxpool.New(Ctx, "postgresql://postgres:2121@localhost:5433/imgo")
	if err != nil {
		fmt.Printf("Unable to create pool connection: %v", err)
	}

	fmt.Println("Connected to database.")
}
