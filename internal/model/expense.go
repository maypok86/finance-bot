package model

import (
	"time"
)

// Expense model
type Expense struct {
	ID               int       `json:"id"`
	Amount           int       `json:"amount"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	CategoryCodename string    `json:"category_codename"`
	RawText          string    `json:"raw_text"`
}

// TableName returns a Expense table name
func (e *Expense) TableName() string {
	return "expenses"
}
