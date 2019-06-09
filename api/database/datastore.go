package database

import (
	"log"
	"os"
	"time"

	"app/models"

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

	seedKanjiData()
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

// TODO: extract
func seedKanjiData() {

	log.Println("Seeding mock data to MongoDB")

	session, collection := GetCollection("kanji")
	defer session.Close()

	RemoveCollection("kanji")

	kanji := models.Kanji{}

	kanji.Migrate(collection)

	err := collection.Insert(
		bson.M{"writing": "日", "reading": "にち", "meaning": "day"},
		bson.M{"writing": "木", "reading": "き", "meaning": "tree"},
		bson.M{"writing": "目", "reading": "め", "meaning": "eye"},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mock data added successfully!")
}
