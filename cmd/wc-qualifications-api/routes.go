package main

import (
	"github.com/gin-gonic/gin"
)

// setupRoutes returns a Gin server ready to rise up with all the available endpoints.
func setupRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	return router
}
