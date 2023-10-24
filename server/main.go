package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, HTTP/3!")
	})
	return r
}

func main() {
	router := setupRouter()

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	server := &http3.Server{
		Addr: ":8081",
		Handler: router,
	}

	if err := server.ListenAndServeTLS("./../localhost.crt", "./../localhost.key"); err != nil {
		log.Fatalf("Failed to run Http/3 server: %v", err)
	}
}