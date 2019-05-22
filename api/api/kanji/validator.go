package kanji

import (
	"app/utils/validation"
)

type Payload struct {
	Writing string `json:"writing"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
}

func rules() map[string][]string {
	return map[string][]string{
		"writing": {"required", "max_chars:1", "kanji"},
		"reading": {"required", "kana"},
		"meaning": {"required"},
	}
}

func (payload Payload) Validate() map[string][]string {

	err := validation.Validate(payload, rules())

	if err != nil {
		return err
	}

	return nil
}
