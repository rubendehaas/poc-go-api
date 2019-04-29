package database

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const host string = "mongo:27017"
const database string = "go_japanese"

type DataStore struct {
	session *mgo.Session
}

var (
	dataStore DataStore
)

func Init() {

	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout(host, maxWait)

	if err != nil {
		log.Fatal(err)
	}

	dataStore = DataStore{session}
}

func NewSession() *mgo.Session {

	return dataStore.session.Copy()
}

func GetCollection(table string) (*mgo.Session, *mgo.Collection) {

	session := NewSession()

	collection := session.DB(database).C(table)

	return session, collection
}

func RemoveCollection(col string) {

	log.Println("Cleaning up MongoDB...")

	session, collection := GetCollection(col)

	defer session.Close()

	_, err := collection.RemoveAll(bson.M{})
	if err != nil {
		log.Fatal(err)
	}
}
