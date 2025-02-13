package database

import (
	"context"
	"erp-api/util/configuration"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var Conn *pgxpool.Pool

func InitDB(){
	logger := configuration.Logger()
	
	err := godotenv.Load()
  if err != nil {
    logger.Fatal("Error loading .env file")
  }

	Conn, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Sugar().Fatalf("Failed to connect to the database: %v", err)
	}

	// Example query to test connection
	var version string
	if err := Conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		logger.Sugar().Fatalf("Query failed: %v", err)
	}

	logger.Sugar().Infof("Connected to:", version)
}