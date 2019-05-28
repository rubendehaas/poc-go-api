package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// "input" cannot be empty.
func Required(name string, input string) {

	if input == "" {

		addError(name, "Cannot be empty.")
	}
}

func Integer(name string, input string) {

	_, err := strconv.Atoi(input)

	if err != nil {
		addError(name, "Not a number.")
	}
}

// "input" can have "max" amount of characters.
func StrMax(name string, input string, param string) {

	max, _ := strconv.Atoi(param)

	if utf8.RuneCountInString(input) > max {

		output := fmt.Sprintf("Maximum %s characters.", strconv.Itoa(max))

		addError(name, output)
	}
}

// "input" must have "min" amount of characters.
func StrMin(name string, input string, param string) {

	min, _ := strconv.Atoi(param)

	if utf8.RuneCountInString(input) < min {

		output := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(min))

		addError(name, output)
	}
}

// "input" must have "min" amount of characters.
func IntMax(name string, input string, param string) {

	value, _ := strconv.Atoi(input)
	max, _ := strconv.Atoi(param)

	if value > max {

		output := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(max))

		addError(name, output)
	}
}

// "input" must have "min" amount of characters.
func IntMin(name string, input string, param string) {

	value, _ := strconv.Atoi(input)
	min, _ := strconv.Atoi(param)

	if value < min {

		output := fmt.Sprintf("Minimum %s characters.", strconv.Itoa(min))

		addError(name, output)
	}
}

func Url(name string, input string) {

	match, _ := regexp.MatchString("^(?:[http|https]+:\\/\\/)?(?:www\\.)?.+\\.[a-z]{2,3}$", input)

	if !match {
		addError(name, "This is not a valid URL.")
	}
}

func JPKanji(name string, input string) {

	match, _ := regexp.MatchString("^\\p{Han}+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}

func JPHiragana(name string, input string) {

	match, _ := regexp.MatchString("^\\p{Hiragana}+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}

func JPKatakana(name string, input string) {

	match, _ := regexp.MatchString("^\\p{Katakana}+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}

func JPKana(name string, input string) {

	match, _ := regexp.MatchString("^[\\p{Katakana}\\p{Hiragana}]+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}

func JPAll(name string, input string) {

	match, _ := regexp.MatchString("^[\\p{Han}\\p{Katakana}\\p{Hiragana}]+$", input)

	if !match {
		addError(name, "Contains invalid characters.")
	}
}
