package validation

import (
	"net/url"
	"reflect"
	"strings"
)

type (
	DataFormat map[string][]string

	Options struct {
		Rules   DataFormat
		Payload interface{}
	}

	Validator struct {
		Options Options
	}
)

var (
	validationErrors url.Values
)

func New(options Options) *Validator {
	return &Validator{options}
}

func (validator *Validator) Validate() url.Values {

	validationErrors = url.Values{}

	for key, value := range validator.getPayloadProperties() {
		validator.validateProperty(key, value)
	}

	return validationErrors
}

func (validator *Validator) getPayloadProperties() map[string]interface{} {

	concreteValues := reflect.ValueOf(validator.Options.Payload)

	properties := make(map[string]interface{})

	for i := 0; i < concreteValues.Type().NumField(); i++ {

		name := concreteValues.Type().Field(i).Tag.Get("json")
		value := concreteValues.Field(i).Interface()

		properties[name] = value
	}

	return properties
}

func (validator *Validator) validateProperty(name string, value interface{}) {

	fieldRules := validator.Options.Rules[name]

	for _, rule := range fieldRules {
		resolveValidationMethod(rule, name, value)
	}
}

func resolveValidationMethod(ruleName string, name string, value interface{}) {

	var params string

	if strings.Contains(ruleName, ":") {

		parts := strings.Split(ruleName, ":")

		ruleName = parts[0]
		params = parts[1]
	}

	switch ruleName {
	case "required":
		Required(name, value.(string))
	case "integer":
		Integer(name, value.(string))
	case "str_max":
		StrMax(name, value.(string), params)
	case "str_min":
		StrMin(name, value.(string), params)
	case "int_max":
		IntMax(name, value.(string), params)
	case "int_min":
		IntMin(name, value.(string), params)
	case "url":
		Url(name, value.(string))
	case "jp_kanji":
		JPKanji(name, value.(string))
	case "jp_hiragana":
		JPHiragana(name, value.(string))
	case "jp_katakana":
		JPKatakana(name, value.(string))
	case "jp_kana":
		JPKana(name, value.(string))
	case "jp_all":
		JPAll(name, value.(string))
	}
}

func addError(name string, message string) {

	propertyErrors := []string{}

	if value, ok := validationErrors[name]; ok {
		propertyErrors = value
	}

	propertyErrors = append(propertyErrors, message)

	validationErrors[name] = propertyErrors
}
