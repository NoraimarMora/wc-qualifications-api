package model

type Standings []Standing

type Standing struct {
	CountryID      int    `json:"country_id"`
	LeagueID       int    `json:"league_id"`
	ActualPosition int    `json:"actual_position"`
	PrevPosition   int    `json:"prev_position"`
	MatchesPlayed  int    `json:"matches_played"`
	Wins           int    `json:"wins"`
	Draws          int    `json:"draws"`
	Loss           int    `json:"loss"`
	GoalsScored    int    `json:"goals_scored"`
	GoalsAgainst   int    `json:"goals_against"`
	GoalDifference int    `json:"goal_difference"`
	Group          string `json:"group"`
	Stage          string `json:"stage"`
}
