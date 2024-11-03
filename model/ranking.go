package model

import (
	"encoding/json"
	"log"
	"os"
)

type Ranking struct {
	CountryID      int     `json:"country_id"`
	LeagueID       int     `json:"league_id"`
	Rank           int     `json:"rank"`
	PreviousRank   int     `json:"previous_rank"`
	Points         float64 `json:"points"`
	PreviousPoints float64 `json:"previous_points"`
}

func RankingFromJSONFile(path string) []Ranking {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[ranking_from_json][read_file][err:%v]", err)
		return []Ranking{}
	}

	var data []Ranking
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("[ranking_from_json][json_unmarshal][err:%v]", err)
		return []Ranking{}
	}

	return data
}
