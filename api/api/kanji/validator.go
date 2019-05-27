package kanji

import (
	"app/utils/validation"
)

type Payload struct {
	Writing string `json:"writing"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
}

func (payload Payload) Validate() map[string][]string {

	rules := validation.DataFormat{
		"writing": {"required", "max_chars:1", "kanji"},
		"reading": {"required", "kana"},
		"meaning": {"required"},
	}

	options := validation.Options{
		Rules:   rules,
		Payload: payload,
	}

	validator := validation.New(options)

	err := validator.Validate()

	if err != nil {
		return err
	}

	return nil
}
