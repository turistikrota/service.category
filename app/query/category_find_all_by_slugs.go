package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindAllBySlugQuery struct {
	Locale string   `json:"-"`
	Slugs  []string `json:"slugs" param:"slugs" query:"slugs" validate:"omitempty,min=1,max=10,dive,slug"`
}

type CategoryFindAllBySlugResult struct {
	List []*category.ListDto `json:"list"`
}

type CategoryFindAllBySlugHandler cqrs.HandlerFunc[CategoryFindAllBySlugQuery, *CategoryFindAllBySlugResult]

func NewCategoryFindAllBySlugHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindAllBySlugHandler {
	c := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindAllBySlugQuery) (*CategoryFindAllBySlugResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.FindAllBySlugs(ctx, query.Locale, query.Slugs)
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
		}
		res, err := c.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_all_by_slugs_%v_%v", query.Locale, query.Slugs))
		if err != nil {
			return nil, err
		}
		return &CategoryFindAllBySlugResult{
			List: res,
		}, nil
	}
}
