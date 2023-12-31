package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindFieldsByUUIDsQuery struct {
	UUIDs []string `json:"uuids" param:"uuids" query:"uuids" validate:"required,min=1,max=10,dive,object_id"`
}

type CategoryFindFieldsByUUIDsResult struct {
	Alerts      []*category.AlertDto      `json:"alerts"`
	Rules       []*category.RuleDto       `json:"rules"`
	InputGroups []*category.InputGroupDto `json:"inputGroups"`
}

type CategoryFindFieldsByUUIDsHandler cqrs.HandlerFunc[CategoryFindFieldsByUUIDsQuery, *CategoryFindFieldsByUUIDsResult]

func NewCategoryFindFieldsByUUIDsHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindFieldsByUUIDsHandler {
	c := cache.New[*CategoryFindFieldsByUUIDsResult](cacheSrv)

	createCacheEntity := func() *CategoryFindFieldsByUUIDsResult {
		return &CategoryFindFieldsByUUIDsResult{}
	}

	return func(ctx context.Context, query CategoryFindFieldsByUUIDsQuery) (*CategoryFindFieldsByUUIDsResult, *i18np.Error) {
		cacheHandler := func() (*CategoryFindFieldsByUUIDsResult, *i18np.Error) {
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
			return &CategoryFindFieldsByUUIDsResult{
				InputGroups: inputGroups,
				Alerts:      alerts,
				Rules:       rules,
			}, nil
		}
		return c.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_fields_by_uuids_%v", query.UUIDs))
	}
}
