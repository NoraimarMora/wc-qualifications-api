package model

import (
	"encoding/json"
	"log"
	"os"
)

type CountriesByID map[int]Country

type Country struct {
	ID   int                        `json:"id"`
	Name map[string]NameTranslation `json:"name"`
	Flag map[string]string          `json:"flags"`
	Fifa string                     `json:"fifa"`
}

type NameTranslation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

func CountriesFromJSONFile(path string) map[int]Country {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[countries_from_json][read_file][err:%v]", err)
		return make(map[int]Country, 0)
	}

	var data []Country
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("[countries_from_json][json_unmarshal][err:%v]", err)
		return make(map[int]Country, 0)
	}

	return countriesSliceToMap(data)
}

func countriesSliceToMap(data []Country) map[int]Country {
	countries := make(map[int]Country, 0)

	for _, country := range data {
		countries[country.ID] = country
	}

	return countries
}

func (c CountriesByID) ToSlice() []Country {
	countries := make([]Country, 0)

	for _, country := range c {
		countries = append(countries, country)
	}

	return countries
}
