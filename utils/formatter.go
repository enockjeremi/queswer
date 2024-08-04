package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var globalLang = "en"

type ErrorFormatter struct{}
type ErrorMessages struct {
	Required string
	Len      string
	Min      string
	Max      string
	Eq       string
	Ne       string
	Gt       string
	Gte      string
	Lt       string
	Lte      string
	Oneof    string
	Email    string
	URL      string
	UUID     string
	UUID4    string
	Alpha    string
	Alphanum string
	Numeric  string
	Contains string
	Default  string
}

func (*ErrorFormatter) LengFormatter(lang string) {
	globalLang = lang
}

func NewErrorFormatter() *ErrorFormatter {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	return &ErrorFormatter{}
}

func (*ErrorFormatter) Formatter(err error) map[string]string {
	var verr validator.ValidationErrors
	errs := make(map[string]string)
	if errors.As(err, &verr) {

		for _, f := range verr {
			errs[f.Field()] = NewErrorFormatter().GetErrorMessage(f.ActualTag(), f.Field(), f.Param())
		}
	}
	return errs
}

var errorMessages = map[string]ErrorMessages{
	"en": {
		Required: "{field} is required.",
		Len:      "{field} must be {param} characters.",
		Min:      "{field} must be at min {param} characters.",
		Max:      "{field} must be at max {param} characters.",
		Eq:       "{field} must be {param}.",
		Ne:       "{field} must not be {param}.",
		Gt:       "{field} must be greater than {param}.",
		Gte:      "{field} must be at least {param}.",
		Lt:       "{field} must be less than {param}.",
		Lte:      "{field} must be at most {param}.",
		Oneof:    "{field} must be one of {param}.",
		Email:    "{field} must be a valid email.",
		URL:      "{field} must be a valid URL.",
		UUID:     "{field} must be a valid UUID.",
		UUID4:    "{field} must be a valid UUIDv4.",
		Alpha:    "{field} must contain only letters.",
		Alphanum: "{field} must contain only letters and numbers.",
		Numeric:  "{field} must be a number.",
		Contains: "{field} must contain '{param}'.",
		Default:  "The {field} field is invalid.",
	},
	"es": {
		Required: "{field} es obligatorio.",
		Len:      "{field} debe tener {param} caracteres.",
		Min:      "{field} debe tener minimo {param} caracteres.",
		Max:      "{field} debe tener máximo {param} caracteres.",
		Eq:       "{field} debe ser igual {param}.",
		Ne:       "{field} no debe ser igual {param}.",
		Gt:       "{field} debe ser mayor que {param}.",
		Gte:      "{field} debe ser mayor o igual que {param}.",
		Lt:       "{field} debe ser menor que {param}.",
		Lte:      "{field} debe ser menor o igual que {param}.",
		Oneof:    "{field} debe ser uno de {param}.",
		Email:    "{field} debe ser un correo electrónico válido.",
		URL:      "{field} debe ser una URL válida.",
		UUID:     "{field} debe ser un UUID válido.",
		UUID4:    "{field} debe ser un UUIDv4 válido.",
		Alpha:    "{field} solo debe contener letras.",
		Alphanum: "{field} solo debe contener letras y números.",
		Numeric:  "{field} debe ser un número.",
		Contains: "{field} debe contener '{param}'.",
		Default:  "{field} campo invalido.",
	},
}

func (*ErrorFormatter) GetErrorMessage(tag, field, param string) string {
	messages := errorMessages[globalLang]
	var message string
	switch tag {
	case "required":
		message = messages.Required
	case "len":
		message = messages.Len
	case "min":
		message = messages.Min
	case "max":
		message = messages.Max
	case "eq":
		message = messages.Eq
	case "ne":
		message = messages.Ne
	case "gt":
		message = messages.Gt
	case "gte":
		message = messages.Gte
	case "lt":
		message = messages.Lt
	case "lte":
		message = messages.Lte
	case "oneof":
		message = messages.Oneof
	case "email":
		message = messages.Email
	case "url":
		message = messages.URL
	case "uuid":
		message = messages.UUID
	case "uuid4":
		message = messages.UUID4
	case "alpha":
		message = messages.Alpha
	case "alphanum":
		message = messages.Alphanum
	case "numeric":
		message = messages.Numeric
	case "contains":
		message = messages.Contains
	default:
		message = messages.Default
	}
	message = strings.ReplaceAll(message, "{field}", field)
	message = strings.ReplaceAll(message, "{param}", param)
	return message
}
