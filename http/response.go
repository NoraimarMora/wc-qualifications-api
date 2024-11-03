package http

import "ws-qualifications-api/model"

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"error"`
}

type CountriesResponse struct {
	Countries []model.Country `json:"countries"`
}

type CountryResponse struct {
	Country model.Country `json:"country"`
}

type LeaguesResponse struct {
	Leagues []model.League `json:"leagues"`
}

type LeagueResponse struct {
	League model.League `json:"league"`
}

type MatchesResponse struct {
	Matches []model.Match `json:"matches"`
}

type MatchResponse struct {
	Match model.Match `json:"match"`
}

type StandingsResponse struct {
	Standings []model.Standing `json:"standings"`
}

type NewsResponse struct {
	News model.NewsList `json:"news"`
}

type RankingResponse struct {
	Ranking []model.Ranking `json:"ranking"`
}
