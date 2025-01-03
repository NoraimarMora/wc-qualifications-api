package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	wc_http "ws-qualifications-api/http"
	"ws-qualifications-api/inmem"
)

// setupRoutes returns a Gin server ready to rise up with all the available endpoints.
func setupRoutes(repository inmem.Repository) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.SetHTMLTemplate(template.Must(template.ParseFiles("./front/html/index.html")))

	prefix := router.Group("api/v1")

	handler := wc_http.NewHandler(repository)

	prefix.StaticFile("/styles.css", "./front/assets/styles.css")
	prefix.StaticFile("/functions.js", "./front/assets/functions.js")

	prefix.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "World Cup Qualifiers API",
		})
	})

	// Health check of the app.
	prefix.GET("/health", corsMiddleware(), handler.HealthCheck)

	prefix.GET("/countries", corsMiddleware(), handler.GetCountries)
	prefix.GET("/countries/:country_id", corsMiddleware(), handler.GetCountryByID)

	prefix.GET("/leagues", corsMiddleware(), handler.GetLeagues)
	prefix.GET("/leagues/:league_id", corsMiddleware(), handler.GetLeagueByID)

	prefix.GET("/matches", corsMiddleware(), handler.GetMatches)
	prefix.GET("/matches/:league_id", corsMiddleware(), handler.GetMatchesByLeagueID)
	prefix.GET("/matches/:league_id/:match_id", corsMiddleware(), handler.GetMatchByID)

	prefix.GET("/standings", corsMiddleware(), handler.GetStandings)
	prefix.GET("/standings/:league_id", corsMiddleware(), handler.GetStandingsByLeagueID)
	prefix.GET("/standings/:league_id/:country_id", corsMiddleware(), handler.GetStandingsByCountryID)

	prefix.GET("/news", corsMiddleware(), handler.GetNews)

	prefix.GET("/ranking", corsMiddleware(), handler.GetRanking)

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, PATCH, OPTIONS")
		c.Writer.Header().Set("Content-Type", "application/json")

		c.Next()
	}
}
