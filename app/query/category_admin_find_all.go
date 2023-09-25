package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryAdminFindAllQuery struct{}

type CategoryAdminFindAllResult struct {
	List []*category.Entity `json:"list"`
}

type CategoryAdminFindAllHandler cqrs.HandlerFunc[CategoryAdminFindAllQuery, *CategoryAdminFindAllResult]

func NewCategoryAdminFindAllHandler(repo category.Repository, cacheSrv cache.Service) CategoryAdminFindAllHandler {
	cache := cache.New[[]*category.Entity](cacheSrv)

	createCacheEntity := func() []*category.Entity {
		return []*category.Entity{}
	}

	return func(ctx context.Context, query CategoryAdminFindAllQuery) (*CategoryAdminFindAllResult, *i18np.Error) {
		cacheHandler := func() ([]*category.Entity, *i18np.Error) {
			return repo.AdminFindAll(ctx)
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
