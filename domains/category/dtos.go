package category

import "time"

type ListDto struct {
	UUID      string                  `json:"uuid" bson:"_id,omitempty"`
	MainUUIDs []string                `json:"mainUUIDs" bson:"main_uuids"`
	Images    []Image                 `json:"images" bson:"images"`
	Meta      map[Locale]*MetaListDto `json:"meta" bson:"meta"`
}

type DetailDto struct {
	UUID      string           `json:"uuid" bson:"_id,omitempty"`
	MainUUIDs []string         `json:"mainUUIDs" bson:"main_uuids"`
	Images    []Image          `json:"images" bson:"images"`
	Meta      map[Locale]*Meta `json:"meta" bson:"meta"`
	CreatedAt time.Time        `json:"createdAt" bson:"created_at"`
}

type AdminDetailDto struct {
	*Entity
}

type InputGroupDto struct {
	UUID         string                     `json:"uuid" bson:"uuid" validate:"required,uuid4"`
	Icon         string                     `json:"icon" bson:"icon" validate:"required,max=255,min=3"`
	Translations map[Locale]BaseTranslation `json:"translations" bson:"translations" validate:"required,dive"`
	Inputs       []Input                    `json:"inputs" bson:"inputs" validate:"required,dive"`
}

type AdminListDto struct {
	UUID      string                  `json:"uuid" bson:"_id,omitempty"`
	MainUUIDs []string                `json:"mainUUIDs" bson:"main_uuids"`
	Images    []Image                 `json:"images" bson:"images"`
	Meta      map[Locale]*MetaListDto `json:"meta" bson:"meta"`
	IsActive  bool                    `json:"isActive" bson:"is_active"`
	IsDeleted bool                    `json:"isDeleted" bson:"is_deleted"`
	UpdatedAt time.Time               `json:"updatedAt" bson:"updated_at"`
}

type MetaListDto struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Title       string `json:"title" bson:"title"`
	Slug        string `json:"slug" bson:"slug"`
}

func (e *Entity) ToList() *ListDto {
	return &ListDto{
		UUID:      e.UUID,
		MainUUIDs: e.MainUUIDs,
		Images:    e.Images,
		Meta:      e.ToMetaList(),
	}
}

func (e *Entity) ToDetail() *DetailDto {
	return &DetailDto{
		UUID:      e.UUID,
		MainUUIDs: e.MainUUIDs,
		Images:    e.Images,
		Meta:      e.Meta,
		CreatedAt: e.CreatedAt,
	}
}

func (e *Entity) ToAdminDetail() *AdminDetailDto {
	return &AdminDetailDto{
		Entity: e,
	}
}

func (e *Entity) ToAdminList() *AdminListDto {
	return &AdminListDto{
		UUID:      e.UUID,
		MainUUIDs: e.MainUUIDs,
		Images:    e.Images,
		Meta:      e.ToMetaList(),
		IsActive:  e.IsActive,
		IsDeleted: e.IsDeleted,
		UpdatedAt: e.UpdatedAt,
	}
}

func (e *Entity) ToInputGroup() []*InputGroupDto {
	groupInputs := map[string][]Input{}
	for _, input := range e.Inputs {
		groupInputs[input.GroupUUID] = append(groupInputs[input.GroupUUID], input)
	}
	res := []*InputGroupDto{}
	for _, group := range e.InputGroups {
		res = append(res, &InputGroupDto{
			UUID:         group.UUID,
			Icon:         group.Icon,
			Translations: group.Translations,
			Inputs:       groupInputs[group.UUID],
		})
	}
	return res
}

func (m *Entity) ToMetaList() map[Locale]*MetaListDto {
	res := map[Locale]*MetaListDto{}
	for k, v := range m.Meta {
		res[k] = v.ToList()
	}
	return res
}

func (m *Meta) ToList() *MetaListDto {
	return &MetaListDto{
		Name:        m.Name,
		Description: m.Description,
		Title:       m.Title,
		Slug:        m.Slug,
	}
}
