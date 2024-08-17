package initializers

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DbConnection *pgx.Conn

func InitDatabase(ctx context.Context) error {
	// ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		// Decide whether to return the error or continue with default values
		// For now, we'll just print the error and continue
	}

	var err error
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	DbConnection, err = pgx.Connect(ctx, databaseURL)
	// DbConnection, err = pgx.Connect(ctx, "user=postgres password=password host=localhost port=5432 dbname=test sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return err
	}

	return nil
}
