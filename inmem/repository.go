package inmem

import "ws-qualifications-api/model"

type Repository interface {
	LoadCountries() map[int]model.Country
	LoadStandings() map[int]model.Standings
	LoadMatches() map[int]model.Matches
	LoadLeagues() map[int]model.League

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
}

type Provider interface {
	LoadLeagues() map[int]model.League
	LoadCountries() map[int]model.Country
	LoadMatches() map[int]model.Matches
	LoadStandings() map[int]model.Standings
}
