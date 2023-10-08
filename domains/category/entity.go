package category

import (
	"time"

	"github.com/ssibrahimbas/slug"
)

type Entity struct {
	UUID        string           `json:"uuid" bson:"_id,omitempty"`
	MainUUID    string           `json:"mainUUID" bson:"main_uuid,omitempty"  validate:"omitempty,object_id"`
	MainUUIDs   []string         `json:"mainUUIDs" bson:"main_uuids"  validate:"omitempty,dive,object_id"`
	Images      []Image          `json:"images" bson:"images"  validate:"min=1,max=30,dive,required"`
	Meta        map[Locale]*Meta `json:"meta" bson:"meta" validate:"required,dive"`
	InputGroups []InputGroup     `json:"inputGroups" bson:"input_groups" validate:"required,dive"`
	Inputs      []Input          `json:"inputs" bson:"inputs" validate:"required,dive"`
	Rules       []Rule           `json:"rules" bson:"rules" validate:"required,dive"`
	Alerts      []Alert          `json:"alerts" bson:"alerts" validate:"required,dive"`
	Validators  []string         `json:"validators" bson:"validators" validate:"required,min=1"`
	Order       int              `json:"order" bson:"order" validate:"required,min=0,max=100"`
	IsActive    bool             `json:"isActive" bson:"is_active" validate:"required,boolean"`
	IsDeleted   bool             `json:"isDeleted" bson:"is_deleted" validate:"required,boolean"`
	CreatedAt   time.Time        `json:"createdAt" bson:"created_at" validate:"required"`
	UpdatedAt   time.Time        `json:"updatedAt" bson:"updated_at" validate:"required"`
}

type Image struct {
	Url   string `json:"url" bson:"url" validate:"required,url"`
	Order int16  `json:"order" bson:"order" validate:"required,min=0,max=20"`
}

type BaseTranslation struct {
	Name        string `json:"name" bson:"name" validate:"required,max=255,min=3"`
	Description string `json:"description" bson:"description" validate:"required,max=255,min=3"`
}

type Rule struct {
	UUID         string                     `json:"uuid" bson:"uuid" validate:"required,uuid4"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations" validate:"required,dive"`
	StrictLevel  int16                      `json:"strictLevel" bson:"strict_level"  validate:"required,min=0,max=10"`
}

type Alert struct {
	UUID         string                     `json:"uuid" bson:"uuid" validate:"required,uuid4"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations"`
	Type         string                     `json:"type" bson:"type"` // info, warning, error
}

type InputGroup struct {
	UUID         string                     `json:"uuid" bson:"uuid" validate:"required,uuid4"`
	Icon         string                     `json:"icon" bson:"icon" validate:"required,max=255,min=3"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations" validate:"required,dive"`
}
type Input struct {
	UUID         string                       `json:"uuid" bson:"uuid" validate:"required,uuid4"`
	GroupUUID    string                       `json:"groupUUID" bson:"group_uuid"`
	Type         InputType                    `json:"type" bson:"type"  validate:"required"`
	Translations map[Locale]*InputTranslation `json:"translations" bson:"translations" validate:"required,dive"`
	IsRequired   bool                         `json:"isRequired" bson:"is_required"  validate:"required,boolean"`
	IsMultiple   bool                         `json:"isMultiple" bson:"is_multiple"  validate:"required,boolean"`
	IsUnique     bool                         `json:"isUnique" bson:"is_unique"  validate:"required,boolean"`
	IsPayed      bool                         `json:"isPayed" bson:"is_payed"  validate:"required,boolean"`
	Extra        []InputExtra                 `json:"extra" bson:"extra"  validate:"required,dive"`
	Options      []string                     `json:"options" bson:"options"  validate:"required,min=0"`
}

type InputExtra struct {
	Name  string `json:"name" bson:"name" validate:"required,max=255,min=1"`
	Value string `json:"value" bson:"value" validate:"required,max=255,min=1"`
}

type InputTranslation struct {
	Name        string `json:"name" bson:"name" validate:"required,max=255,min=3"`
	Placeholder string `json:"placeholder" bson:"placeholder" validate:"required,max=255,min=3"`
	Help        string `json:"help" bson:"help" validate:"required,max=255,min=3"`
}

