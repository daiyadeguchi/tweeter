package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "pass123"
	dbname   = "tweeter"
)

func ConnectPostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// Check the connection
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected")
	return db
}
