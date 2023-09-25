package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryDeleteCmd struct {
	AdminUUID    string `json:"-"`
	CategoryUUID string `json:"categoryUUID" param:"categoryUUID" validate:"required,object_id"`
}

type CategoryDeleteRes struct {
}

type CategoryDeleteHandler cqrs.HandlerFunc[CategoryDeleteCmd, *CategoryDeleteRes]

func NewCategoryDeleteHandler(repo category.Repository, events category.Events) CategoryDeleteHandler {
	return func(ctx context.Context, cmd CategoryDeleteCmd) (*CategoryDeleteRes, *i18np.Error) {
		err := repo.Delete(ctx, cmd.CategoryUUID)
		if err != nil {
			return nil, err
		}
		events.Deleted(category.DeletedEvent{
			AdminUUID:  cmd.AdminUUID,
			EntityUUID: cmd.CategoryUUID,
		})
		return &CategoryDeleteRes{}, nil
	}
}
