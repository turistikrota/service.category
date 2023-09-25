package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindChildQuery struct {
	MainUUID string `json:"mainUUID" param:"mainUUID" validate:"required,object_id"`
}

type CategoryFindChildResult struct {
	List []*category.Entity `json:"list"`
}

type CategoryFindChildHandler cqrs.HandlerFunc[CategoryFindChildQuery, *CategoryFindChildResult]

func NewCategoryFindChildHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindChildHandler {
	cache := cache.New[[]*category.Entity](cacheSrv)

	createCacheEntity := func() []*category.Entity {
		return []*category.Entity{}
	}

	return func(ctx context.Context, query CategoryFindChildQuery) (*CategoryFindChildResult, *i18np.Error) {
		cacheHandler := func() ([]*category.Entity, *i18np.Error) {
			return repo.FindChild(ctx, query.MainUUID)
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_child_%v", query.MainUUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindChildResult{
			List: res,
		}, nil
	}
}
