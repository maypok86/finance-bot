package service

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/repository"
	"github.com/pkg/errors"
)

type Manager struct {
	Category Category
}

func NewManager(ctx context.Context, repo *repository.Repository) (*Manager, error) {
	if repo == nil {
		return nil, errors.New("No repo provided")
	}
	return &Manager{
		Category: NewCategory(ctx, repo.Category),
	}, nil
}
