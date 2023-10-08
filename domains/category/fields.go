package category

type fieldsType struct {
	UUID        string
	MainUUID    string
	MainUUIDs   string
	Images      string
	Meta        string
	Inputs      string
	InputGroups string
	Rules       string
	Alerts      string
	Validators  string
	Order       string
	IsActive    string
	IsDeleted   string
	CreatedAt   string
	UpdatedAt   string
}

type imageFieldsType struct {
	Url   string
	Order string
}

type baseTranslationFieldsType struct {
	Locale      string
	Name        string
	Description string
}

type ruleFieldsType struct {
	UUID         string
	Translations string
	StrictLevel  string
}

type alertFieldsType struct {
	UUID         string
	Translations string
	Type         string
}

type inputGroupFieldsType struct {
	UUID         string
	Icon         string
	Translations string
}

type inputFieldsType struct {
	UUID         string
	GroupUUID    string
	Type         string
	Translations string
	IsRequired   string
	IsMultiple   string
	IsUnique     string
	IsPayed      string
	Extra        string
	Options      string
}

type inputExtraFieldsType struct {
	Name  string
	Value string
}

type inputTranslationFieldsType struct {
	Locale      string
	Name        string
	Placeholder string
	Help        string
}

type metaFieldsType struct {
	Locale      string
	Name        string
	Description string
	Title       string
	Slug        string
	MarkdownURL string
	Seo         string
}

type seoFieldsType struct {
	Title       string
	Description string
	Keywords    string
	Canonical   string
	Extra       string
}

type seoExtraFieldsType struct {
	Name       string
	Content    string
	Attributes string
}

type seoAttributesFieldsType struct {
	Name  string
	Value string
}

var fields = fieldsType{
	UUID:        "_id",
	MainUUID:   "main_uuid",
	MainUUIDs:   "main_uuids",
	Images:      "images",
	Meta:        "meta",
	Inputs:      "inputs",
	InputGroups: "input_groups",
	Rules:       "rules",
	Alerts:      "alerts",
	Validators:  "validators",
	Order:       "order",
	IsActive:    "is_active",
	IsDeleted:   "is_deleted",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var imageFields = imageFieldsType{
	Url:   "url",
	Order: "order",
}

var baseTranslationFields = baseTranslationFieldsType{
	Locale:      "locale",
	Name:        "name",
	Description: "description",
}

var ruleFields = ruleFieldsType{
	UUID:         "uuid",
	Translations: "translations",
	StrictLevel:  "strict_level",
}

var alertFields = alertFieldsType{
	UUID:         "uuid",
	Translations: "translations",
	Type:         "type",
}

var inputGroupFields = inputGroupFieldsType{
	UUID:         "uuid",
	Icon:         "icon",
	Translations: "translations",
}

var inputFields = inputFieldsType{
	UUID:         "uuid",
	GroupUUID:    "group_uuid",
	Type:         "type",
	Translations: "translations",
	IsRequired:   "is_required",
	IsMultiple:   "is_multiple",
	IsUnique:     "is_unique",
	IsPayed:      "is_payed",
	Extra:        "extra",
	Options:      "options",
}

var inputExtraFields = inputExtraFieldsType{
	Name:  "name",
	Value: "value",
}

var inputTranslationFields = inputTranslationFieldsType{
	Locale:      "locale",
	Name:        "name",
	Placeholder: "placeholder",
	Help:        "help",
}

var metaFields = metaFieldsType{
	Locale:      "locale",
	Name:        "name",
	Description: "description",
	Title:       "title",
	Slug:        "slug",
	MarkdownURL: "markdown_url",
	Seo:         "seo",
}

var seoFields = seoFieldsType{
	Title:       "title",
	Description: "description",
	Keywords:    "keywords",
	Canonical:   "canonical",
	Extra:       "extra",
}

var seoExtraFields = seoExtraFieldsType{
	Name:       "name",
	Content:    "content",
	Attributes: "attributes",
}

var seoAttributesFields = seoAttributesFieldsType{
	Name:  "name",
	Value: "value",
}

func metaField(locale string, field string) string {
	return fields.Meta + "." + locale + "." + field
}
