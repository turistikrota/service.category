package category

type messageTypes struct {
	InvalidInputType          string
	InvalidInputUUID          string
	InvalidInputGroupUUID     string
	Failed                    string
	NotFound                  string
	MinOneTranslationRequired string
	InvalidMetaLength         string
	InvalidImagesLength       string
	InvalidUUID               string
	InvalidCategoryInputType  string
	ParentUUIDIsNotEmpty      string
	ParentUUIDIsNotCorrect    string
	CategoryUUIDsIsNotCorrect string
	FeatureIsNotCorrect       string
}

var messages = messageTypes{
	InvalidInputType:          "category_invalid_input_type",
	InvalidInputUUID:          "category_invalid_input_uuid",
	InvalidInputGroupUUID:     "category_invalid_input_group_uuid",
	Failed:                    "category_failed",
	NotFound:                  "category_not_found",
	MinOneTranslationRequired: "category_min_one_translation_required",
	InvalidMetaLength:         "category_invalid_meta_length",
	InvalidImagesLength:       "category_invalid_images_length",
	InvalidUUID:               "category_invalid_uuid",
	InvalidCategoryInputType:  "category_invalid_category_input_type",
	ParentUUIDIsNotCorrect:    "category_parent_uuid_is_not_correct",
	ParentUUIDIsNotEmpty:      "category_parent_uuid_is_not_empty",
	CategoryUUIDsIsNotCorrect: "category_category_uuids_is_not_correct",
	FeatureIsNotCorrect:       "category_feature_is_not_correct",
}
