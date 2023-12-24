package main

import (
	"database/sql"
	"fmt"
	"github.com/daiyadeguchi/tweeter/backend/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	postgres := db.ConnectPostgres()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return postGetHandler(c, postgres)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func postGetHandler(c echo.Context, db *sql.DB) error {
	var id string
	var body string
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		c.Logger().Fatal(err)
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&id, &body); err != nil {
			c.Logger().Fatal(err)
			return err
		}
		fmt.Println(id, body)
	}
	return c.String(http.StatusOK, body)
}
