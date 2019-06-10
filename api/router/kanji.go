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

	create := validateRequest(http.HandlerFunc(kanji.Post))
	delete := modelBinding(http.HandlerFunc(kanji.Delete))
	list := http.HandlerFunc(kanji.List)
	get := modelBinding(http.HandlerFunc(kanji.Get))
	update := validateRequest(modelBinding(http.HandlerFunc(kanji.Put)))

	p.router.Handle("/kanji", middleware.Authorize(create)).Methods("POST")
	p.router.Handle("/kanji/{id}", middleware.Authorize(delete)).Methods("DELETE")
	p.router.Handle("/kanji", middleware.Authorize(list)).Methods("GET")
	p.router.Handle("/kanji/{id}", middleware.Authorize(get)).Methods("GET")
	p.router.Handle("/kanji/{id}", middleware.Authorize(update)).Methods("PUT")
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
