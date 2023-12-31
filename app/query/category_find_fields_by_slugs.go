package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindFieldsBySlugsQuery struct {
	Locale string   `json:"-"`
	Slugs  []string `json:"slugs" param:"slugs" query:"slugs" validate:"required,min=1,max=10,dive,slug"`
}

type CategoryFindFieldsBySlugsResult struct {
	Alerts      []*category.AlertDto      `json:"alerts"`
	Rules       []*category.RuleDto       `json:"rules"`
	InputGroups []*category.InputGroupDto `json:"inputGroups"`
}

type CategoryFindFieldsBySlugsHandler cqrs.HandlerFunc[CategoryFindFieldsBySlugsQuery, *CategoryFindFieldsBySlugsResult]

func NewCategoryFindFieldsBySlugsHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindFieldsBySlugsHandler {
	c := cache.New[*CategoryFindFieldsBySlugsResult](cacheSrv)

	createCacheEntity := func() *CategoryFindFieldsBySlugsResult {
		return &CategoryFindFieldsBySlugsResult{}
	}

	return func(ctx context.Context, query CategoryFindFieldsBySlugsQuery) (*CategoryFindFieldsBySlugsResult, *i18np.Error) {
		cacheHandler := func() (*CategoryFindFieldsBySlugsResult, *i18np.Error) {
			res, err := repo.FindFieldsBySlugs(ctx, query.Locale, query.Slugs)
			if err != nil {
				return nil, err
			}
			inputGroups := make([]*category.InputGroupDto, 0)
			alerts := make([]*category.AlertDto, 0)
			rules := make([]*category.RuleDto, 0)
			for _, v := range res {
				inputGroups = append(inputGroups, v.ToInputGroup()...)
				alerts = append(alerts, v.ToAlert()...)
				rules = append(rules, v.ToRule()...)
			}
			return &CategoryFindFieldsBySlugsResult{
				InputGroups: inputGroups,
				Alerts:      alerts,
				Rules:       rules,
			}, nil
		}
		return c.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_fields_by_slugs_%v_%v", query.Locale, query.Slugs))
	}
}
