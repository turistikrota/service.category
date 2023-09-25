package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindAllQuery struct{}

type CategoryFindAllResult struct {
	List []*category.Entity `json:"list"`
}

type CategoryFindAllHandler cqrs.HandlerFunc[CategoryFindAllQuery, *CategoryFindAllResult]

func NewCategoryFindAllHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindAllHandler {
	cache := cache.New[[]*category.Entity](cacheSrv)

	createCacheEntity := func() []*category.Entity {
		return []*category.Entity{}
	}

	return func(ctx context.Context, query CategoryFindAllQuery) (*CategoryFindAllResult, *i18np.Error) {
		cacheHandler := func() ([]*category.Entity, *i18np.Error) {
			return repo.FindAll(ctx)
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_all"))
		if err != nil {
			return nil, err
		}
		return &CategoryFindAllResult{
			List: res,
		}, nil
	}
}
