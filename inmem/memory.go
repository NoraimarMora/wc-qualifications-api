package inmem

import "ws-qualifications-api/model"

type Memory struct {
	provider  Provider
	matches   map[int]model.Matches
	countries map[int]model.Country
	leagues   map[int]model.League
	standings map[int]model.Standings
}

func NewMemoryRepository(provider Provider) Repository {
	r := Memory{
		provider: provider,
	}

	r.GetCountries()
	r.GetLeagues()
	r.GetStandings()
	r.GetMatches()

	return &r
}

func (r *Memory) GetCountries() map[int]model.Country {
	if r.countries == nil {
		r.countries = r.provider.LoadCountries()
	}
	return r.countries
}

func (r *Memory) GetStandings() map[int]model.Standings {
	if r.standings == nil {
		r.standings = r.provider.LoadStandings()
	}
	return r.standings
}

func (r *Memory) GetMatches() map[int]model.Matches {
	if r.matches == nil {
		r.matches = r.provider.LoadMatches()
	}
	return r.matches
}

func (r *Memory) GetLeagues() map[int]model.League {
	if r.leagues == nil {
		r.leagues = r.provider.LoadLeagues()
	}
	return r.leagues
}
