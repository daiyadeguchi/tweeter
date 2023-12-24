package main

import (
	"database/sql"
	"fmt"
	"github.com/daiyadeguchi/tweeter/backend/types"
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "pass123"
	dbname   = "tweeter"
)

type Storage interface {
	CreateAccount(*types.Account) error
	DeleteAccount(int) error
	UpdateAccount(*types.Account) error
	GetAccountByID(int) (*types.Account, error)
	CreatePost(*types.Post) error
	DeletePost(*types.Post) error
	UpdatePost(*types.Post) error
	GetPostByID(int) (*types.Post, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresConnection() (*PostgresStore, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	// Check the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to postgres")
	return &PostgresStore{
		db: db,
	}, nil
}
