package command

import (
	"context"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
)

type CategoryCreateCmd struct {
	AdminUUID   string                            `json:"-"`
	MainUUID    string                            `json:"mainUUID"  validate:"omitempty,object_id"`
	Images      []category.Image                  `json:"images" validate:"min=1,max=30,dive,required"`
	Meta        map[category.Locale]category.Meta `json:"meta" validate:"required,dive"`
	InputGroups []category.InputGroup             `json:"inputGroups" validate:"required,dive"`
	Inputs      []category.Input                  `json:"inputs" bson:"inputs" validate:"required,dive"`
	Rules       []category.Rule                   `json:"rules" bson:"rules" validate:"required,dive"`
	Alerts      []category.Alert                  `json:"alerts" bson:"alerts" validate:"required,dive"`
	Validators  []string                          `json:"validators" bson:"validators" validate:"required,min=1"`
	Order       int                               `json:"order" bson:"order" validate:"required,min=0,max=100"`
}

type CategoryCreateRes struct {
	UUID string `json:"uuid"`
}

type CategoryCreateHandler cqrs.HandlerFunc[CategoryCreateCmd, *CategoryCreateRes]

func NewCategoryCreateHandler(factory category.Factory, repo category.Repository, events category.Events) CategoryCreateHandler {
	return func(ctx context.Context, cmd CategoryCreateCmd) (*CategoryCreateRes, *i18np.Error) {
		e := factory.New(category.NewConfig{
			MainUUID:    cmd.MainUUID,
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
		saved, _err := repo.Create(ctx, e)
		if _err != nil {
			return nil, _err
		}
		events.Created(category.CreatedEvent{
			AdminUUID: cmd.AdminUUID,
			Entity:    saved,
		})
		return &CategoryCreateRes{
			UUID: saved.UUID,
		}, nil
	}
}
