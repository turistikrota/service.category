package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryUpdateCmd struct {
	AdminUUID    string                             `json:"-"`
	CategoryUUID string                             `json:"-" params:"uuid" validate:"required,object_id"`
	MainUUIDs    []string                           `json:"mainUUIDs"  validate:"required,dive,object_id"`
	Images       []category.Image                   `json:"images" validate:"min=1,max=30,dive,required"`
	Meta         map[category.Locale]*category.Meta `json:"meta" validate:"required,dive"`
	InputGroups  []category.InputGroup              `json:"inputGroups" validate:"required,dive"`
	Inputs       []category.Input                   `json:"inputs" bson:"inputs" validate:"required,dive"`
	Rules        []category.Rule                    `json:"rules" bson:"rules" validate:"required,dive"`
	Alerts       []category.Alert                   `json:"alerts" bson:"alerts" validate:"required,dive"`
	Validators   []string                           `json:"validators" bson:"validators" validate:"required,min=1"`
	Order        int                                `json:"order" bson:"order" validate:"required,min=0,max=100"`
}

type CategoryUpdateRes struct {
}

type CategoryUpdateHandler cqrs.HandlerFunc[CategoryUpdateCmd, *CategoryUpdateRes]

func NewCategoryUpdateHandler(factory category.Factory, repo category.Repository, events category.Events) CategoryUpdateHandler {
	return func(ctx context.Context, cmd CategoryUpdateCmd) (*CategoryUpdateRes, *i18np.Error) {
		e := factory.New(category.NewConfig{
			MainUUIDs:   cmd.MainUUIDs,
			Meta:        cmd.Meta,
			Images:      cmd.Images,
			Inputs:      cmd.Inputs,
			InputGroups: cmd.InputGroups,
			Rules:       cmd.Rules,
			Alerts:      cmd.Alerts,
			Validators:  cmd.Validators,
			Order:       cmd.Order,
		})
		err := factory.Validate(e)
		if err != nil {
			return nil, err
		}
		e.UUID = cmd.CategoryUUID
		err = repo.Update(ctx, e)
		if err != nil {
			return nil, err
		}
		events.Updated(category.UpdatedEvent{
			AdminUUID: cmd.AdminUUID,
			Entity:    e,
		})
		return &CategoryUpdateRes{}, nil
	}
}
