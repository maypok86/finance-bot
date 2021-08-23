package repository

import (
	"context"

	"github.com/pkg/errors"
	"gitlab.com/LazyBearCT/finance-bot/internal/config"
	"gitlab.com/LazyBearCT/finance-bot/internal/repository/postgres"
)

// Database ...
type Database interface {
	Connect() error
	KeepAlive()
}

// Repository ...
type Repository struct {
	db       Database
	Category Category
	Budget   Budget
	Expense  Expense
}

// New creates a new Repository instance.
func New(ctx context.Context, config *config.DB) (*Repository, error) {
	var r *Repository
	switch config.Type {
	case "postgres":
		db, err := postgres.New(ctx, config)
		if err != nil {
			return nil, errors.Wrap(err, "[postgres.New] failed")
		}
		r = &Repository{
			db:       db,
			Category: postgres.NewCategoryRepository(db),
			Budget:   postgres.NewBudgetRepository(db),
			Expense:  postgres.NewExpenseRepository(db),
		}
	default:
		return nil, errors.New("unknown type db")
	}
	go r.db.KeepAlive()
	return r, nil
}
