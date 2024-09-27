package model

type Matches map[int]Match

type Match struct {
	MatchID              int               `json:"match_id"`
	LeagueID             int               `json:"league_id"`
	Date                 string            `json:"date"`
	Time                 string            `json:"time"`
	Status               string            `json:"status"`
	Round                string            `json:"round"`
	Referee              string            `json:"referee"`
	Stadium              string            `json:"stadium"`
	Stage                string            `json:"stage"`
	HomeTeamID           int               `json:"home_team_id"`
	AwayTeamID           int               `json:"away_team_id"`
	HomeTeamScore        int               `json:"home_team_score"`
	AwayTeamScore        int               `json:"away_team_score"`
	HomeTeamPenaltyScore int               `json:"home_team_penalty_score"`
	AwayTeamPenaltyScore int               `json:"away_team_penalty_score"`
	LineUps              map[string]LineUp `json:"lineups"`
	GoalScorers          []Goal            `json:"goal_scorers"`
	Cards                []Card            `json:"cards"`
	Substitutions        []Substitution    `json:"substitutions"`
	Statistics           []Stat            `json:"statistics"`
}

type LineUp struct {
	System          string   `json:"system"`
	Coach           string   `json:"coach"`
	StartingPlayers []Player `json:"starting_players"`
	Substitutes     []Player `json:"substitutes"`
}

type Player struct {
	Name     string `json:"player_name"`
	Number   int    `json:"player_number"`
	Position int    `json:"player_position"`
}

type Goal struct {
	Time       string `json:"time"`
	PlayerName string `json:"player"`
	Team       string `json:"team"`
}

type Card struct {
	Time       string `json:"time"`
	CardType   string `json:"card"`
	PlayerName string `json:"player"`
	Team       string `json:"team"`
}

type Substitution struct {
	Time string `json:"time"`
	In   string `json:"in"`
	Out  string `json:"out"`
}

type Stat struct {
	Type map[string]string `json:"type"`
	Home string            `json:"home"`
	Away string            `json:"away"`
}
