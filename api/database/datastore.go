package database

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var host string
var databaseName string

type DataStore struct {
	session *mgo.Session
}

var (
	dataStore DataStore
)

func Load() {

	host = os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT")
	databaseName = os.Getenv("DATABASE_Name")

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

	collection := session.DB(databaseName).C(table)

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
