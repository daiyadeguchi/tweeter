package main

import (
	"database/sql"
	"fmt"
	"github.com/daiyadeguchi/tweeter/backend/types"
	_ "github.com/lib/pq"
	"log"
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
	GetPosts() ([]*types.Post, error)
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

func (s *PostgresStore) Init() error {
	err := s.createAccountTable()
	err = s.createPostTable()
	return err
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
			id serial primary key, 
			username varchar(100), 
			email varchar(100), 
			encrypted_password varchar(100)
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createPostTable() error {
	query := `create table if not exists post (
			id serial primary key, 
			user_id integer, 
			post_body varchar(500)
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *types.Account) error {
	panic("implement me")
}

func (s *PostgresStore) DeleteAccount(i int) error {
	panic("implement me")
}

func (s *PostgresStore) UpdateAccount(account *types.Account) error {
	panic("implement me")
}

func (s *PostgresStore) GetAccountByID(i int) (*types.Account, error) {
	panic("implement me")
}

func (s *PostgresStore) CreatePost(post *types.Post) error {
	rows, err := s.db.Query("insert into post (user_id, post_body) values ($1, $2)", post.UserID, post.Body)
	if err != nil {
		return err
	}
	log.Println("Successfully inserted:", rows)
	return nil
}

func (s *PostgresStore) DeletePost(post *types.Post) error {
	panic("implement me")
}

func (s *PostgresStore) UpdatePost(post *types.Post) error {
	panic("implement me")
}

func (s *PostgresStore) GetPostByID(i int) (*types.Post, error) {
	panic("implement me")
}

func (s *PostgresStore) GetPosts() ([]*types.Post, error) {
	rows, err := s.db.Query("select * from post")
	if err != nil {
		return nil, err
	}

	var posts []*types.Post
	for rows.Next() {
		post := new(types.Post)
		err := rows.Scan(&post.ID, &post.UserID, &post.Body)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
