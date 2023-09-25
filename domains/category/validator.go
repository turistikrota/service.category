package category

import "github.com/cilloparch/cillop/i18np"

type ValidatorFunc func(input Input, value interface{}) *i18np.Error

type Validator interface {
	Validate(input Input, value interface{}) *i18np.Error
	GetName() string
	GetType() InputType
}

type validatorImpl struct {
	Type      InputType
	Validator ValidatorFunc
}

func (v *validatorImpl) Validate(input Input, value interface{}) *i18np.Error {
	return v.Validator(input, value)
}

func (v *validatorImpl) GetName() string {
	return v.Type.String()
}

func (v *validatorImpl) GetType() InputType {
	return v.Type
}

func newValidator(t InputType, validator ValidatorFunc) Validator {
	return &validatorImpl{
		Type:      t,
		Validator: validator,
	}
}
