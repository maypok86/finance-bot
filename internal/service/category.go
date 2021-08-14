package service

import (
	"context"
	"strings"
	"sync"

	"github.com/LazyBearCT/finance-bot/internal/model"
	"github.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=category.go -destination=mocks/mock_category.go

type Category interface {
	GetAll() ([]*model.Category, error)
	GetByName(name string) *model.Category
}

type categoryService struct {
	ctx        context.Context
	repo       repository.Category
	categories []*model.Category
}

func NewCategory(ctx context.Context, repo repository.Category) Category {
	return &categoryService{
		ctx:  ctx,
		repo: repo,
	}
}

var once sync.Once
var e error

func (cs *categoryService) GetAll() ([]*model.Category, error) {
	once.Do(func() {
		categories, err := cs.repo.GetAllCategories(cs.ctx)
		if err != nil {
			e = err
			return
		}
		for _, category := range categories {
			aliases := filterAliases(strings.Split(category.Aliases, ", "))
			aliases = append(aliases, category.Codename, category.Name)
			cs.categories = append(cs.categories, &model.Category{
				Codename:      category.Codename,
				Name:          category.Name,
				IsBaseExpense: category.IsBaseExpense,
				Aliases:       aliases,
			})
		}
	})
	return cs.categories, e
}

func (cs *categoryService) GetByName(name string) *model.Category {
	var foundedCategory *model.Category = nil
	var otherCategory *model.Category = nil
	for _, category := range cs.categories {
		if category.Codename == "other" {
			otherCategory = category
		}
		for _, alias := range category.Aliases {
			if strings.Contains(alias, name) {
				foundedCategory = category
			}
		}
	}
	if foundedCategory == nil {
		foundedCategory = otherCategory
	}
	return foundedCategory
}

func filterAliases(s []string) []string {
	var aliases []string
	for _, str := range s {
		alias := strings.TrimSpace(str)
		if alias != "" {
			aliases = append(aliases, alias)
		}
	}
	return aliases
}
