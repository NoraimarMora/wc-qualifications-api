package inmem

import "ws-qualifications-api/model"

type Repository interface {
	GetCountries() map[int]model.Country
	GetStandings() map[int]model.Standings
	GetMatches() map[int]model.Matches
	GetLeagues() map[int]model.League
}

type Provider interface {
	LoadLeagues() map[int]model.League
	LoadCountries() map[int]model.Country
	LoadMatches() map[int]model.Matches
	LoadStandings() map[int]model.Standings
}
