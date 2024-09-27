package main

import (
	"github.com/gin-gonic/gin"

	wc_http "ws-qualifications-api/http"
	"ws-qualifications-api/inmem"
)

// setupRoutes returns a Gin server ready to rise up with all the available endpoints.
func setupRoutes(repository inmem.Repository) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	prefix := router.Group("wc-qualifications/api/v1")

	handler := wc_http.NewHandler(repository)

	// Health check of the app.
	prefix.GET("/health", handler.HealthCheck)

	return router
}
