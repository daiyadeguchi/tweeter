package main

import (
	"log"
)

func main() {
	store, err := NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":1323", store)
	server.Run()
}
