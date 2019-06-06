package router

import (
	"app/api/kanji"
	"app/database"
	"app/models"
	"app/utils/response"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (p *Provider) RegisterKanji() {

	createHandler := http.HandlerFunc(createKanji)
	deleteHandler := http.HandlerFunc(deleteKanji)
	getHandler := http.HandlerFunc(getKanji)
	updateHandler := http.HandlerFunc(updateKanji)

	p.router.Handle("/kanji", requestMiddleware(createHandler)).Methods("POST")
	p.router.Handle("/kanji/{id}", modelBindingMiddleware(deleteHandler)).Methods("DELETE")
	p.router.HandleFunc("/kanji", getAllKanji).Methods("GET")
	p.router.Handle("/kanji/{id}", modelBindingMiddleware(getHandler)).Methods("GET")
	p.router.Handle("/kanji/{id}", requestMiddleware(modelBindingMiddleware(updateHandler))).Methods("PUT")
}

func createKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.Post(writer, request)
}

func deleteKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.Delete(writer, request)
}

func getKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.Get(writer, request)
}

func getAllKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.GetAll(writer, request)
}

func updateKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.Put(writer, request)
}

func requestMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		resource, errs := kanji.RequestHandler(request)
		if errs != nil {
			response.UnprocessableEntity(writer, errs)
			return
		}

		context.Set(request, "resource", resource)

		next.ServeHTTP(writer, request)
	})
}

func modelBindingMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		vars := mux.Vars(request)
		id, _ := vars["id"]

		session, collection := database.GetCollection(models.TableKanji)
		defer session.Close()

		kanji := models.Kanji{}

		err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&kanji)
		if err == mgo.ErrNotFound {
			response.NotFound(writer)
			return
		}

		context.Set(request, "kanji", &kanji)

		next.ServeHTTP(writer, request)
	})
}
