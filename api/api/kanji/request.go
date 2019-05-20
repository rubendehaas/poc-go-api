package kanji

import (
	"app/models"
	"app/utils/validation"
	"encoding/json"
	"net/url"
)

type Payload struct {
	Writing string `json:"writing"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
}

var (
	errors url.Values
)

func normalize(payload *Payload) models.Kanji {

	kanji := models.Kanji{}

	kanji.Writing = payload.Writing
	kanji.Reading = payload.Reading
	kanji.Meaning = payload.Meaning

	return kanji
}

func unserialize(rb []byte) *Payload {

	p := &Payload{}

	if err := json.Unmarshal(rb, p); err != nil {
		return nil
	}

	return p
}

func validate(rb []byte) (*models.Kanji, url.Values) {

	errors = url.Values{}

	p := unserialize(rb)

	if p == nil {

		errors.Add("invalid_payload", "Your payload sucks dude.")

		return nil, errors
	}

	p.Validate()

	// validateWriting(p.Writing)
	// validateReading(p.Reading)
	// validateMeaning(p.Meaning)

	if len(errors) > 0 {

		return nil, errors
	}

	kanji := normalize(p)

	return &kanji, nil
}

func validateWriting(w string) {

	propertyName := "writing"

	if err, success := validation.Required(w); success != true {
		errors.Add(propertyName, err)
	}

	if err, success := validation.MaxChars(w, 1); success != true {
		errors.Add(propertyName, err)
	}

	if err, success := validation.Kanji(w); success != true {
		errors.Add(propertyName, err)
	}

	// if err, success := validation.Unique(models.Kanji{}, propertyName, w); success != true {
	// 	errors.Add(propertyName, err)
	// }
}

func validateReading(w string) {

	propertyName := "reading"

	if err, success := validation.Required(w); success != true {
		errors.Add(propertyName, err)
	}

	if err, success := validation.Kana(w); success != true {
		errors.Add(propertyName, err)
	}
}

func validateMeaning(w string) {

	propertyName := "meaning"

	if err, success := validation.Required(w); success != true {
		errors.Add(propertyName, err)
	}
}
