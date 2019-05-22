package validation

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var validationErrors map[string][]string

var rules map[string][]string

func Validate(payload interface{}, _rules map[string][]string) map[string][]string {

	defer reset()

	validationErrors = make(map[string][]string)

	rules = _rules

	for key, value := range getPayloadProperties(payload) {
		validateProperty(key, value)
	}

	return validationErrors
}

func reset() {
	rules = map[string][]string{}
	validationErrors = map[string][]string{}
}

func getPayloadProperties(payload interface{}) map[string]interface{} {

	concreteValues := reflect.ValueOf(payload)

	properties := make(map[string]interface{})

	for i := 0; i < concreteValues.Type().NumField(); i++ {

		name := concreteValues.Type().Field(i).Tag.Get("json")
		value := concreteValues.Field(i).Interface()

		properties[name] = value
	}

	return properties
}

func validateProperty(name string, value interface{}) {

	fieldRules := rules[name]

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
	case "max_chars":
		MaxChars(name, value.(string), params)
	case "min_chars":
		MinChars(name, value.(string), params)
	case "url":
		URL(name, value.(string))
	case "kanji":
		KanjiJP(name, value.(string))
	case "kana":
		KanaJP(name, value.(string))
	default:

	}
}

// "input" cannot be empty.
func Required(name string, input string) {

	if input == "" {

		addError(name, "Cannot be empty.")
	}
}

// "input" can have "max" amount of characters.
func MaxChars(name string, input string, param string) {

	max, _ := strconv.Atoi(param)

	if utf8.RuneCountInString(input) > max {

		output := fmt.Sprintf("Maximum %s characters.", strconv.Itoa(max))

		addError(name, output)
	}
}

// "input" must have "min" amount of characters.
func MinChars(name string, input string, param string) {

	min, _ := strconv.Atoi(param)

	if utf8.RuneCountInString(input) < min {

		output := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(min))

		addError(name, output)
	}
}

func URL(name string, input string) {

	match, _ := regexp.MatchString("^(?:[http|https]+:\\/\\/)?(?:www\\.)?.+\\.[a-z]{2,3}$", input)

	if !match {
		addError(name, "This is not a valid URL.")
	}
}

func KanjiJP(name string, input string) {

	match, _ := regexp.MatchString("^\\p{Han}+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}

func KanaJP(name string, input string) {

	match, _ := regexp.MatchString("^[\\p{Katakana}\\p{Hiragana}]+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
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
