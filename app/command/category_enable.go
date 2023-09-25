package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryEnableCmd struct {
	AdminUUID    string `json:"-"`
	CategoryUUID string `json:"categoryUUID" param:"categoryUUID" validate:"required,object_id"`
}

type CategoryEnableRes struct {
}

type CategoryEnableHandler cqrs.HandlerFunc[CategoryEnableCmd, *CategoryEnableRes]

func NewCategoryEnableHandler(repo category.Repository, events category.Events) CategoryEnableHandler {
	return func(ctx context.Context, cmd CategoryEnableCmd) (*CategoryEnableRes, *i18np.Error) {
		err := repo.Enable(ctx, cmd.CategoryUUID)
		if err != nil {
			return nil, err
		}
		events.Enabled(category.EnabledEvent{
			AdminUUID:  cmd.AdminUUID,
			EntityUUID: cmd.CategoryUUID,
		})
		return &CategoryEnableRes{}, nil
	}
}
