package models

import (
	_ "fmt"
	"log"

	"app/database"

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

func (k Kanji) Migrate() {

	session, collection := database.GetCollection("kanji")
	defer session.Close()

	database.RemoveCollection("kanji")

	index := mgo.Index{
		Key:    []string{"writing"},
		Unique: true,
	}

	if err := collection.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}
}

func (k Kanji) Seed() {

	session, collection := database.GetCollection("kanji")
	defer session.Close()

	err := collection.Insert(
		bson.M{"writing": "日", "reading": "にち", "meaning": "day"},
		bson.M{"writing": "木", "reading": "き", "meaning": "tree"},
		bson.M{"writing": "目", "reading": "め", "meaning": "eye"},
	)

	if err != nil {
		log.Fatal(err)
	}
}
