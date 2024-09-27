package model

import (
	"encoding/json"
	"log"
	"os"
)

type Matches map[int]Match

type Match struct {
	ID                   int                      `json:"id"`
	LeagueID             int                      `json:"league_id"`
	Date                 string                   `json:"date"`
	Time                 string                   `json:"time"`
	Status               string                   `json:"status"`
	Round                string                   `json:"round"`
	Referee              string                   `json:"referee"`
	Stadium              string                   `json:"stadium"`
	Stage                string                   `json:"stage"`
	HomeTeamID           string                   `json:"hometeam_id"`
	AwayTeamID           string                   `json:"awayteam_id"`
	HomeTeamScore        string                   `json:"hometeam_score"`
	AwayTeamScore        string                   `json:"awayteam_score"`
	HomeTeamPenaltyScore string                   `json:"hometeam_penalty_score"`
	AwayTeamPenaltyScore string                   `json:"awayteam_penalty_score"`
	LineUps              map[string]LineUp        `json:"lineups"`
	GoalScorers          []Goal                   `json:"goal_scorers"`
	Cards                []Card                   `json:"cards"`
	Substitutions        map[string]Substitutions `json:"substitutions"`
	Statistics           []Stat                   `json:"statistics"`
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

type Substitutions []Substitution

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

func MatchesFromJSONFile(paths []string) map[int]Matches {
	matches := map[int]Matches{}

	for _, path := range paths {
		file, err := os.ReadFile(path)
		if err != nil {
			log.Printf("[matches_from_json][read_file][err: %v]", err)
			continue
		}

		var data []Match
		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Printf("[matches_from_json][json_unmarshal][err: %v]", err)
			continue
		}

		matches = matchesSliceToMap(data, matches)
	}

	return matches
}

func matchesSliceToMap(data []Match, matches map[int]Matches) map[int]Matches {
	for _, match := range data {
		if _, ok := matches[match.LeagueID]; !ok {
			matches[match.LeagueID] = Matches{}
		}

		matches[match.LeagueID][match.ID] = match
	}

	return matches
}
