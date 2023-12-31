package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindAllByUUIDsQuery struct {
	UUIDs []string `json:"uuids" param:"uuids" query:"uuids" validate:"omitempty,min=1,max=10,dive,object_id"`
}

type CategoryFindAllByUUIDsResult struct {
	List []*category.ListDto `json:"list"`
}

type CategoryFindAllByUUIDsHandler cqrs.HandlerFunc[CategoryFindAllByUUIDsQuery, *CategoryFindAllByUUIDsResult]

func NewCategoryFindAllByUUIDsHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindAllByUUIDsHandler {
	c := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindAllByUUIDsQuery) (*CategoryFindAllByUUIDsResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.FindAllByUUIDs(ctx, query.UUIDs)
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
		}
		res, err := c.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_all_by_uuids_%v", query.UUIDs))
		if err != nil {
			return nil, err
		}
		return &CategoryFindAllByUUIDsResult{
			List: res,
		}, nil
	}
}
