package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ws-qualifications-api/inmem"
	"ws-qualifications-api/model"
)

var stages = []string{
	"Current",
	"1st Round",
	"2nd Round",
	"3rd Round",
	"Group Stage",
}

var status = []string{
	"Finished",
	"Cancelled",
	"Postponed",
}

type Handler struct {
	repository inmem.Repository
}

func NewHandler(repository inmem.Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthCheckResponse{"OK"})
}

func (h *Handler) GetCountries(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, CountriesResponse{h.repository.GetCountries()})
}

func (h *Handler) GetCountryByID(ctx *gin.Context) {
	countryID, err := strconv.Atoi(ctx.Param("country_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_country_by_id][parse_country_id:%s][err:%v]", ctx.Param("country_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	country := h.repository.GetCountryByID(countryID)
	if country.ID == 0 {
		msg := fmt.Sprintf("[get_country_by_id][get_country_by_id][err:country not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, CountryResponse{country})
}

func (h *Handler) GetLeagues(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, LeaguesResponse{h.repository.GetLeagues()})
}

func (h *Handler) GetLeagueByID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_league_by_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	league := h.repository.GetLeagueByID(leagueID)
	if league.ID == 0 {
		msg := fmt.Sprintf("[get_league_by_id][get_league_by_id][err:league not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, LeagueResponse{league})
}

func (h *Handler) GetMatches(ctx *gin.Context) {
	filters, err := validateQueryParams(ctx)
	if err != nil {
		msg := fmt.Sprintf("[get_matches]%v", err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, MatchesResponse{h.repository.GetMatches(filters)})
}

func (h *Handler) GetMatchsByLeagueID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_matches_by_league_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	filters, err := validateQueryParams(ctx)
	if err != nil {
		msg := fmt.Sprintf("[get_matches_by_league_id]%v", err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	matches := h.repository.GetMatchesByLeagueID(leagueID, filters)
	if len(matches) == 0 {
		msg := fmt.Sprintf("[get_matches_by_league_id][get_matches_by_league_id][err:matches not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, MatchesResponse{matches})
}

func (h *Handler) GetMatchByID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_match_by_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	matchID, err := strconv.Atoi(ctx.Param("match_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_match_by_id][parse_match_id:%s][err:%v]", ctx.Param("match_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	match := h.repository.GetMatchByID(leagueID, matchID)
	if match.ID == 0 {
		msg := fmt.Sprintf("[get_match_by_id][get_match_by_id][err:match not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, MatchResponse{match})
}

func (h *Handler) GetStandings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StandingsResponse{h.repository.GetStandings()})
}

func (h *Handler) GetStandingsByLeagueID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	filters, err := validateQueryParams(ctx)
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id]%v", err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	standings := h.repository.GetStandingsByLeagueID(leagueID, filters)
	if len(standings) == 0 {
		msg := fmt.Sprintf("[get_standings_by_league_id][get_standings_by_league_id][err:standings not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, StandingsResponse{standings})
}

func (h *Handler) GetStandingsByCountryID(ctx *gin.Context) {
	leagueID, err := strconv.Atoi(ctx.Param("league_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id_and_country_id][parse_league_id:%s][err:%v]", ctx.Param("league_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	countryID, err := strconv.Atoi(ctx.Param("country_id"))
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id_and_country_id][parse_country_id:%s][err:%v]", ctx.Param("country_id"), err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	filters, err := validateQueryParams(ctx)
	if err != nil {
		msg := fmt.Sprintf("[get_standings_by_league_id_and_country_id]%v", err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	standings := h.repository.GetStandingsByCountryID(leagueID, countryID, filters)
	if len(standings) == 0 {
		msg := fmt.Sprintf("[get_standings_by_league_id_and_country_id][get_standings_by_league_id][err:standings not found]")
		log.Println(msg)

		ctx.JSON(http.StatusNotFound, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, StandingsResponse{standings})
}

func validateQueryParams(ctx *gin.Context) (model.Filters, error) {
	filters := model.Filters{}

	if ctx.Query("stage") != "" {
		stage, err := strconv.Atoi(ctx.Query("stage"))
		if err != nil {
			return filters, fmt.Errorf("[parse_stage:%s][err:%v]", ctx.Query("stage"), err)
		}

		if !(stage >= 0 && stage <= 4) {
			return filters, fmt.Errorf("[parse_stage:%s][err:index out of range]", ctx.Query("stage"))
		}

		filters.Stage = stages[stage]
	}

	if ctx.Query("status") != "" {
		st, err := strconv.Atoi(ctx.Query("status"))
		if err != nil {
			return filters, fmt.Errorf("[parse_status:%s][err:%v]", ctx.Query("status"), err)
		}

		if !(st >= 0 && st <= 2) {
			return filters, fmt.Errorf("[parse_status:%s][err:index out of range]", ctx.Query("status"))
		}

		filters.Status = status[st]
	}

	if ctx.Query("from") != "" {
		from, err := time.Parse("2006-01-02", ctx.Query("from"))
		if err != nil {
			return filters, fmt.Errorf("[parse_from:%s][err:%v]", ctx.Query("from"), err)
		}

		filters.From = from
	}

	if ctx.Query("to") != "" {
		to, err := time.Parse("2006-01-02", ctx.Query("to"))
		if err != nil {
			return filters, fmt.Errorf("[parse_to:%s][err:%v]", ctx.Query("to"), err)
		}

		filters.To = to
	}

	filters.HometeamID = ctx.Query("hometeam_id")
	filters.AwayteamID = ctx.Query("awayteam_id")

	return filters, nil
}
