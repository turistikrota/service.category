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
	MainUUID string `json:"mainUUID" param:"mainUUID" validate:"required,object_id"`
}

type CategoryAdminFindChildResult struct {
	List []*category.Entity `json:"list"`
}

type CategoryAdminFindChildHandler cqrs.HandlerFunc[CategoryAdminFindChildQuery, *CategoryAdminFindChildResult]

func NewCategoryAdminFindChildHandler(repo category.Repository, cacheSrv cache.Service) CategoryAdminFindChildHandler {
	cache := cache.New[[]*category.Entity](cacheSrv)

	createCacheEntity := func() []*category.Entity {
		return []*category.Entity{}
	}

	return func(ctx context.Context, query CategoryAdminFindChildQuery) (*CategoryAdminFindChildResult, *i18np.Error) {
		cacheHandler := func() ([]*category.Entity, *i18np.Error) {
			return repo.AdminFindChild(ctx, query.MainUUID)
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_admin_find_child_%v", query.MainUUID))
		if err != nil {
			return nil, err
		}
		return &CategoryAdminFindChildResult{
			List: res,
		}, nil
	}
}