type Meta struct {
	Name        string `json:"name" bson:"name" validate:"required,max=255,min=3"`
	Description string `json:"description" bson:"description" validate:"required,max=255,min=5"`
	Title       string `json:"title" bson:"title" validate:"required,max=100,min=5"`
	Slug        string `json:"slug" bson:"slug"`
	Seo         Seo    `json:"seo" bson:"seo"  validate:"required"`
}

type Seo struct {
	Title       string     `json:"title" bson:"title" validate:"required,max=100,min=5"`
	Description string     `json:"description" bson:"description" validate:"required,max=255,min=5"`
	Keywords    string     `json:"keywords" bson:"keywords" validate:"required,max=255,min=5"`
	Canonical   string     `json:"canonical" bson:"canonical" validate:"omitempty,url"`
	Extra       []SeoExtra `json:"extra" bson:"extra" validate:"required"`
}

type SeoExtra struct {
	Name       string         `json:"name" bson:"name" validate:"required,max=255,min=3"`
	Content    string         `json:"content" bson:"content" validate:"required,max=255,min=3"`
	Attributes []SeoAttribute `json:"attributes" bson:"attributes" validate:"required"`
}

type SeoAttribute struct {
	Name  string `json:"name" bson:"name" validate:"required,max=255,min=3"`
	Value string `json:"value" bson:"value" validate:"required,max=255,min=3"`
}

type InputType string

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

const (
	InputTypeText     InputType = "text"
	InputTypeTextarea InputType = "textarea"
	InputTypeNumber   InputType = "number"
	InputTypeSelect   InputType = "select"
	InputTypeRadio    InputType = "radio"
	InputTypeCheckbox InputType = "checkbox"
	InputTypeDate     InputType = "date"
	InputTypeTime     InputType = "time"
	InputTypeDatetime InputType = "datetime"
	InputTypeFile     InputType = "file"
	InputTypeImage    InputType = "image"
	InputTypePDF      InputType = "pdf"
	InputTypeRange    InputType = "range"
	InputTypeColor    InputType = "color"
	InputTypeURL      InputType = "url"
	InputTypeEmail    InputType = "email"
	InputTypeTel      InputType = "tel"
	InputTypeLocation InputType = "location"
	InputTypePrice    InputType = "price"
	InputTypeRating   InputType = "rating"
)

func (e *Entity) IsMain() bool {
	return len(e.MainUUIDs) == 0
}

func (e *Entity) HasValidator(name string) bool {
	for _, v := range e.Validators {
		if v == name {
			return true
		}
	}
	return false
}

func (t InputType) String() string {
	return string(t)
}

func (t InputType) In(types ...InputType) bool {
	for _, v := range types {
		if v == t {
			return true
		}
	}
	return false
}

func (i Input) HasOption(option string) bool {
	for _, v := range i.Options {
		if v == option {
			return true
		}
	}
	return false
}

func (i Input) GetExtra(name string) (string, bool) {
	for _, v := range i.Extra {
		if v.Name == name {
			return v.Value, true
		}
	}
	return "", false
}

func (e Entity) GetInput(uuid string) (Input, bool) {
	for _, v := range e.Inputs {
		if v.UUID == uuid {
			return v, true
		}
	}
	return Input{}, false
}

func (l Locale) String() string {
	return string(l)
}

func (e *Entity) BeforeCreate() {
	e.IsActive = true
	e.IsDeleted = false
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
	e.Meta[LocaleEN].Slug = slug.New(e.Meta[LocaleEN].Title, slug.EN)
	e.Meta[LocaleTR].Slug = slug.New(e.Meta[LocaleTR].Title, slug.TR)
}

func (e *Entity) BeforeUpdate() {
	e.UpdatedAt = time.Now()
	e.Meta[LocaleEN].Slug = slug.New(e.Meta[LocaleEN].Title, slug.EN)
	e.Meta[LocaleTR].Slug = slug.New(e.Meta[LocaleTR].Title, slug.TR)
}
