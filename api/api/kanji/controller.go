package kanji

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"app/database"
	"app/models"
	"app/utils/pagination"
	"app/utils/response"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	log.Println("innit?")
}

// response-code: 204, response-body: empty
func Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	err := collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {

		response.NotFound(w)
		return
	}

	response.NoContent(w)
}

func Get(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := vars["id"]

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	kanji := models.Kanji{}

	collection.FindId(bson.ObjectIdHex(id)).One(&kanji)

	response.Ok(w, kanji)
}

func GetAll(w http.ResponseWriter, r *http.Request) {

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	kanji := []models.Kanji{}
	query := r.URL.Query()

	pagination.Build(collection, query, &kanji)

	response.Ok(w, pagination.Paginator)
}

func Post(w http.ResponseWriter, r *http.Request) {

	// Read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {

		response.InternalServerError(w)
		return
	}

	// Read kanji
	kanji := &models.Kanji{}
	err = json.Unmarshal(payload, kanji)
	if err != nil {

		response.InternalServerError(w)
		return
	}

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	// Insert new
	if err := collection.Insert(kanji); err != nil {

		response.NotFound(w)
		return
	}

	response.Ok(w, kanji)
}

func Put(w http.ResponseWriter, r *http.Request) {

	// Read body
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {

		response.InternalServerError(w)
		return
	}

	// Read kanji
	kanji := &models.Kanji{}
	err = json.Unmarshal(payload, kanji)
	if err != nil {

		response.InternalServerError(w)
		return
	}

	vars := mux.Vars(r)
	id, _ := vars["id"]

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	// Insert new
	if err := collection.UpdateId(bson.ObjectIdHex(id), kanji); err != nil {

		response.NotFound(w)
		return
	}

	response.Ok(w, kanji)
}
