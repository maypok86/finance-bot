package model

import "strings"

// Category model.
type Category struct {
	Codename      string   `json:"codename"`
	Name          string   `json:"name"`
	IsBaseExpense bool     `json:"is_base_expense"`
	Aliases       []string `json:"aliases"`
}

// DBCategory model.
type DBCategory struct {
	Codename      string
	Name          string
	IsBaseExpense bool
	Aliases       string
}

// ToCategory converts a Category instance to DBCategory instance.
func (dc *DBCategory) ToCategory() *Category {
	return &Category{
		Codename:      dc.Codename,
		Name:          dc.Name,
		IsBaseExpense: dc.IsBaseExpense,
		Aliases:       strings.Split(dc.Aliases, ", "),
	}
}

// ToDB converts a DBCategory instance to Category instance.
func (c *Category) ToDB() *DBCategory {
	return &DBCategory{
		Codename:      c.Codename,
		Name:          c.Name,
		IsBaseExpense: c.IsBaseExpense,
		Aliases:       strings.Join(c.Aliases, ", "),
	}
}

// TableName returns a DBCategory table name.
func (dc *DBCategory) TableName() string {
	return "categories"
}
