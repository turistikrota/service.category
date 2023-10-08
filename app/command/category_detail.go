package command

type CategoryDetailCmd struct {
	CategoryUUID string `json:"-" params:"uuid" validate:"required,object_id"`
}
