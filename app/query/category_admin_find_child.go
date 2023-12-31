package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryAdminFindChildQuery struct {
	UUID string `json:"uuid" param:"uuid" validate:"required,object_id"`
}

type CategoryAdminFindChildResult struct {
	List []*category.AdminListDto `json:"list"`
}

type CategoryAdminFindChildHandler cqrs.HandlerFunc[CategoryAdminFindChildQuery, *CategoryAdminFindChildResult]

func NewCategoryAdminFindChildHandler(repo category.Repository, cacheSrv cache.Service) CategoryAdminFindChildHandler {
	cache := cache.New[[]*category.AdminListDto](cacheSrv)

	createCacheEntity := func() []*category.AdminListDto {
		return []*category.AdminListDto{}
	}

	return func(ctx context.Context, query CategoryAdminFindChildQuery) (*CategoryAdminFindChildResult, *i18np.Error) {
		cacheHandler := func() ([]*category.AdminListDto, *i18np.Error) {
			res, err := repo.AdminFindChild(ctx, query.UUID)
			if err != nil {
				return nil, err
			}
			list := make([]*category.AdminListDto, len(res))
			for i, v := range res {
				list[i] = v.ToAdminList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_admin_find_child_%v", query.UUID))
		if err != nil {
			return nil, err
		}
		return &CategoryAdminFindChildResult{
			List: res,
		}, nil
	}
}
