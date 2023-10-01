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
	List []*category.ListDto `json:"list"`
}

type CategoryFindAllHandler cqrs.HandlerFunc[CategoryFindAllQuery, *CategoryFindAllResult]

func NewCategoryFindAllHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindAllHandler {
	cache := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindAllQuery) (*CategoryFindAllResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.FindAll(ctx)
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
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
