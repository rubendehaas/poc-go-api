package models

import (
	_ "fmt"

	"gopkg.in/mgo.v2/bson"
)

const TableToken string = "token"

type Token struct {
	Id    bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Token string        `json:"token" bson:"token"`
}
