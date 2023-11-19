package command

import (
	"context"
	"fmt"

	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/turistikrota/service.category/domains/category"
	"github.com/turistikrota/service.category/domains/listing"
)

type CategoryValidateListingCmd struct {
	Listing *listing.Entity    `json:"listing" validate:"required"`
	User    listing.UserDetail `json:"user" validate:"required"`
}

type CategoryValidateListingRes struct {
}

type CategoryValidateListingHandler cqrs.HandlerFunc[CategoryValidateListingCmd, *CategoryValidateListingRes]

func NewCategoryValidateListingHandler(factory category.Factory, repo category.Repository, validators category.Validators, events category.Events) CategoryValidateListingHandler {
	return func(ctx context.Context, cmd CategoryValidateListingCmd) (*CategoryValidateListingRes, *i18np.Error) {
		categories, err := loadCategories(ctx, repo, cmd.Listing.CategoryUUIDs)
		if err != nil {
			return nil, err
		}
		if len(categories) != len(cmd.Listing.CategoryUUIDs) {
			return failValidation(events, "categories", factory.Errors.CategoryUUIDsIsNotCorrect(), cmd.Listing, cmd.User)
		}
		index, err := validateCategoryChain(ctx, factory, categories)
		if err != nil {
			return failValidation(events, fmt.Sprintf("categories[%d]", index), err, cmd.Listing, cmd.User)
		}
		for _, category := range categories {
			for _, input := range category.Inputs {
				feature, idx, exist := cmd.Listing.GetFeatureByCategoryInputUUID(input.UUID)
				if !exist {
					if input.IsRequired != nil && *input.IsRequired {
						return failValidation(events, "features", factory.Errors.FeatureIsNotCorrect(), cmd.Listing, cmd.User)
					}
				} else {
					validator, exist := validators.GetValidator(input.Type)
					if exist {
						err := validator.Validate(input, feature.Value)
						if err != nil {
							return failValidation(events, fmt.Sprintf("features[%v]", idx), err, cmd.Listing, cmd.User)
						}
					}
					isPayed := input.IsPayed != nil && *input.IsPayed
					feature.IsPayed = isPayed
				}
			}
		}
		cmd.Listing.IsValid = true
		return successValidation(events, cmd.Listing, cmd.User)
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
	length := len(categories)
	for i := length - 1; i >= 0; i-- {
		category := categories[i]
		if category.IsMain() {
			if i != 0 {
				return i, factory.Errors.CategoryUUIDsIsNotCorrect()
			}
		} else {
			if i == 0 {
				return i, factory.Errors.CategoryUUIDsIsNotCorrect()
			}
			if category.MainUUIDs[i-1] != categories[i-1].UUID {
				return i, factory.Errors.CategoryUUIDsIsNotCorrect()
			}
		}
	}
	return -1, nil
}

func successValidation(events category.Events, p *listing.Entity, u listing.UserDetail) (*CategoryValidateListingRes, *i18np.Error) {
	events.ListingValidationSuccess(category.ListingValidationSuccessEvent{
		ListingUUID: p.UUID,
		Listing:     p,
		User: category.UserDetailEvent{
			UUID: u.UUID,
			Name: u.Name,
		},
	})
	return &CategoryValidateListingRes{}, nil
}

func failValidation(events category.Events, field string, err *i18np.Error, p *listing.Entity, u listing.UserDetail) (*CategoryValidateListingRes, *i18np.Error) {
	errors := make([]*listing.ValidationError, 0)
	errors = append(errors, &listing.ValidationError{
		Field:   field,
		Message: err.Key,
		Params:  *err.Params,
	})
	events.ListingValidationFailed(category.ListingValidationFailedEvent{
		ListingUUID: p.UUID,
		Listing:     p,
		Errors:      errors,
		User: category.UserDetailEvent{
			UUID: u.UUID,
			Name: u.Name,
		},
	})
	return &CategoryValidateListingRes{}, nil
}
