package router

import (
	"app/api/kanji"
	"app/database"
	"app/models"
	"app/router/middleware"
	"app/utils/response"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (p *Provider) RegisterKanji() {

	getAll := http.HandlerFunc(kanji.GetAll)

	p.router.HandleFunc("/kanji", createKanji).Methods("POST")
	p.router.HandleFunc("/kanji/{id}", deleteKanji).Methods("DELETE")
	p.router.Handle("/kanji", middleware.Authorize(getAll)).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", getKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", updateKanji).Methods("PUT")
}

func createKanji(writer http.ResponseWriter, request *http.Request) {

	handler := http.HandlerFunc(kanji.Post)

	middleware.Authorize(
		validateRequest(
			handler,
		),
	)

	handler.ServeHTTP(writer, request)
}

func deleteKanji(writer http.ResponseWriter, request *http.Request) {

	handler := http.HandlerFunc(kanji.Delete)

	middleware.Authorize(
		validateRequest(
			modelBinding(
				handler,
			),
		),
	)

	handler.ServeHTTP(writer, request)
}

func getKanji(writer http.ResponseWriter, request *http.Request) {

	handler := http.HandlerFunc(kanji.Get)

	middleware.Authorize(
		modelBinding(
			handler,
		),
	)

	handler.ServeHTTP(writer, request)
}

func updateKanji(writer http.ResponseWriter, request *http.Request) {

	handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		kanji.Put(writer, request)
	})

	middleware.Authorize(
		validateRequest(
			modelBinding(
				handler,
			),
		),
	)

	handler.ServeHTTP(writer, request)
}

func validateRequest(next http.Handler) http.Handler {

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

func modelBinding(next http.Handler) http.Handler {

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
