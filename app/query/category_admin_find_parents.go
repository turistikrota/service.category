package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryAdminFindParentsQuery struct {
	MainUUIDs []string `json:"mainUUIDs" param:"mainUUIDs" query:"mainUUIDs" validate:"required,min=1,max=10,dive,object_id"`
}

type CategoryAdminFindParentsResult struct {
	List []*category.AdminListDto `json:"list"`
}

type CategoryAdminFindParentsHandler cqrs.HandlerFunc[CategoryAdminFindParentsQuery, *CategoryAdminFindParentsResult]

func NewCategoryAdminFindParentsHandler(repo category.Repository, cacheSrv cache.Service) CategoryAdminFindParentsHandler {
	cache := cache.New[[]*category.AdminListDto](cacheSrv)

	createCacheEntity := func() []*category.AdminListDto {
		return []*category.AdminListDto{}
	}

	return func(ctx context.Context, query CategoryAdminFindParentsQuery) (*CategoryAdminFindParentsResult, *i18np.Error) {
		cacheHandler := func() ([]*category.AdminListDto, *i18np.Error) {
			res, err := repo.FindByUUIDs(ctx, query.MainUUIDs)
			if err != nil {
				return nil, err
			}
			list := make([]*category.AdminListDto, len(res))
			for i, v := range res {
				list[i] = v.ToAdminList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_admin_find_parents_%v", query.MainUUIDs))
		if err != nil {
			return nil, err
		}
		return &CategoryAdminFindParentsResult{
			List: res,
		}, nil
	}
}
