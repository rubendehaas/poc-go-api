package kanji

import (
	"app/models"
	"app/utils/validation"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Payload struct {
	Writing string `json:"writing"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
}

func normalize(payload *Payload) models.Kanji {

	kanji := models.Kanji{}

	kanji.Writing = payload.Writing
	kanji.Reading = payload.Reading
	kanji.Meaning = payload.Meaning

	return kanji
}

func unserialize(requestBody io.ReadCloser) *Payload {

	payload := &Payload{}
	err := json.NewDecoder(requestBody).Decode(payload)
	if err != nil {
		return nil
	}

	return payload
}

func (payload Payload) validate() map[string][]string {

	rules := validation.DataFormat{
		"writing": {"required", "str_max:1", "jp_kanji"},
		"reading": {"required", "jp_kana"},
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

func RequestHandler(request *http.Request) (*models.Kanji, url.Values) {

	errors := url.Values{}
	payload := unserialize(request.Body)

	if payload == nil {
		errors.Add("invalid_payload", "Your payload sucks dude.")
		return nil, errors
	}

	errors = payload.validate()

	if len(errors) > 0 {
		return nil, errors
	}

	kanji := normalize(payload)

	return &kanji, nil
}
