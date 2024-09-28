package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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
	ctx.JSON(http.StatusOK, gin.H{"countries": h.repository.GetCountries()})
}

func (h *Handler) GetCountryByID(ctx *gin.Context) {
	countryID, err := strconv.Atoi(ctx.Param("country_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_country_by_id][parse_country_id:%s][err:%v]", ctx.Param("country_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	country := h.repository.GetCountryByID(countryID)
	if country.ID == 0 {
		msg := fmt.Sprintf("[get_country_by_id][get_country_by_id][err:country not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"country": country})
}

func (h *Handler) GetLeagues(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"leagues": h.repository.GetLeagues()})
}

func (h *Handler) GetLeagueByID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_league_by_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	league := h.repository.GetLeagueByID(leagueID)
	if league.ID == 0 {
		msg := fmt.Sprintf("[get_league_by_id][get_league_by_id][err:league not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"league": league})
}

func (h *Handler) GetMatches(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"matches": h.repository.GetMatches()})
}

func (h *Handler) GetMatchsByLeagueID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_matches_by_league_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	matches := h.repository.GetMatchesByLeagueID(leagueID)
	if len(matches) == 0 {
		msg := fmt.Sprintf("[get_matches_by_league_id][get_matches_by_league_id][err:matches not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"matches": matches})
}

func (h *Handler) GetMatchByID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_match_by_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	matchID, err := strconv.Atoi(ctx.Param("match_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_match_by_id][parse_match_id:%s][err:%v]", ctx.Param("match_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	match := h.repository.GetMatchByID(leagueID, matchID)
	if match.ID == 0 {
		msg := fmt.Sprintf("[get_match_by_id][get_match_by_id][err:match not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"match": match})
}

func (h *Handler) GetStandings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"standings": h.repository.GetStandings()})
}

func (h *Handler) GetStandingsByLeagueID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	standings := h.repository.GetStandingsByLeagueID(leagueID)
	if len(standings) == 0 {
		msg := fmt.Sprintf("[get_standings_by_league_id][get_standings_by_league_id][err:standings not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"standings": standings})
}

func (h *Handler) GetStandingsByCountryID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	countryID, err := strconv.Atoi(ctx.Param("country_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id][parse_country_id:%s][err:%v]", ctx.Param("country_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	standings := h.repository.GetStandingsByCountryID(leagueID, countryID)
	if len(standings) == 0 {
		msg := fmt.Sprintf("[get_standings_by_league_id][get_standings_by_league_id][err:standings not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"standings": standings})
}
