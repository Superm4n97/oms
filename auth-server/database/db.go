package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var DB *pgxpool.Pool

func ConnectDB() {
	connString, found := os.LookupEnv("DB_CONN")
	if !found {
		connString = "postgres://postgres:postgres@localhost:5432/postgres"
	}

	var err error
	DB, err = pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to the database.")
}

func CloseDB() {
	DB.Close()
}
