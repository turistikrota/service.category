package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindFieldsQuery struct {
	UUIDs []string `json:"uuids" param:"uuids" query:"uuids" validate:"required,min=1,max=10,dive,object_id"`
}

type CategoryFindFieldsResult struct {
	Alerts      []*category.AlertDto      `json:"alerts"`
	Rules       []*category.RuleDto       `json:"rules"`
	InputGroups []*category.InputGroupDto `json:"inputGroups"`
}

type CategoryFindFieldsHandler cqrs.HandlerFunc[CategoryFindFieldsQuery, *CategoryFindFieldsResult]

func NewCategoryFindFieldsHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindFieldsHandler {
	cache := cache.New[*CategoryFindFieldsResult](cacheSrv)

	createCacheEntity := func() *CategoryFindFieldsResult {
		return &CategoryFindFieldsResult{}
	}

	return func(ctx context.Context, query CategoryFindFieldsQuery) (*CategoryFindFieldsResult, *i18np.Error) {
		cacheHandler := func() (*CategoryFindFieldsResult, *i18np.Error) {
			res, err := repo.FindFieldsByUUIDs(ctx, query.UUIDs)
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
			return &CategoryFindFieldsResult{
				InputGroups: inputGroups,
				Alerts:      alerts,
				Rules:       rules,
			}, nil
		}
		return cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_Fields_%v", query.UUIDs))
	}
}
