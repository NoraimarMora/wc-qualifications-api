package main

import (
	"log"
	"net/http"
	"time"

	"ws-qualifications-api/inmem"
	"ws-qualifications-api/provider"
)

// @title WORLD CUP QUALIFICATION API
// @version 1.0
// @description This service is in charge of selecting which stores are the most suitable according to the user’s latitude and longitude. This microservice provides the necessary endpoints for store listing, store details, checkout, and other critical services. It is an essential component for the operation and user experience on our platform.
// @termsOfService http://swagger.io/terms/
// @BasePath /wc-qualification/api/v1
// @Schemes http https
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
