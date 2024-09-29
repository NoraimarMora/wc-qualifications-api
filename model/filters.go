package model

import "time"

type Filters struct {
	Stage      string    `json:"stage"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	Status     string    `json:"status"`
	HometeamID string    `json:"hometeam_id"`
	AwayteamID string    `json:"awayteam_id"`
}
