package model

import (
	"encoding/json"
	"log"
	"os"
)

type League struct {
	ID                   int                 `json:"id"`
	Name                 map[string]string   `json:"name"`
	Alias                string              `json:"alias"`
	Logo                 string              `json:"logo"`
	Year                 int                 `json:"year"`
	QualificationProcess map[string][]string `json:"qualification_process"`
}

func LeaguesFromJSONFile(path string) map[int]League {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[leagues_from_json][read_file][err: %v]", err)
		return make(map[int]League, 0)
	}

	var data []League
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("[leagues_from_json][json_unmarshal][err: %v]", err)
		return make(map[int]League, 0)
	}

	return leaguesSliceToMap(data)
}

func leaguesSliceToMap(data []League) map[int]League {
	leagues := make(map[int]League, 0)

	for _, league := range data {
		leagues[league.ID] = league
	}

	return leagues
}
