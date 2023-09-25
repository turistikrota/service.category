package command

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
	"github.com/turistikrota/service.category/domains/post"
)

type CategoryValidatePostCmd struct {
	Post *post.Entity    `json:"post" validate:"required"`
	User post.UserDetail `json:"user" validate:"required"`
}

type CategoryValidatePostRes struct {
}

type CategoryValidatePostHandler cqrs.HandlerFunc[CategoryValidatePostCmd, *CategoryValidatePostRes]

func NewCategoryValidatePostHandler(factory category.Factory, repo category.Repository, validators category.Validators, events category.Events) CategoryValidatePostHandler {
	return func(ctx context.Context, cmd CategoryValidatePostCmd) (*CategoryValidatePostRes, *i18np.Error) {
		categories, err := loadCategories(ctx, repo, cmd.Post.CategoryUUIDs)
		if err != nil {
			return nil, err
		}
		if len(categories) != len(cmd.Post.CategoryUUIDs) {
			return failValidation(events, "categories", factory.Errors.CategoryUUIDsIsNotCorrect(), cmd.Post, cmd.User)
		}
		index, err := validateCategoryChain(ctx, factory, categories)
		if err != nil {
			return failValidation(events, fmt.Sprintf("categories[%d]", index), err, cmd.Post, cmd.User)
		}
		for _, category := range categories {
			for _, input := range category.Inputs {
				feature, idx, exist := cmd.Post.GetFeatureByCategoryInputUUID(input.UUID)
				if !exist {
					if input.IsRequired {
						return failValidation(events, "features", factory.Errors.FeatureIsNotCorrect(), cmd.Post, cmd.User)
					}
				} else {
					validator, exist := validators.GetValidator(input.Type)
					if exist {
						err := validator.Validate(input, feature.Value)
						if err != nil {
							return failValidation(events, fmt.Sprintf("features[%v]", idx), err, cmd.Post, cmd.User)
						}
					}
					feature.IsPayed = input.IsPayed
				}
			}
		}
		return successValidation(events, cmd.Post, cmd.User)
	}
}

func loadCategories(ctx context.Context, repo category.Repository, categoryUUIDs []string) ([]*category.Entity, *i18np.Error) {
	res, err := repo.FindByUUIDs(ctx, categoryUUIDs)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func validateCategoryChain(ctx context.Context, factory category.Factory, categories []*category.Entity) (int, *i18np.Error) {
	len := len(categories)
	for i := 0; i < len; i++ {
		if i == 0 {
			if categories[i].MainUUID != "" {
				return i, factory.Errors.ParentUUIDIsNotEmpty()
			}
		} else {
			if categories[i].MainUUID != categories[i-1].UUID {
				return i, factory.Errors.ParentUUIDIsNotCorrect()
			}
		}
	}
	return -1, nil
}

func successValidation(events category.Events, p *post.Entity, u post.UserDetail) (*CategoryValidatePostRes, *i18np.Error) {
	events.PostValidationSuccess(category.PostValidationSuccessEvent{
		PostUUID: p.UUID,
		Post:     p,
		User: category.UserDetailEvent{
			UUID: u.UUID,
			Name: u.Name,
			Code: u.Code,
		},
	})
	return &CategoryValidatePostRes{}, nil
}

func failValidation(events category.Events, field string, err *i18np.Error, p *post.Entity, u post.UserDetail) (*CategoryValidatePostRes, *i18np.Error) {
	errors := make([]*post.ValidationError, 0)
	errors = append(errors, &post.ValidationError{
		Field:   field,
		Message: err.Key,
		Params:  *err.Params,
	})
	events.PostValidationFailed(category.PostValidationFailedEvent{
		PostUUID: p.UUID,
		Post:     p,
		Errors:   errors,
		User: category.UserDetailEvent{
			UUID: u.UUID,
			Name: u.Name,
			Code: u.Code,
		},
	})
	return &CategoryValidatePostRes{}, nil
}
