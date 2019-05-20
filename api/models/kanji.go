package models

import (
	_ "fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TableKanji string = "kanji"

type Kanji struct {
	Id      bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Writing string        `json:"writing" bson:"writing"`
	Reading string        `json:"reading" bson:"reading"`
	Meaning string        `json:"meaning" bson:"meaning"`
}

func (k Kanji) Migrate(collection *mgo.Collection) {

	index := mgo.Index{
		Key:    []string{"writing"},
		Unique: true,
	}

	if err := collection.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}
}
