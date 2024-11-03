package inmem

import "ws-qualifications-api/model"

type Memory struct {
	provider       Provider
	matches        map[int]model.Matches
	countries      model.CountriesByID
	leagues        model.LeaguesByID
	standings      model.StandingsByLeague
	news           model.NewsList
	ranking        []model.Ranking
}

func NewMemoryRepository(provider Provider) Repository {
	r := Memory{
		provider: provider,
	}

	r.LoadCountries()
	r.LoadLeagues()
	r.LoadStandings()
	r.LoadMatches()
	r.LoadNews()
	r.LoadRanking()

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

func (r *Memory) LoadNews() model.NewsList {
	if r.news == nil {
		r.news = r.provider.LoadNews()
	}
	return r.news
}

func (r *Memory) LoadRanking() []model.Ranking {
	if r.ranking == nil {
		r.ranking = r.provider.LoadRanking()
	}
	return r.ranking
}

func (r *Memory) GetCountries() []model.Country {
	return r.countries.ToSlice()
}

func (r *Memory) GetCountryByID(countryID int) model.Country {
	if country, ok := r.countries[countryID]; ok {
		return country
	}

	return model.Country{}
}

func (r *Memory) GetLeagues() []model.League {
	return r.leagues.ToSlice()
}

func (r *Memory) GetLeagueByID(leagueID int) model.League {
	if league, ok := r.leagues[leagueID]; ok {
		return league
	}

	return model.League{}
}

func (r *Memory) GetMatches(filters model.Filters) []model.Match {
	matches := make([]model.Match, 0)

	for _, matchesByLeague := range r.matches {
		matches = append(
			matches,
			matchesByLeague.Copy().
				ByStage(filters.Stage).
				ByStatus(filters.Status).
				ByHomeTeamID(filters.HometeamID).
				ByAwayTeamID(filters.AwayteamID).
				ByFromDate(filters.From).
				ByToDate(filters.To).
				ToSlice()...,
		)
	}

	return matches
}

func (r *Memory) GetMatchesByLeagueID(leagueID int, filters model.Filters) []model.Match {
	if matchesByLeague, ok := r.matches[leagueID]; ok {
		return matchesByLeague.Copy().
			ByStage(filters.Stage).
			ByStatus(filters.Status).
			ByHomeTeamID(filters.HometeamID).
			ByAwayTeamID(filters.AwayteamID).
			ByFromDate(filters.From).
			ByToDate(filters.To).
			ToSlice()
	}

	return []model.Match{}
}

func (r *Memory) GetMatchByID(leagueID, matchID int) model.Match {
	if matchesByLeague, ok := r.matches[leagueID]; ok {
		if match, ok2 := matchesByLeague[matchID]; ok2 {
			return match
		}
	}

	return model.Match{}
}

func (r *Memory) GetStandings() []model.Standing {
	return r.standings.ToSlice()
}

func (r *Memory) GetStandingsByLeagueID(leagueID int, filters model.Filters) []model.Standing {
	if standingsByLeague, ok := r.standings[leagueID]; ok {
		return standingsByLeague.ByStage(filters.Stage).ToSlice()
	}

	return []model.Standing{}
}

func (r *Memory) GetStandingsByCountryID(leagueID, countryID int, filters model.Filters) []model.Standing {
	if standingsByLeague, ok := r.standings[leagueID]; ok {
		return standingsByLeague.ByStage(filters.Stage).ByCountry(countryID)
	}

	return []model.Standing{}
}

func (r *Memory) GetNews(filters model.Filters) model.NewsList {
	return r.news.ByFromDate(filters.From).ByToDate(filters.To)
}

func (r *Memory) GetRanking() []model.Ranking {
	return r.ranking
}
