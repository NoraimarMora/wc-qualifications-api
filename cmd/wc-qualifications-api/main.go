package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := setupRoutes()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
