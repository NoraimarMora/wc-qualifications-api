package model

type Countries map[int]Country

type Country struct {
	ID   int                        `json:"country_id"`
	Name map[string]NameTranslation `json:"name"`
	Flag map[string]string          `json:"flag"`
	Fifa string                     `json:"fifa"`
}

type NameTranslation struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}
