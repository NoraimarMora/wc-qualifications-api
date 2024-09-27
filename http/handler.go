package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws-qualifications-api/inmem"
)

type Handler struct {
	repository inmem.Repository
}

func NewHandler(repository inmem.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func (h *Handler) GetCountries(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"countries": "OK"})
}

func (h *Handler) GetCountryByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"country": "OK"})
}

func (h *Handler) GetLeagues(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"leagues": "OK"})
}

func (h *Handler) GetLeagueByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"league": "OK"})
}

func (h *Handler) GetMatches(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"matches": "OK"})
}

func (h *Handler) GetMatchsByLeagueID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"matches": "OK"})
}

func (h *Handler) GetMatchByID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"match": "OK"})
}

func (h *Handler) GetStandings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"standings": "OK"})
}

func (h *Handler) GetStandingsByLeagueID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"standings": "OK"})
}

func (h *Handler) GetStandingsByCountryID(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"standings": "OK"})
}
