package repository

import (
	"context"

	"github.com/maypok86/finance-bot/internal/model"
)

//go:generate mockgen -source=category.go -destination=mocks/mock_category.go

// Category repository.
type Category interface {
	GetAllCategories(ctx context.Context) ([]*model.DBCategory, error)
}
