package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type Validator interface {
	Required() (string, bool)
	String() (string, bool)
	MaxChars() (string, bool)
	MinChars() (string, bool)
	URL() (string, bool)
	Kanji() (string, bool)
	Kana() (string, bool)
	Unique() (string, bool)
}

// "input" cannot be empty.
func Required(input string) (string, bool) {

	if input == "" {
		str := "Cannot be empty."
		return str, false
	}

	return "", true
}

func String(input interface{}) (string, bool) {
	if _, ok := input.(string); !ok {

		str := "Has to be of type string."
		return str, false
	}

	return "", true
}

func Int() {

}

// "input" can have "max" amount of characters.
func MaxChars(input string, max int) (string, bool) {
	if utf8.RuneCountInString(input) > max {

		str := fmt.Sprintf("Maximum %s characters.", strconv.Itoa(max))

		return str, false
	}

	return "", true
}

// "input" must have "min" amount of characters.
func MinChars(min int, input string) (string, bool) {
	if utf8.RuneCountInString(input) < min {

		str := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(min))

		return str, false
	}

	return "", true
}

func URL(s string) (string, bool) {

	m, err := regexp.MatchString("^(?:[http|https]+:\\/\\/)?(?:www\\.)?.+\\.[a-z]{2,3}$", s)

	if err != nil || !m {
		return "This is not a valid URL.", false
	}

	return "", true
}

func Kanji(s string) (string, bool) {

	match, _ := regexp.MatchString("^\\p{Han}+$", s)

	if !match {

		return "Contains invalid characters.", false
	}

	return "", true
}

func Kana(s string) (string, bool) {

	match, _ := regexp.MatchString("^[\\p{Katakana}\\p{Hiragana}]+$", s)

	if !match {

		return "Contains invalid characters.", false
	}

	return "", true
}

// func (model *models.Model) Unique(model instance, field string, value string) (string, bool) {

// 	_, collection := database.GetCollection("kanji")

// 	err := collection.Find(bson.D{{field, value}}).One(&model)
// 	if err == nil {
// 		return "Duplicate", false
// 	}

// 	return "", true
// }
