package main

import (
	"github.com/daiyadeguchi/tweeter/backend/types"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
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
	e.POST("/new-post", s.handleNewPost)
	e.Logger.Fatal(e.Start(s.listenAddr))
}

func (s *APIServer) handleGetPosts(c echo.Context) error {
	posts, err := s.store.GetPosts()
	if err != nil {
		log.Fatal(err)
		return err
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJavaScriptCharsetUTF8)
	return c.JSON(http.StatusOK, posts)
}

func (s *APIServer) handleNewPost(c echo.Context) error {
	userID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	post := types.Post{
		UserID: userID,
		Body:   c.QueryParam("body"),
	}
	if err := s.store.CreatePost(&post); err != nil {
		return err
	}
	return nil
}
