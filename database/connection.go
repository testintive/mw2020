package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

var dbPool *pgxpool.Pool

// InitDb initialises pgx pool to use across app
func InitDb() error {
	var err error
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))

	dbPool, err = pgxpool.Connect(context.Background(), connString)

	if err != nil {
		log.Println("Unable to connect to database")
		dbPool = nil
	}

	return err
}

func GetConnection() (*pgxpool.Conn, error) {
	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, err
}
