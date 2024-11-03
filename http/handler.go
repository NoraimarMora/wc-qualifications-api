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

// GetCountries
// @Summary Returns the list of countries participating in the qualifiers
// @Tags Country
// @Produce json
// @Success 200 {object} CountriesResponse "Retrieves the list of participating countries"
// @Router /countries  [get]
func (h *Handler) GetCountries(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, CountriesResponse{h.repository.GetCountries()})
}

// GetCountryByID
// @Summary Returns the detail of a country according to a given country_id
// @Tags Country
// @Produce json
// @Param country_id path int true "Country ID"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Country not found"
// @Success 200 {object} CountryResponse "Retrieves the detail of a country"
// @Router /countries/{country_id}  [get]
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

// GetLeagues
// @Summary Returns the list of leagues/confederations participating in the qualifiers
// @Tags League
// @Produce json
// @Success 200 {object} LeaguesResponse "Retrieves the list of participating leagues/confederations"
// @Router /leagues  [get]
func (h *Handler) GetLeagues(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, LeaguesResponse{h.repository.GetLeagues()})
}

// GetLeagueByID
// @Summary Returns the detail of a league according to a given league_id
// @Tags League
// @Produce json
// @Param league_id path int true "League ID"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "League not found"
// @Success 200 {object} LeagueResponse "Retrieves the detail of a league"
// @Router /leagues/{league_id}  [get]
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

// GetMatches
// @Summary Returns the list of matches of the qualifiers
// @Tags Match
// @Produce json
// @Param stage query int false "Stage"
// @Param status query int false "Status"
// @Param hometeam_id query int false "Hometeam ID"
// @Param awayteam_id query int false "Awayteam ID"
// @Param from query string false "From date" "2023-10-02"
// @Param to query string false "To date" "2023-10-02"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Success 200 {object} MatchesResponse "Retrieves the list of matches"
// @Router /matches  [get]
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

// GetMatchesByLeagueID
// @Summary Returns the list of matches of the qualifiers according to a given league_id
// @Tags Match
// @Produce json
// @Param league_id path int true "League ID"
// @Param stage query int false "Stage" [0, 1, 2, 3, 4]
// @Param status query int false "Status" [0, 1, 2]
// @Param hometeam_id query int false "Hometeam ID"
// @Param awayteam_id query int false "Awayteam ID"
// @Param from query string false "From date" "2023-10-02"
// @Param to query int false "To date" "2023-10-02"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Matches not found"
// @Success 200 {object} MatchesResponse "Retrieves the list of matches of a league"
// @Router /matches/{league_id}  [get]
func (h *Handler) GetMatchesByLeagueID(ctx *gin.Context) {
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

// GetMatchByID
// @Summary Returns the detail of a match ¿according to a given league_id and match_id
// @Tags Match
// @Produce json
// @Param league_id path int true "League ID"
// @Param match_id path int true "Match ID"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Match not found"
// @Success 200 {object} MatchResponse "Retrieves the detail of a match"
// @Router /matches/{league_id}/{match_id}  [get]
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

// GetStandings
// @Summary Returns the list of standings
// @Tags Standing
// @Produce json
// @Success 200 {object} StandingsResponse "Retrieves the list of standings"
// @Router /standings  [get]
func (h *Handler) GetStandings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, StandingsResponse{h.repository.GetStandings()})
}

// GetStandingsByLeagueID
// @Summary Returns the list of standings according to a given league_id
// @Tags Standing
// @Produce json
// @Param league_id path int true "League ID"
// @Param stage query int false "Stage" [0, 1, 2, 3, 4]
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Standings not found"
// @Success 200 {object} StandingsResponse "Retrieves the list of standings of a league"
// @Router /standings/{league_id}  [get]
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

// GetStandingsByCountryID
// @Summary Returns the list of standings according to a given league_id and country_id
// @Tags Standing
// @Produce json
// @Param league_id path int true "League ID"
// @Param country_id path int true "Country ID"
// @Param stage query int false "Stage" [0, 1, 2, 3, 4]
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 404 {object} ErrorResponse "Standings not found"
// @Success 200 {object} StandingsResponse "Retrieves the list of standings of a country"
// @Router /standings/{league_id}/{country_id}  [get]
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

// GetNews
// @Summary Returns the list of matches of the qualifiers
// @Tags News
// @Produce json
// @Param from query string false "From date" "2023-10-02"
// @Param to query string false "To date" "2023-10-02"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Success 200 {object} NewsResponse "Retrieves the list of news"
// @Router /news  [get]
func (h *Handler) GetNews(ctx *gin.Context) {
	filters, err := validateQueryParams(ctx)
	if err != nil {
		msg := fmt.Sprintf("[get_news]%v", err)
		log.Println(msg)

		ctx.JSON(http.StatusBadRequest, ErrorResponse{msg})
		return
	}

	ctx.JSON(http.StatusOK, NewsResponse{h.repository.GetNews(filters)})
}

// GetRanking
// @Summary Returns the last male world classification
// @Tags Ranking
// @Produce json
// @Failure 400 {object} ErrorResponse "Bad request"
// @Success 200 {object} RankingResponse "Retrieves the male world classification"
// @Router /ranking  [get]
func (h *Handler) GetRanking(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, RankingResponse{h.repository.GetRanking()})
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
