package model

import (
	"time"
)

type Expense struct {
	ID               int       `json:"id"`
	Amount           int       `json:"amount"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	CategoryCodename string    `json:"category_codename"`
	RawText          string    `json:"raw_text"`
}

func (e *Expense) TableName() string {
	return "expenses"
}
