package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryAdminFindAllQuery struct {
	OnlyMains bool `query:"only_mains"`
}

type CategoryAdminFindAllResult struct {
	List []*category.AdminListDto `json:"list"`
}

type CategoryAdminFindAllHandler cqrs.HandlerFunc[CategoryAdminFindAllQuery, *CategoryAdminFindAllResult]

func NewCategoryAdminFindAllHandler(repo category.Repository, cacheSrv cache.Service) CategoryAdminFindAllHandler {
	cache := cache.New[[]*category.AdminListDto](cacheSrv)

	createCacheEntity := func() []*category.AdminListDto {
		return []*category.AdminListDto{}
	}

	return func(ctx context.Context, query CategoryAdminFindAllQuery) (*CategoryAdminFindAllResult, *i18np.Error) {
		cacheHandler := func() ([]*category.AdminListDto, *i18np.Error) {
			res, err := repo.AdminFindAll(ctx, query.OnlyMains)
			if err != nil {
				return nil, err
			}
			list := make([]*category.AdminListDto, len(res))
			for i, v := range res {
				list[i] = v.ToAdminList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_admin_find_all"))
		if err != nil {
			return nil, err
		}
		return &CategoryAdminFindAllResult{
			List: res,
		}, nil
	}
}
