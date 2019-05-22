package kanji

import (
	"app/models"
	"encoding/json"
	"net/url"
)

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

	payload := unserialize(rb)

	if payload == nil {

		errors.Add("invalid_payload", "Your payload sucks dude.")

		return nil, errors
	}

	errors := payload.Validate()

	if len(errors) > 0 {

		return nil, errors
	}

	kanji := normalize(payload)

	return &kanji, nil
}
