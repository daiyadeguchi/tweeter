package main

import (
	"github.com/labstack/echo/v4"
	"log"
)

type APIServer struct {
	listenAddr string
	store      PostgresStore
}

func NewAPIServer(listenAddr string, store PostgresStore) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	e := echo.New()

	log.Println("API server running on port: " + s.listenAddr)
	e.Logger.Fatal(e.Start(s.listenAddr))
}
