package initializers

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var DbConnection *pgx.Conn

func InitDatabase() error {
	ctx := context.Background()

	var err error
	DbConnection, err = pgx.Connect(ctx, "user=postgres password=password host=localhost port=5432 dbname=test sslmode=disable")
	if err != nil {
		return err
	}

	return nil
}
