package category

import (
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cilloparch/cillop/formats"
	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/validation"
	"github.com/turistikrota/service.category/config"
)

func newTextValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeText, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newTextareaValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeTextarea, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newNumberValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeNumber, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(int)
		if err {
			_, err = value.(float64)
			if err {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
		}
		return nil
	})
}

func newSelectValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeSelect, func(input Input, value interface{}) *i18np.Error {
		if *input.IsMultiple {
			val, err := value.([]string)
			if err {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
			for _, v := range val {
				if !input.HasOption(v) {
					return errors.InvalidCategoryInputType(input.Type.String(), value)
				}
			}
		} else {
			val, err := value.(string)
			if err {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
			if !input.HasOption(val) {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
		}
		return nil
	})
}

func newRadioValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeRadio, func(input Input, value interface{}) *i18np.Error {
		if *input.IsMultiple {
			val, err := value.([]string)
			if err {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
			for _, v := range val {
				if !input.HasOption(v) {
					return errors.InvalidCategoryInputType(input.Type.String(), value)
				}
			}
		} else {
			val, err := value.(string)
			if err {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
			if !input.HasOption(val) {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
		}
		return nil
	})
}

func newCheckboxValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeCheckbox, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(bool)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newDateValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeDate, func(input Input, value interface{}) *i18np.Error {
		str, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := time.Parse(formats.DateYYYYMMDD, str)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newTimeValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeTime, func(input Input, value interface{}) *i18np.Error {
		str, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := time.Parse(formats.HHMMSS, str)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newDateTimeValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeDatetime, func(input Input, value interface{}) *i18np.Error {
		str, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := time.Parse(formats.DateYYYYMMDDHHMMSS, str)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newFileValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeFile, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		isCdn := strings.Contains(value.(string), config.CDN.Url)
		if !isCdn {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		fileTypes := input.Options
		if len(fileTypes) == 0 {
			fileTypes = []string{"docx", "md", "pdf", "xlsx", "xls"}
		}
		isFileType := false
		for _, fileType := range fileTypes {
			if strings.Contains(value.(string), fileType) {
				isFileType = true
				break
			}
		}
		if !isFileType {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newImageValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeImage, func(input Input, value interface{}) *i18np.Error {
		_, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		isCdn := strings.Contains(value.(string), config.CDN.Url)
		if !isCdn {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		fileTypes := input.Options
		if len(fileTypes) == 0 {
			fileTypes = []string{"jpg", "jpeg", "png", "gif", "svg"}
		}
		isFileType := false
		for _, fileType := range fileTypes {
			if strings.Contains(value.(string), fileType) {
				isFileType = true
				break
			}
		}
		if !isFileType {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newPDFValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypePDF, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		isCdn := strings.Contains(val, config.CDN.Url)
		if !isCdn {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		fileTypes := input.Options
		if len(fileTypes) == 0 {
			fileTypes = []string{"pdf"}
		}
		isFileType := false
		for _, fileType := range fileTypes {
			if strings.Contains(val, fileType) {
				isFileType = true
				break
			}
		}
		if !isFileType {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newRangeValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeRange, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(float64)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		min, exist := input.GetExtra("min")
		if exist {
			min, _ := strconv.ParseFloat(min, 64)
			if val < min {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
		}
		max, exist := input.GetExtra("max")
		if exist {
			max, _ := strconv.ParseFloat(max, 64)
			if val > max {
				return errors.InvalidCategoryInputType(input.Type.String(), value)
			}
		}
		return nil
	})
}

func newColorValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeColor, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := strconv.ParseUint(val, 16, 64)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newURLValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeURL, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := url.Parse(val)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newEmailValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeEmail, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		_, _err := mail.ParseAddress(val)
		if _err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newPhoneValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeTel, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(string)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if _, err := regexp.MatchString(validation.PhoneWithCountryCodeRegexp, val); err != nil {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newLocationValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeLocation, func(input Input, value interface{}) *i18np.Error {
		val, err := value.([]float64)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if len(val) != 2 {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if val[0] < -180 || val[0] > 180 {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if val[1] < -90 || val[1] > 90 {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newPriceValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypePrice, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(float64)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if val < 0 {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

func newRatingValidator(errors Errors, config config.App) Validator {
	return newValidator(InputTypeRating, func(input Input, value interface{}) *i18np.Error {
		val, err := value.(float64)
		if err {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		if val < 0 || val > 5 {
			return errors.InvalidCategoryInputType(input.Type.String(), value)
		}
		return nil
	})
}

type Validators map[InputType]Validator

func NewValidators(errors Errors, config config.App) Validators {
	return Validators{
		InputTypeText:     newTextValidator(errors, config),
		InputTypeTextarea: newTextareaValidator(errors, config),
		InputTypeNumber:   newNumberValidator(errors, config),
		InputTypeSelect:   newSelectValidator(errors, config),
		InputTypeRadio:    newRadioValidator(errors, config),
		InputTypeCheckbox: newCheckboxValidator(errors, config),
		InputTypeDate:     newDateValidator(errors, config),
		InputTypeTime:     newTimeValidator(errors, config),
		InputTypeDatetime: newDateTimeValidator(errors, config),
		InputTypeFile:     newFileValidator(errors, config),
		InputTypeImage:    newImageValidator(errors, config),
		InputTypePDF:      newPDFValidator(errors, config),
		InputTypeRange:    newRangeValidator(errors, config),
		InputTypeColor:    newColorValidator(errors, config),
		InputTypeURL:      newURLValidator(errors, config),
		InputTypeEmail:    newEmailValidator(errors, config),
		InputTypeTel:      newPhoneValidator(errors, config),
		InputTypeLocation: newLocationValidator(errors, config),
		InputTypePrice:    newPriceValidator(errors, config),
		InputTypeRating:   newRatingValidator(errors, config),
	}
}

func (v Validators) GetValidator(t InputType) (Validator, bool) {
	val, ok := v[t]
	return val, ok
}
