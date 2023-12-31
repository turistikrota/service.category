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

type CategoryFindByUUIDQuery struct {
	Locale string `json:"-"`
	UUID   string `json:"uuid" param:"uuid" validate:"required,object_id"`
}

type CategoryFindByUUIDMarkdownResult struct {
	TR string `json:"tr"`
	EN string `json:"en"`
}

type CategoryFindByUUIDResult struct {
	*category.DetailDto
	Markdown *CategoryFindByUUIDMarkdownResult `json:"md"`
}

type CategoryFindByUUIDHandler cqrs.HandlerFunc[CategoryFindByUUIDQuery, *CategoryFindByUUIDResult]

func NewCategoryFindByUUIDHandler(repo category.Repository, cacheSrv cache.Service, cnf config.App) CategoryFindByUUIDHandler {
	cache := cache.New[*category.DetailDto](cacheSrv)

	createCacheEntity := func() *category.DetailDto {
		return &category.DetailDto{}
	}

	return func(ctx context.Context, query CategoryFindByUUIDQuery) (*CategoryFindByUUIDResult, *i18np.Error) {
		cacheHandler := func() (*category.DetailDto, *i18np.Error) {
			res, err := repo.Find(ctx, query.UUID)
			if err != nil {
				return nil, err
			}
			return res.ToDetail(), nil
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_%v", query.UUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindByUUIDResult{
			DetailDto: res,
			Markdown: &CategoryFindByUUIDMarkdownResult{
				TR: dressCdnMarkdown(cnf, res.UUID, category.LocaleTR.String()),
				EN: dressCdnMarkdown(cnf, res.UUID, category.LocaleEN.String()),
			},
		}, nil
	}
}
