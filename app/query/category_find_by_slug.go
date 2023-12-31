package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/config"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindBySlugQuery struct {
	Locale string `json:"-"`
	Slug   string `json:"slug" param:"slug" validate:"required,slug"`
}

type CategoryFindBySlugResult struct {
	*category.DetailDto
	MarkdownURL string `json:"markdownURL"`
}

type CategoryFindBySlugHandler cqrs.HandlerFunc[CategoryFindBySlugQuery, *CategoryFindBySlugResult]

func NewCategoryFindBySlugHandler(repo category.Repository, cacheSrv cache.Service, cnf config.App) CategoryFindBySlugHandler {
	cache := cache.New[*category.DetailDto](cacheSrv)

	createCacheEntity := func() *category.DetailDto {
		return &category.DetailDto{}
	}

	return func(ctx context.Context, query CategoryFindBySlugQuery) (*CategoryFindBySlugResult, *i18np.Error) {
		cacheHandler := func() (*category.DetailDto, *i18np.Error) {
			res, err := repo.FindBySlug(ctx, category.I18nDetail{
				Locale: query.Locale,
				Slug:   query.Slug,
			})
			if err != nil {
				return nil, err
			}
			return res.ToDetail(), nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_%v_%v", query.Locale, query.Slug))
		if err != nil {
			return nil, err
		}
		return &CategoryFindBySlugResult{
			DetailDto:   res,
			MarkdownURL: dressCdnMarkdown(cnf, res.UUID, query.Locale),
		}, nil
	}
}
