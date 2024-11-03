package inmem

import "ws-qualifications-api/model"

type Repository interface {
	LoadCountries() map[int]model.Country
	LoadStandings() map[int]model.Standings
	LoadMatches() map[int]model.Matches
	LoadLeagues() map[int]model.League
	LoadNews() model.NewsList
	LoadRanking() []model.Ranking

	GetCountries() []model.Country
	GetCountryByID(countryID int) model.Country
	GetLeagues() []model.League
	GetLeagueByID(leagueID int) model.League
	GetMatches(filters model.Filters) []model.Match
	GetMatchesByLeagueID(leagueID int, filters model.Filters) []model.Match
	GetMatchByID(leagueID, matchID int) model.Match
	GetStandings() []model.Standing
	GetStandingsByLeagueID(leagueID int, filters model.Filters) []model.Standing
	GetStandingsByCountryID(leagueID, countryID int, filters model.Filters) []model.Standing
	GetNews(filters model.Filters) model.NewsList
	GetRanking() []model.Ranking
}

type Provider interface {
	LoadLeagues() map[int]model.League
	LoadCountries() map[int]model.Country
	LoadMatches() map[int]model.Matches
	LoadStandings() map[int]model.Standings
	LoadNews() model.NewsList
	LoadRanking() []model.Ranking
}
