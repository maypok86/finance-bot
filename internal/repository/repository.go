package repository

import (
	"context"

	"github.com/LazyBearCT/finance-bot/internal/config"
	"github.com/LazyBearCT/finance-bot/internal/repository/postgres"
	"github.com/pkg/errors"
)

type Database interface {
	Connect() error
	KeepAlive()
}

type Repository struct {
	db       Database
	Category Category
	Budget   Budget
	Expense  Expense
}

func New(ctx context.Context, config *config.DB) (*Repository, error) {
	r := new(Repository)
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
