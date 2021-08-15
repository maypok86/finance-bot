package postgres

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/model"
	"gorm.io/gorm"
)

type CategoryPostgres struct {
	db *gorm.DB
}

func NewCategoryRepository(db *DB) *CategoryPostgres {
	return &CategoryPostgres{
		db: db.db,
	}
}

func (c *CategoryPostgres) GetAllCategories(ctx context.Context) ([]*model.DBCategory, error) {
	var categories []*model.DBCategory
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
