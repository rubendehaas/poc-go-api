package models

import (
	"app/database"
	_ "fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TableToken string = "token"

type Token struct {
	Id    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Token string        `json:"token" bson:"token"`
}

func (t Token) Migrate() {

	session, collection := database.GetCollection("kanji")
	defer session.Close()

	database.RemoveCollection("kanji")

	index := mgo.Index{
		Key:    []string{"token"},
		Unique: true,
	}

	if err := collection.EnsureIndex(index); err != nil {
		log.Fatal(err)
	}
}
