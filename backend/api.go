package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	e := echo.New()
	log.Println("API server running on port: " + s.listenAddr)

	e.GET("/", s.handleGetPosts)
	e.Logger.Fatal(e.Start(s.listenAddr))
}

func (s *APIServer) handleGetPosts(c echo.Context) error {
	posts, err := s.store.GetPosts()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("handle")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScriptCharsetUTF8)
	return c.JSON(http.StatusOK, posts)
}
