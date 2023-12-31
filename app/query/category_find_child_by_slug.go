package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindChildBySlugQuery struct {
	Locale string `json:"-"`
	Slug   string `json:"slug" params:"slug" validate:"required,slug"`
}

type CategoryFindChildBySlugResult struct {
	List []*category.ListDto `json:"list"`
}

type CategoryFindChildBySlugHandler cqrs.HandlerFunc[CategoryFindChildBySlugQuery, *CategoryFindChildBySlugResult]

func NewCategoryFindChildBySlugHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindChildBySlugHandler {
	cache := cache.New[[]*category.ListDto](cacheSrv)

	createCacheEntity := func() []*category.ListDto {
		return []*category.ListDto{}
	}

	return func(ctx context.Context, query CategoryFindChildBySlugQuery) (*CategoryFindChildBySlugResult, *i18np.Error) {
		cacheHandler := func() ([]*category.ListDto, *i18np.Error) {
			res, err := repo.AdminFindChildBySlug(ctx, category.I18nDetail{
				Locale: query.Locale,
				Slug:   query.Slug,
			})
			if err != nil {
				return nil, err
			}
			list := make([]*category.ListDto, len(res))
			for i, v := range res {
				list[i] = v.ToList()
			}
			return list, nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_child_by_slug_%v_%v", query.Locale, query.Slug))
		if err != nil {
			return nil, err
		}
		return &CategoryFindChildBySlugResult{
			List: res,
		}, nil
	}
}
