package service

import (
	"context"
	"strings"

	"gitlab.com/LazyBearCT/finance-bot/internal/model"
	"gitlab.com/LazyBearCT/finance-bot/internal/repository"
)

//go:generate mockgen -source=category.go -destination=mocks/mock_category.go

// Category service.
type Category interface {
	GetAll() []*model.Category
	GetByName(name string) *model.Category
}

type categoryService struct {
	ctx        context.Context
	repo       repository.Category
	categories []*model.Category
}

// NewCategory creates a new Category instance.
func NewCategory(ctx context.Context, repo repository.Category) (Category, error) {
	category := &categoryService{
		ctx:  ctx,
		repo: repo,
	}
	if err := category.loadCategories(); err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *categoryService) loadCategories() error {
	categories, err := cs.repo.GetAllCategories(cs.ctx)
	if err != nil {
		return err
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
	return nil
}

// GetAll returns all model.Category.
func (cs *categoryService) GetAll() []*model.Category {
	return cs.categories
}

// GetByName returns model.Category by name.
func (cs *categoryService) GetByName(name string) *model.Category {
	var foundedCategory *model.Category
	var otherCategory *model.Category
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
