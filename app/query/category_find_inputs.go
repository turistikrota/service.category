package query

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/helpers/cache"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryFindInputsQuery struct {
	UUIDs []string `json:"uuids" param:"uuids" query:"uuids" validate:"required,min=1,max=10,dive,object_id"`
}

type CategoryFindInputsResult struct {
	RuleAndAlerts []*category.CategoryForAlertAndRuleDto `json:"ruleAndAlerts"`
	Inputs        []*category.InputGroupDto              `json:"inputs"`
}

type CategoryFindInputsHandler cqrs.HandlerFunc[CategoryFindInputsQuery, *CategoryFindInputsResult]

func NewCategoryFindInputsHandler(repo category.Repository, cacheSrv cache.Service) CategoryFindInputsHandler {
	cache := cache.New[*CategoryFindInputsResult](cacheSrv)

	createCacheEntity := func() *CategoryFindInputsResult {
		return &CategoryFindInputsResult{}
	}

	return func(ctx context.Context, query CategoryFindInputsQuery) (*CategoryFindInputsResult, *i18np.Error) {
		cacheHandler := func() (*CategoryFindInputsResult, *i18np.Error) {
			res, err := repo.FindInputsByUUIDs(ctx, query.UUIDs)
			if err != nil {
				return nil, err
			}
			result := make([]*category.InputGroupDto, 0)
			ruleAndAlerts := make([]*category.CategoryForAlertAndRuleDto, 0)
			for _, v := range res {
				result = append(result, v.ToInputGroup()...)
				ruleAndAlerts = append(ruleAndAlerts, v.ToCategoryForAlertAndRule())
			}
			return &CategoryFindInputsResult{
				Inputs:        result,
				RuleAndAlerts: ruleAndAlerts,
			}, nil
		}
		return cache.Creator(createCacheEntity).Handler(cacheHandler).Get(ctx, fmt.Sprintf("category_find_inputs_%v", query.UUIDs))
	}
}
