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

type CategoryFindQuery struct {
	Locale       string
	CategoryUUID string `json:"uuid" params:"uuid" validate:"required,object_id"`
}

type CategoryFindResult struct {
	*category.AdminDetailDto
	MarkdownURL string `json:"markdownURL"`
}

type CategoryFindHandler cqrs.HandlerFunc[CategoryFindQuery, *CategoryFindResult]

func NewCategoryFindHandler(repo category.Repository, cacheSrv cache.Service, cnf config.App) CategoryFindHandler {
	cache := cache.New[*category.Entity](cacheSrv)

	createCacheEntity := func() *category.Entity {
		return &category.Entity{}
	}

	return func(ctx context.Context, query CategoryFindQuery) (*CategoryFindResult, *i18np.Error) {
		cacheHandler := func() (*category.Entity, *i18np.Error) {
			return repo.Find(ctx, query.CategoryUUID)
		}
		res, err := cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_%v", query.CategoryUUID))
		if err != nil {
			return nil, err
		}
		return &CategoryFindResult{
			AdminDetailDto: res.ToAdminDetail(),
			MarkdownURL:    dressCdnMarkdown(cnf, res.UUID, query.Locale),
		}, nil
	}
}

func dressCdnMarkdown(cnf config.App, identity string, locale string) string {
	return fmt.Sprintf("%s/categories/md/%s.%s.md", cnf.CDN.Url, identity, locale)
}
