package category

import "time"

type Entity struct {
	UUID        string          `json:"uuid" bson:"_id,omitempty"`
	MainUUID    string          `json:"mainUUID" bson:"main_uuid"`
	Images      []Image         `json:"images" bson:"images"`
	Meta        map[Locale]Meta `json:"meta" bson:"meta"`
	InputGroups []InputGroup    `json:"inputGroups" bson:"input_groups"`
	Inputs      []Input         `json:"inputs" bson:"inputs"`
	Rules       []Rule          `json:"rules" bson:"rules"`
	Alerts      []Alert         `json:"alerts" bson:"alerts"`
	Validators  []string        `json:"validators" bson:"validators"`
	Order       int             `json:"order" bson:"order"`
	IsActive    bool            `json:"isActive" bson:"is_active"`
	IsDeleted   bool            `json:"isDeleted" bson:"is_deleted"`
	CreatedAt   time.Time       `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updatedAt" bson:"updated_at"`
}

type Image struct {
	Url   string `json:"url" bson:"url"`
	Order int16  `json:"order" bson:"order"`
}

type BaseTranslation struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Rule struct {
	UUID         string                     `json:"uuid" bson:"uuid"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations"`
	StrictLevel  int16                      `json:"strictLevel" bson:"strict_level"`
}

type Alert struct {
	UUID         string                     `json:"uuid" bson:"uuid"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations"`
	Type         string                     `json:"type" bson:"type"` // info, warning, error
}

type InputGroup struct {
	UUID         string                     `json:"uuid" bson:"uuid"`
	Icon         string                     `json:"icon" bson:"icon"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations"`
}

type Input struct {
	UUID         string                      `json:"uuid" bson:"uuid"`
	GroupUUID    string                      `json:"groupUUID" bson:"group_uuid"`
	Type         InputType                   `json:"type" bson:"type"`
	Translations map[Locale]InputTranslation `json:"translations" bson:"translations"`
	IsRequired   bool                        `json:"isRequired" bson:"is_required"`
	IsMultiple   bool                        `json:"isMultiple" bson:"is_multiple"`
	IsUnique     bool                        `json:"isUnique" bson:"is_unique"`
	IsPayed      bool                        `json:"isPayed" bson:"is_payed"`
	Extra        []InputExtra                `json:"extra" bson:"extra"`
	Options      []string                    `json:"options" bson:"options"`
}

type InputExtra struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
}

type InputTranslation struct {
	Name        string `json:"name" bson:"name"`
	Placeholder string `json:"placeholder" bson:"placeholder"`
	Help        string `json:"help" bson:"help"`
}

type Meta struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Title       string `json:"title" bson:"title"`
	Slug        string `json:"slug" bson:"slug"`
	MarkdownURL string `json:"markdownURL" bson:"markdown_url"`
	Seo         Seo
}

type Seo struct {
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Keywords    string     `json:"keywords" bson:"keywords"`
	Canonical   string     `json:"canonical" bson:"canonical"`
	Extra       []SeoExtra `json:"extra" bson:"extra"`
}

type SeoExtra struct {
	Name       string         `json:"name" bson:"name"`
	Content    string         `json:"content" bson:"content"`
	Attributes []SeoAttribute `json:"attributes" bson:"attributes"`
}

type SeoAttribute struct {
	Name  string `json:"name" bson:"name"`
	Value string `json:"value" bson:"value"`
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
	return e.MainUUID == ""
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
