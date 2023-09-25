package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryDisableCmd struct {
	AdminUUID    string `json:"-"`
	CategoryUUID string `json:"categoryUUID" param:"categoryUUID" validate:"required,object_id"`
}

type CategoryDisableRes struct {
}

type CategoryDisableHandler cqrs.HandlerFunc[CategoryDisableCmd, *CategoryDisableRes]

func NewCategoryDisableHandler(repo category.Repository, events category.Events) CategoryDisableHandler {
	return func(ctx context.Context, cmd CategoryDisableCmd) (*CategoryDisableRes, *i18np.Error) {
		err := repo.Disable(ctx, cmd.CategoryUUID)
		if err != nil {
			return nil, err
		}
		events.Disabled(category.DisabledEvent{
			AdminUUID:  cmd.AdminUUID,
			EntityUUID: cmd.CategoryUUID,
		})
		return &CategoryDisableRes{}, nil
	}
}
