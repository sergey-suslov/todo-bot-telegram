package db

import (
	"os"
	"context"
	"github.com/jackc/pgx"
)


// GetConnection Get connection to DB
func GetConnection() (*pgx.Conn, error) {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	conn, err := pgx.Connect(context.Background(), "postgresql://" + username + ":" + password + "@" + host + ":" + port)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
