package models

import (
	_ "fmt"

	"gopkg.in/mgo.v2/bson"
)

const TableKanji string = "kanji"

type Kanji struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Writing string        `json:"writing" bson:"writing"`
	Reading string        `json:"reading" bson:"reading"`
	Meaning string        `json:"meaning" bson:"meaning"`
}
