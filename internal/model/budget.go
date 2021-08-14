package model

type Budget struct {
	Codename   string `json:"codename"`
	DailyLimit int    `json:"daily_limit"`
}
