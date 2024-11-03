package model

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type NewsList []News

type News struct {
	URL         map[string]string `json:"url"`
	Image       string            `json:"image"`
	Title       map[string]string `json:"title"`
	PreviewText map[string]string `json:"preview_text"`
	Date        string            `json:"date"`
}

func NewsFromJSONFile(path string) []News {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[news_from_json][read_file][err:%v]", err)
		return []News{}
	}

	var data []News
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("[news_from_json][json_unmarshal][err:%v]", err)
		return []News{}
	}

	return data
}

func (n NewsList) ByFromDate(from time.Time) NewsList {
	if from.IsZero() {
		return n
	}

	result := NewsList{}
	for _, news := range n {
		newsDate, err := time.Parse("2006-01-02", news.Date)
		if err != nil {
			continue
		}

		if from.Equal(newsDate) || from.Before(newsDate) {
			result = append(result, news)
		}
	}

	return result
}

func (n NewsList) ByToDate(to time.Time) NewsList {
	if to.IsZero() {
		return n
	}

	result := NewsList{}
	for _, news := range n {
		newsDate, err := time.Parse("2006-01-02", news.Date)
		if err != nil {
			continue
		}

		if to.Equal(newsDate) || to.After(newsDate) {
			result = append(result, news)
		}
	}

	return result
}
