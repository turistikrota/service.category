package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindBySlugQuery struct {
	Locale string `json:"-"`
	Slug   string `json:"slug" param:"slug" validate:"required,slug"`
}

type CategoryFindBySlugResult struct {
	*category.Entity
}

type CategoryFindBySlugHandler cqrs.HandlerFunc[CategoryFindBySlugQuery, *CategoryFindBySlugResult]

func NewCategoryFindBySlugHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindBySlugHandler {
	cache := cache.New[*category.Entity](cacheSrv)

	createCacheEntity := func() *category.Entity {
		return &category.Entity{}
	}

	return func(ctx context.Context, query CategoryFindBySlugQuery) (*CategoryFindBySlugResult, *i18np.Error) {
		cacheHandler := func() (*category.Entity, *i18np.Error) {
			return repo.FindBySlug(ctx, category.I18nDetail{
				Locale: query.Locale,
				Slug:   query.Slug,
			})
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_%v_%v", query.Locale, query.Slug))
		if err != nil {
			return nil, err
		}
		return &CategoryFindBySlugResult{
			Entity: res,
		}, nil
	}
}
