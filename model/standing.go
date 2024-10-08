package model

import (
	"encoding/json"
	"log"
	"os"
)

type StandingsByLeague map[int]Standings

// Standings by stage
type Standings map[string][]Standing

type Standing struct {
	CountryID      int    `json:"country_id"`
	LeagueID       int    `json:"league_id"`
	ActualPosition int    `json:"actual_pos"`
	PrevPosition   int    `json:"prev_pos"`
	MatchesPlayed  int    `json:"matches_played"`
	Wins           int    `json:"wins"`
	Draws          int    `json:"draws"`
	Loss           int    `json:"loss"`
	GoalsScored    int    `json:"goals_scored"`
	GoalsAgainst   int    `json:"goals_against"`
	GoalDifference int    `json:"goal_difference"`
	Points         int    `json:"points"`
	Group          string `json:"group"`
	Stage          string `json:"stage"`
}

func StandingsFromJSONFile(paths []string) map[int]Standings {
	standings := map[int]Standings{}

	for _, path := range paths {
		file, err := os.ReadFile(path)
		if err != nil {
			log.Printf("[standings_from_json][read_file][err:%v]", err)
			continue
		}

		var data []Standing
		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Printf("[standings_from_json][json_unmarshal][err: %v]", err)
			continue
		}

		standings = standingsSliceToMap(data, standings)
	}

	return standings
}

func standingsSliceToMap(data []Standing, standings map[int]Standings) map[int]Standings {
	for _, standing := range data {
		if s, ok := standings[standing.LeagueID]; !ok {
			standings[standing.LeagueID] = map[string][]Standing{
				standing.Stage: make([]Standing, 0),
			}
		} else {
			if _, ok2 := s[standing.Stage]; !ok2 {
				standings[standing.LeagueID][standing.Stage] = make([]Standing, 0)
			}
		}

		standings[standing.LeagueID][standing.Stage] = append(standings[standing.LeagueID][standing.Stage], standing)
	}

	return standings
}

func (s StandingsByLeague) ToSlice() []Standing {
	standings := make([]Standing, 0)

	for _, standingsByLeague := range s {
		for _, standingsByStage := range standingsByLeague {
			standings = append(standings, standingsByStage...)
		}
	}

	return standings
}

func (s Standings) ToSlice() []Standing {
	standings := make([]Standing, 0)

	for _, standingsByStage := range s {
		standings = append(standings, standingsByStage...)
	}

	return standings
}

func (s Standings) ByStage(stage string) Standings {
	if stage == "" {
		return s
	}

	if standingsByStage, ok := s[stage]; ok {
		return Standings{
			stage: standingsByStage,
		}
	}

	return Standings{}
}

func (s Standings) ByCountry(countryID int) []Standing {
	standings := make([]Standing, 0)

	for _, standingsByStage := range s {
		for _, standing := range standingsByStage {
			if standing.CountryID == countryID {
				standings = append(standings, standing)
				break
			}
		}
	}

	return standings
}
