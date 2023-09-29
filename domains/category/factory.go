package category

import (
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
	"github.com/turistikrota/service.category/domains/post"
)

type Factory struct {
	Errors Errors
}

func NewFactory() Factory {
	return Factory{
		Errors: newCategoryErrors(),
	}
}

func (f Factory) IsZero() bool {
	return f.Errors == nil
}

type NewConfig struct {
	MainUUID    string
	Images      []Image
	Meta        map[Locale]*Meta
	InputGroups []InputGroup
	Inputs      []Input
	Rules       []Rule
	Alerts      []Alert
	Validators  []string
	Order       int
}

func (f Factory) New(cnf NewConfig) *Entity {
	t := time.Now()
	return &Entity{
		MainUUID:    cnf.MainUUID,
		Meta:        cnf.Meta,
		Images:      cnf.Images,
		InputGroups: cnf.InputGroups,
		Inputs:      cnf.Inputs,
		Rules:       cnf.Rules,
		Alerts:      cnf.Alerts,
		Order:       cnf.Order,
		IsActive:    false,
		IsDeleted:   false,
		CreatedAt:   t,
		UpdatedAt:   t,
	}
}

func (f Factory) Validate(entity *Entity) *i18np.Error {
	if err := f.validateInputGroups(entity); err != nil {
		return err
	}
	if err := f.validateInputs(entity); err != nil {
		return err
	}
	return nil
}

func (f Factory) validateInputGroups(entity *Entity) *i18np.Error {
	for _, group := range entity.InputGroups {
		if err := f.validateInputGroupUUIDs(group); err != nil {
			return err
		}
	}
	return nil
}

func (f Factory) validateInputs(entity *Entity) *i18np.Error {
	for _, input := range entity.Inputs {
		if err := f.validateInputType(input.Type); err != nil {
			return err
		}
		if err := f.validateInputUUIDs(input); err != nil {
			return err
		}
	}
	return nil
}

func (f Factory) validateInputGroupUUIDs(group InputGroup) *i18np.Error {
	if !IsUUID(group.UUID) {
		return f.Errors.InvalidInputGroupUUID(group.UUID)
	}
	return nil
}

func (f Factory) validateInputUUIDs(input Input) *i18np.Error {
	if !IsUUID(input.UUID) {
		return f.Errors.InvalidInputUUID(input.UUID)
	}
	return nil
}

func (f Factory) validatePostFeatureByInputType(input Input, feature post.Feature) *i18np.Error {
	return nil
}

func (f Factory) validateInputType(t InputType) *i18np.Error {
	types := []string{
		InputTypeText.String(),
		InputTypeTextarea.String(),
		InputTypeNumber.String(),
		InputTypeSelect.String(),
		InputTypeRadio.String(),
		InputTypeCheckbox.String(),
		InputTypeDate.String(),
		InputTypeTime.String(),
		InputTypeDatetime.String(),
		InputTypeFile.String(),
		InputTypeImage.String(),
		InputTypePDF.String(),
		InputTypeRange.String(),
		InputTypeColor.String(),
		InputTypeURL.String(),
		InputTypeEmail.String(),
		InputTypeTel.String(),
		InputTypeLocation.String(),
		InputTypePrice.String(),
		InputTypeRating.String(),
	}
	for _, v := range types {
		if v == t.String() {
			return nil
		}
	}
	return f.Errors.InvalidInputType(t.String())
}

func NewUUID() string {
	return uuid.New().String()
}

func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
