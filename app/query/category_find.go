package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindQuery struct {
	CategoryUUID string `json:"categoryUUID" param:"categoryUUID" validate:"required,object_id"`
}

type CategoryFindResult struct {
	*category.Entity
}

type CategoryFindHandler cqrs.HandlerFunc[CategoryFindQuery, *CategoryFindResult]

func NewCategoryFindHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindHandler {
	cache := cache.New[*category.Entity](cacheSrv)

	createCacheEntity := func() *category.Entity {
		return &category.Entity{}
	}

	return func(ctx context.Context, query CategoryFindQuery) (*CategoryFindResult, *i18np.Error) {
		cacheHandler := func() (*category.Entity, *i18np.Error) {
			return repo.Find(ctx, query.CategoryUUID)
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_%v", query.CategoryUUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindResult{
			Entity: res,
		}, nil
	}
}