package database

import (
	"log"

	"app/models"
)

func SeedData() {

	log.Println("Seeding mock data to MongoDB")

	session, collection := GetCollection("kanji")
	defer session.Close()

	RemoveCollection(models.TableKanji)

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
