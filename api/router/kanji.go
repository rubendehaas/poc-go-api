package router

import (
	"app/api/kanji"
	"app/database"
	"app/models"
	"app/utils/response"
	"net/http"
	"net/url"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (p *Provider) RegisterKanji() {

	p.router.HandleFunc("/kanji", createKanji).Methods("POST")
	p.router.HandleFunc("/kanji/{id}", deleteKanji).Methods("DELETE")
	p.router.HandleFunc("/kanji", getAllKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", getKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", updateKanji).Methods("PUT")
}

func createKanji(w http.ResponseWriter, r *http.Request) {

	rawResource, errs := kanji.RequestHandler(r)
	if errs != nil {
		response.UnprocessableEntity(w, errs)
		return
	}

	// TODO: db request validation
	// Kanji is validated by its unique Writing or ID

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	kanjiResource := models.Kanji{}

	err := collection.Find(bson.M{"writing": rawResource.Writing}).One(&kanjiResource)
	if err != mgo.ErrNotFound {
		response.UnprocessableEntity(w, url.Values{"illegal_operation": []string{"Resource already exists."}})
		return
	}

	kanji.Post(w, r, rawResource)
}

func deleteKanji(w http.ResponseWriter, r *http.Request) {
	kanji.Delete(w, r)
}

func getKanji(w http.ResponseWriter, r *http.Request) {
	kanji.Get(w, r)
}

func getAllKanji(w http.ResponseWriter, r *http.Request) {
	kanji.GetAll(w, r)
}

func updateKanji(w http.ResponseWriter, r *http.Request) {

	k, errs := kanji.RequestHandler(r)
	if errs != nil {
		response.UnprocessableEntity(w, errs)
		return
	}

	kanji.Put(w, r, k)
}
