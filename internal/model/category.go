package model

import "strings"

type Category struct {
	Codename      string   `json:"codename"`
	Name          string   `json:"name"`
	IsBaseExpense bool     `json:"is_base_expense"`
	Aliases       []string `json:"aliases"`
}

type DBCategory struct {
	Codename      string
	Name          string
	IsBaseExpense bool
	Aliases       string
}

func (dc *DBCategory) ToCategory() *Category {
	return &Category{
		Codename:      dc.Codename,
		Name:          dc.Name,
		IsBaseExpense: dc.IsBaseExpense,
		Aliases:       strings.Split(dc.Aliases, ", "),
	}
}

func (c *Category) ToDB() *DBCategory {
	return &DBCategory{
		Codename:      c.Codename,
		Name:          c.Name,
		IsBaseExpense: c.IsBaseExpense,
		Aliases:       strings.Join(c.Aliases, ", "),
	}
}

func (dc *DBCategory) TableName() string {
	return "categories"
}
