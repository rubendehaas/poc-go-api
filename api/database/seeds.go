package database

import (
	"app/models"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func Seed() {
	kanjiData()
}

func kanjiData() {

	session, collection := GetCollection("kanji")
	defer session.Close()

	RemoveCollection("kanji")

	(models.Kanji{}).Migrate(collection)

	err := collection.Insert(
		bson.M{"writing": "日", "reading": "にち", "meaning": "day"},
		bson.M{"writing": "木", "reading": "き", "meaning": "tree"},
		bson.M{"writing": "目", "reading": "め", "meaning": "eye"},
	)

	if err != nil {
		log.Fatal(err)
	}
}
