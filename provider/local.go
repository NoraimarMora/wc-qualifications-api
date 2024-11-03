package provider

import "ws-qualifications-api/model"

type Local struct {
	Path string
}

func (l Local) LoadLeagues() map[int]model.League {
	return model.LeaguesFromJSONFile(l.Path + "/leagues.json")
}

func (l Local) LoadCountries() map[int]model.Country {
	return model.CountriesFromJSONFile(l.Path + "/countries.json")
}

func (l Local) LoadMatches() map[int]model.Matches {
	paths := []string{
		l.Path + "/afc.json",
		l.Path + "/caf.json",
		l.Path + "/concacaf.json",
		l.Path + "/conmebol.json",
		l.Path + "/ofc.json",
	}

	return model.MatchesFromJSONFile(paths)
}

func (l Local) LoadStandings() map[int]model.Standings {
	paths := []string{
		l.Path + "/standings_afc.json",
		l.Path + "/standings_caf.json",
		l.Path + "/standings_concacaf.json",
		l.Path + "/standings_conmebol.json",
		l.Path + "/standings_ofc.json",
	}

	return model.StandingsFromJSONFile(paths)
}

func (l Local) LoadNews() model.NewsList {
	return model.NewsFromJSONFile(l.Path + "/news.json")
}

func (l Local) LoadRanking() []model.Ranking {
	return model.RankingFromJSONFile(l.Path + "/ranking.json")
}
