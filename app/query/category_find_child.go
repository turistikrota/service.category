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
	MainUUID string `json:"mainUUID" params:"mainUUID" validate:"required,object_id"`
}

type CategoryFindChildResult struct {
	List []*category.ListDto `json:"list"`
}

type CategoryFindChildHandler cqrs.HandlerFunc[CategoryFindChildQuery, *CategoryFindChildResult]

func NewCategoryFindChildHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindChildHandler {
	cache := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindChildQuery) (*CategoryFindChildResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.AdminFindChild(ctx, query.MainUUID)
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_admin_find_child_%v", query.MainUUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindChildResult{
			List: res,
		}, nil
	}
}
