package main

import (
	"log"
)

func main() {
	store, err := NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":1323", *store)
	server.Run()
}

//
//func postGetHandler(c echo.Context, store *db.PostgresStore) echo.HandlerFunc {
//	var id string
//	var body string
//	rows, err := .Query("SELECT * FROM posts")
//	if err != nil {
//		c.Logger().Fatal(err)
//		return err
//	}
//	for rows.Next() {
//		if err := rows.Scan(&id, &body); err != nil {
//			c.Logger().Fatal(err)
//			return err
//		}
//		fmt.Println(id, body)
//	}
//	return c.String(http.StatusOK, body)
//}
