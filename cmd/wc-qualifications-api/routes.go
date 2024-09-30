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

	prefix.GET("/countries", handler.GetCountries)
	prefix.GET("/countries/:country_id", handler.GetCountryByID)

	prefix.GET("/leagues", handler.GetLeagues)
	prefix.GET("/leagues/:league_id", handler.GetLeagueByID)

	prefix.GET("/matches", handler.GetMatches)
	prefix.GET("/matches/:league_id", handler.GetMatchesByLeagueID)
	prefix.GET("/matches/:league_id/:match_id", handler.GetMatchByID)

	prefix.GET("/standings", handler.GetStandings)
	prefix.GET("/standings/:league_id", handler.GetStandingsByLeagueID)
	prefix.GET("/standings/:league_id/:country_id", handler.GetStandingsByCountryID)

	return router
}
