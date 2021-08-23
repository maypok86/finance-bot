package model

// Budget model.
type Budget struct {
	Codename   string `json:"codename"`
	DailyLimit int    `json:"daily_limit"`
}

// TableName returns a Budget table name.
func (b *Budget) TableName() string {
	return "budgets"
}
