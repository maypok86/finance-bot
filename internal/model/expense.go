package model

import "gorm.io/datatypes"

type Expense struct {
	ID               int            `json:"id"`
	Amount           int            `json:"amount"`
	CreatedDate      datatypes.Date `json:"created_date"`
	CategoryCodename string         `json:"category_codename"`
	RawText          string         `json:"raw_text"`
}

func (e *Expense) TableName() string {
	return "expenses"
}
