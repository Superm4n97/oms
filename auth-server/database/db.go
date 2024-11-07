package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB() {
	connString := "postgres://username:password@localhost:5432/dbname"
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
