package listing

import (
	"time"

	"github.com/cilloparch/cillop/i18np"
)

type ValidationError struct {
	Field   string  `json:"field"`
	Message string  `json:"message"`
	Params  i18np.P `json:"params"`
}

type UserDetail struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
}

type Entity struct {
	UUID          string          `json:"uuid"`
	Owner         Owner           `json:"owner"`
	Images        []Image         `json:"images"`
	Meta          map[Locale]Meta `json:"meta"`
	CategoryUUIDs []string        `json:"categoryUUIDs"`
	Features      []Feature       `json:"features"`
	Prices        []Price         `json:"prices"`
	Location      Location        `json:"location"`
	Boosts        []Boost         `json:"boosts"`
	Validation    Validation      `json:"validation"`
	Type          Type            `json:"type"`
	Order         int             `json:"order"`
	IsActive      bool            `json:"isActive"`
	IsDeleted     bool            `json:"isDeleted"`
	IsValid       bool            `json:"isValid"`
	CreatedAt     time.Time       `json:"createdAt"`
	UpdatedAt     time.Time       `json:"updatedAt"`
}
type Owner struct {
	UUID     string `json:"uuid"`
	NickName string `json:"nickName"`
}

type Image struct {
	Url   string `json:"url"`
	Order int16  `json:"order"`
}

type Validation struct {
	MinAdult   *int  `json:"minAdult" bson:"min_adult" validate:"required,min=1,max=50,ltefield=MaxAdult"`
	MaxAdult   *int  `json:"maxAdult" bson:"max_adult" validate:"required,min=0,max=50,gtefield=MinAdult"`
	MinKid     *int  `json:"minKid" bson:"min_kid" validate:"required,min=0,max=50,ltefield=MaxKid"`
	MaxKid     *int  `json:"maxKid" bson:"max_kid" validate:"required,min=0,max=50,gtefield=MinKid"`
	MinBaby    *int  `json:"minBaby" bson:"min_baby" validate:"required,min=0,max=50,ltefield=MaxBaby"`
	MaxBaby    *int  `json:"maxBaby" bson:"max_baby" validate:"required,min=0,max=50,gtefield=MinBaby"`
	MinDate    *int  `json:"minDate" bson:"min_date" validate:"required,min=0,max=50,ltefield=MaxDate"`
	MaxDate    *int  `json:"maxDate" bson:"max_date" validate:"required,min=0,max=50,gtefield=MinDate"`
	OnlyFamily *bool `json:"onlyFamily" bson:"only_family" validate:"required"`
	NoPet      *bool `json:"noPet" bson:"no_pet" validate:"required"`
	NoSmoke    *bool `json:"noSmoke" bson:"no_smoke" validate:"required"`
	NoAlcohol  *bool `json:"noAlcohol" bson:"no_alcohol" validate:"required"`
}

type Meta struct {
	Locale      string `json:"locale"`
	Description string `json:"description"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	MarkdownURL string `json:"markdownURL"`
}

type Feature struct {
	CategoryInputUUID string      `json:"categoryInputUUID"`
	Value             interface{} `json:"value"`
	IsPayed           bool        `json:"isPayed"`
}

type Price struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Price     float64   `json:"price"`
}

type Location struct {
	Country     string    `json:"country"`
	City        string    `json:"city"`
	Street      string    `json:"street"`
	Address     string    `json:"address"`
	IsStrict    bool      `json:"isStrict"`
	Coordinates []float64 `json:"coordinates"`
}

type Boost struct {
	UUID      string    `json:"uuid"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type Type string

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleTR Locale = "tr"
)

const (
	TypeEstate     Type = "estate"
	TypeCar        Type = "car"
	TypeBoat       Type = "boat"
	TypeMotorcycle Type = "motorcycle"
	TypeOther      Type = "other"
)

func (t Type) String() string {
	return string(t)
}

func (e Entity) GetFeatureByCategoryInputUUID(uuid string) (Feature, int, bool) {
	for idx, f := range e.Features {
		if f.CategoryInputUUID == uuid {
			return f, idx, true
		}
	}
	return Feature{}, -1, false
}

func (l Locale) String() string {
	return string(l)
}
