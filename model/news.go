package model

import (
	"encoding/json"
	"log"
	"os"
	"sort"
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

func NewsFromJSONFile(path string) NewsList {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[news_from_json][read_file][err:%v]", err)
		return NewsList{}
	}

	var data NewsList
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Printf("[news_from_json][json_unmarshal][err:%v]", err)
		return NewsList{}
	}

	return data.OrderByDate()
}

func (n NewsList) OrderByDate() NewsList {
	sort.Sort(n)

	return n
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

func (n NewsList) Len() int {
	return len(n)
}

func (n NewsList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NewsList) Less(i, j int) bool {
	newsDateI, err := time.Parse("2006-01-02", n[i].Date)
	if err != nil {
		return false
	}

	newsDateJ, err := time.Parse("2006-01-02", n[j].Date)
	if err != nil {
		return false
	}

	return newsDateI.After(newsDateJ)
}
