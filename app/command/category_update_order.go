package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryUpdateOrderCmd struct {
	AdminUUID    string `json:"-"`
	CategoryUUID string `json:"categoryUUID" param:"categoryUUID" validate:"required,object_id"`
	Order        int16  `json:"order" bson:"order" validate:"required,min=0,max=100"`
}

type CategoryUpdateOrderRes struct {
}

type CategoryUpdateOrderHandler cqrs.HandlerFunc[CategoryUpdateOrderCmd, *CategoryUpdateOrderRes]

func NewCategoryUpdateOrderHandler(repo category.Repository, events category.Events) CategoryUpdateOrderHandler {
	return func(ctx context.Context, cmd CategoryUpdateOrderCmd) (*CategoryUpdateOrderRes, *i18np.Error) {
		err := repo.UpdateOrder(ctx, cmd.CategoryUUID, cmd.Order)
		if err != nil {
			return nil, err
		}
		events.UpdateOrder(category.OrderUpdatedEvent{
			AdminUUID:  cmd.AdminUUID,
			EntityUUID: cmd.CategoryUUID,
			Order:      cmd.Order,
		})
		return &CategoryUpdateOrderRes{}, nil
	}
}
