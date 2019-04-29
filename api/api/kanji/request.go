package kanji

import (
	"app/models"
)

type Payload struct {
	Writing string `json:"writing"`
	Reading string `json:"reading"`
	Meaning string `json:"meaning"`
}

func MorphToModel(payload Payload) models.Kanji {

	kanji := models.Kanji{}

	kanji.Writing = payload.Writing
	kanji.Reading = payload.Reading
	kanji.Meaning = payload.Meaning

	return kanji
}
