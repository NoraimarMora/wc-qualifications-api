package model

type Leagues map[int]League

type League struct {
	LeagueID             int                 `json:"league_id"`
	Name                 map[string]string   `json:"name"`
	Alias                string              `json:"alias"`
	Logo                 string              `json:"logo"`
	Year                 int                 `json:"year"`
	QualificationProcess map[string][]string `json:"qualification_process"`
}
