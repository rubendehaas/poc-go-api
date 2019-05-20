package kanji

import (
	"reflect"
)

func rules() map[string][]string {
	return map[string][]string{
		"writing": {"required", "max_chars:1", "kanji"},
		"reading": {"required", "kana"},
		"meaning": {"required"},
	}
}

func validateProperty(name string, value interface{}) {

	rules := rules()
	fieldRules := rules[name]

	for _, rule := range fieldRules {
		ruleMappings(rule, value)
	}
}

func ruleMappings(ruleName string, value interface{}) {

	switch ruleName {
	case "required":
		Required(value.(string))
	default:

	}
}

func (p Payload) Validate() {

	for key, value := range getProperties(p) {
		validateProperty(key, value)
	}
}

func getProperties(p Payload) map[string]interface{} {

	val := reflect.ValueOf(p)

	properties := make(map[string]interface{})

	for i := 0; i < val.Type().NumField(); i++ {

		name := val.Type().Field(i).Tag.Get("json")
		value := val.Field(i).Interface()

		properties[name] = value

		// fmt.Println(val.Field(i).Interface())
	}

	return properties
}

// "input" cannot be empty.
func Required(input string) (string, bool) {

	if input == "" {
		str := "Cannot be empty."
		return str, false
	}

	return "", true
}

// func String(input interface{}) (string, bool) {
// 	if _, ok := input.(string); !ok {

// 		str := "Has to be of type string."
// 		return str, false
// 	}

// 	return "", true
// }

// func Int() {

// }

// // "input" can have "max" amount of characters.
// func MaxChars(input string, max int) (string, bool) {
// 	if utf8.RuneCountInString(input) > max {

// 		str := fmt.Sprintf("Maximum %s characters.", strconv.Itoa(max))

// 		return str, false
// 	}

// 	return "", true
// }

// // "input" must have "min" amount of characters.
// func MinChars(min int, input string) (string, bool) {
// 	if utf8.RuneCountInString(input) < min {

// 		str := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(min))

// 		return str, false
// 	}

// 	return "", true
// }

// func URL(s string) (string, bool) {

// 	m, err := regexp.MatchString("^(?:[http|https]+:\\/\\/)?(?:www\\.)?.+\\.[a-z]{2,3}$", s)

// 	if err != nil || !m {
// 		return "This is not a valid URL.", false
// 	}

// 	return "", true
// }

// func Kanji(s string) (string, bool) {

// 	match, _ := regexp.MatchString("^\\p{Han}+$", s)

// 	if !match {

// 		return "Contains invalid characters.", false
// 	}

// 	return "", true
// }

// func Kana(s string) (string, bool) {

// 	match, _ := regexp.MatchString("^[\\p{Katakana}\\p{Hiragana}]+$", s)

// 	if !match {

// 		return "Contains invalid characters.", false
// 	}

// 	return "", true
// }

// // func (model *models.Model) Unique(model instance, field string, value string) (string, bool) {

// // 	_, collection := database.GetCollection("kanji")

// // 	err := collection.Find(bson.D{{field, value}}).One(&model)
// // 	if err == nil {
// // 		return "Duplicate", false
// // 	}

// // 	return "", true
// // }
