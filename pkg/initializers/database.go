package initializers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var DbConnection *pgx.Conn

func InitDatabase(ctx context.Context) error {
	// ctx := context.Background()

	var err error
	DbConnection, err = pgx.Connect(ctx, "user=postgres password=password host=localhost port=5432 dbname=test sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return err
	}

	return nil
}
