package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindChildByUUIDQuery struct {
	UUID string `json:"uuid" params:"uuid" validate:"required,object_id"`
}

type CategoryFindChildByUUIDResult struct {
	List []*category.ListDto `json:"list"`
}

type CategoryFindChildByUUIDHandler cqrs.HandlerFunc[CategoryFindChildByUUIDQuery, *CategoryFindChildByUUIDResult]

func NewCategoryFindChildByUUIDHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindChildByUUIDHandler {
	cache := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindChildByUUIDQuery) (*CategoryFindChildByUUIDResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.AdminFindChildByUUID(ctx, query.UUID)
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_child_by_uuid_%v", query.UUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindChildByUUIDResult{
			List: res,
		}, nil
	}
}
