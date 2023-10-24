package main

import (
	"log"

	"server/infra/db"
	"server/router"

	"github.com/quic-go/quic-go/http3"
)

var DB *db.DBService

func init() {
	DB = db.Connect()
}

func main() {
	router := router.SetupRouter()

	go func() {
		if err := router.RunTLS(":8080", "./../localhost.crt", "./../localhost.key"); err != nil {
			log.Fatal("Failed to run Http/2 server")
		}
	}()

	server := &http3.Server{
		Addr:    ":8081",
		Handler: router,
	}

	if err := server.ListenAndServeTLS("./../localhost.crt", "./../localhost.key"); err != nil {
		log.Fatal("Failed to run Http/3 server")
	}
}
