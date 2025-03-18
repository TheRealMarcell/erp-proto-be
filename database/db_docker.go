//go:build docker
// +build docker

package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var DB *pgxpool.Pool

func InitDB(logger zap.Logger) {
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if dbUser == "" || dbPass == "" || dbAddress == "" || dbPort == "" || dbName == "" {
		logger.Fatal("Database environment variables are not set")
	}

	databaseUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPass, dbAddress, dbPort, dbName)

	config, err := pgxpool.ParseConfig(databaseUrl)

	if err != nil {
		logger.Fatal("Failed to parse config")
	}

	config.MaxConns = 10
	config.MinConns = 5

	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Fatal("Failed to create connection pool")
	}

	logger.Sugar().Infof("Connected to database at URL: %v", databaseUrl)

	createTables()
}
