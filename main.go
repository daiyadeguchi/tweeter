package main

import (
	"github.com/daiyadeguchi/tweeter/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	db.ConnectPostgres()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, mfers!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
