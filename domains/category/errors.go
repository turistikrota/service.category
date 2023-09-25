package category

import "github.com/cilloparch/cillop/i18np"

type Errors interface {
	InvalidInputType(t string) *i18np.Error
	InvalidInputUUID(uuid string) *i18np.Error
	InvalidInputGroupUUID(uuid string) *i18np.Error
	Failed(action string) *i18np.Error
	NotFound() *i18np.Error
	MinOneTranslationRequired(field string) *i18np.Error
	InvalidMetaLength(len int) *i18np.Error
	InvalidImagesLength(len int) *i18np.Error
	InvalidUUID(uuid string) *i18np.Error
	InvalidCategoryInputType(t string, value interface{}) *i18np.Error
	ParentUUIDIsNotEmpty() *i18np.Error
	ParentUUIDIsNotCorrect() *i18np.Error
	CategoryUUIDsIsNotCorrect() *i18np.Error
	FeatureIsNotCorrect() *i18np.Error
}

type categoryErrors struct{}

func newCategoryErrors() Errors {
	return &categoryErrors{}
}

func (e *categoryErrors) InvalidInputType(t string) *i18np.Error {
	return i18np.NewError(messages.InvalidInputType, i18np.P{"Type": t})
}

func (e *categoryErrors) InvalidInputUUID(uuid string) *i18np.Error {
	return i18np.NewError(messages.InvalidInputUUID, i18np.P{"UUID": uuid})
}

func (e *categoryErrors) InvalidInputGroupUUID(uuid string) *i18np.Error {
	return i18np.NewError(messages.InvalidInputGroupUUID, i18np.P{"UUID": uuid})
}

func (e *categoryErrors) Failed(action string) *i18np.Error {
	return i18np.NewError(messages.Failed, i18np.P{"Action": action})
}

func (e *categoryErrors) NotFound() *i18np.Error {
	return i18np.NewError(messages.NotFound)
}

func (e *categoryErrors) MinOneTranslationRequired(field string) *i18np.Error {
	return i18np.NewError(messages.MinOneTranslationRequired, i18np.P{"Field": field})
}

func (e *categoryErrors) InvalidMetaLength(len int) *i18np.Error {
	return i18np.NewError(messages.InvalidMetaLength, i18np.P{"Length": len})
}

func (e *categoryErrors) InvalidImagesLength(len int) *i18np.Error {
	return i18np.NewError(messages.InvalidImagesLength, i18np.P{"Length": len})
}

func (e *categoryErrors) InvalidUUID(uuid string) *i18np.Error {
	return i18np.NewError(messages.InvalidUUID, i18np.P{"UUID": uuid})
}

func (e *categoryErrors) InvalidCategoryInputType(t string, value interface{}) *i18np.Error {
	return i18np.NewError(messages.InvalidCategoryInputType, i18np.P{"Type": t, "Value": value})
}

func (e *categoryErrors) ParentUUIDIsNotEmpty() *i18np.Error {
	return i18np.NewError(messages.ParentUUIDIsNotEmpty)
}

func (e *categoryErrors) ParentUUIDIsNotCorrect() *i18np.Error {
	return i18np.NewError(messages.ParentUUIDIsNotCorrect)
}

func (e *categoryErrors) CategoryUUIDsIsNotCorrect() *i18np.Error {
	return i18np.NewError(messages.CategoryUUIDsIsNotCorrect)
}

func (e *categoryErrors) FeatureIsNotCorrect() *i18np.Error {
	return i18np.NewError(messages.FeatureIsNotCorrect)
}
