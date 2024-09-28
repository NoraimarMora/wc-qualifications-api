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

	r.LoadCountries()
	r.LoadLeagues()
	r.LoadStandings()
	r.LoadMatches()

	return &r
}

func (r *Memory) LoadCountries() map[int]model.Country {
	if r.countries == nil {
		r.countries = r.provider.LoadCountries()
	}
	return r.countries
}

func (r *Memory) LoadStandings() map[int]model.Standings {
	if r.standings == nil {
		r.standings = r.provider.LoadStandings()
	}
	return r.standings
}

func (r *Memory) LoadMatches() map[int]model.Matches {
	if r.matches == nil {
		r.matches = r.provider.LoadMatches()
	}
	return r.matches
}

func (r *Memory) LoadLeagues() map[int]model.League {
	if r.leagues == nil {
		r.leagues = r.provider.LoadLeagues()
	}
	return r.leagues
}

func (r *Memory) GetCountryByID(countryID int) model.Country {
	if country, ok := r.countries[countryID]; ok {
		return country
	}

	return model.Country{}
}

func (r *Memory) GetCountries() []model.Country {
	countries := make([]model.Country, 0)

	return countries
}

func (r *Memory) GetLeagues() []model.League {
	leagues := make([]model.League, 0)

	return leagues
}

func (r *Memory) GetLeagueByID(leagueID int) model.League {
	return model.League{}
}

func (r *Memory) GetMatches() []model.Match {
	matches := make([]model.Match, 0)

	return matches
}

func (r *Memory) GetMatchesByLeagueID(leagueID int) []model.Match {
	matches := make([]model.Match, 0)

	return matches
}

func (r *Memory) GetMatchByID(leagueID, matchID int) model.Match {
	return model.Match{}
}

func (r *Memory) GetStandings() []model.Standing {
	standings := make([]model.Standing, 0)

	return standings
}

func (r *Memory) GetStandingsByLeagueID(leagueID int) []model.Standing {
	standings := make([]model.Standing, 0)

	return standings
}

func (r *Memory) GetStandingsByCountryID(leagueID, countryID int) []model.Standing {
	standings := make([]model.Standing, 0)

	return standings
}
