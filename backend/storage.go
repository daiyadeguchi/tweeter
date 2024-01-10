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
	DeletePostByID(int) error
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
	_, err := s.db.Query("insert into account (username, email, encrypted_password) values ($1, $2, $3)", account.Username, account.Email, account.EncryptedPassword)
	if err != nil {
		return err
	}
	log.Println("Account successfully created: ", account.Username)
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	if err != nil {
		return nil
	}
	log.Println("Account successfully deleted: ", id)
	return nil
}

func (s *PostgresStore) UpdateAccount(account *types.Account) error {
	_, err := s.db.Query("update account set username = $1, email = $2, encrypted_password = $3 where id = $4", account.Username, account.Email, account.EncryptedPassword, account.ID)
	if err != nil {
		return nil
	}
	log.Println("Account successfully updated: ", account.Username)
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*types.Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		account := new(types.Account)
		err := rows.Scan(&account.ID, &account.Username, &account.Email, &account.EncryptedPassword)
		if err != nil {
			return nil, err
		}
		return account, nil
	}
	return nil, nil
}

func (s *PostgresStore) CreatePost(post *types.Post) error {
	_, err := s.db.Query("insert into post (user_id, post_body) values ($1, $2)", post.UserID, post.Body)
	if err != nil {
		return err
	}
	log.Println("Successfully inserted: ", post.UserID)
	return nil
}

func (s *PostgresStore) DeletePostByID(id int) error {
	_, err := s.db.Query("delete from post where id = $1", id)
	if err != nil {
		return err
	}
	log.Println("Successfully deleted: ", id)
	return nil
}

func (s *PostgresStore) UpdatePost(post *types.Post) error {
	_, err := s.db.Query("update post set post_body = $1 where id = $2", post.Body, post.ID)
	if err != nil {
		return err
	}
	log.Println("Successfully updated", post.ID)
	return nil
}

func (s *PostgresStore) GetPostByID(id int) (*types.Post, error) {
	rows, err := s.db.Query("select * from post where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := new(types.Post)
		err := rows.Scan(&post.ID, &post.UserID, &post.Body)
		if err != nil {
			return nil, err
		}
		return post, nil
	}
	return nil, nil
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
