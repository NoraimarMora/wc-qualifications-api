package main

import (
	"log"
	"net/http"
	"time"

	"ws-qualifications-api/inmem"
	"ws-qualifications-api/provider"
)

func main() {
	localProvider := provider.Local{Path: "./files"}
	repository := inmem.NewMemoryRepository(localProvider)

	router := setupRoutes(repository)

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
