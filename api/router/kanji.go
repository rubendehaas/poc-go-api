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
	p.router.HandleFunc("/kanji", requestMiddleware(resourceMiddleware(createKanji))).Methods("POST")
	p.router.HandleFunc("/kanji/{id}", deleteKanji).Methods("DELETE")
	p.router.HandleFunc("/kanji", getAllKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", getKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", updateKanji).Methods("PUT")
}

func requestMiddleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		
		// TODO: extract the RequestHandler
		rawResource, errs := kanji.RequestHandler(request)
		if errs != nil {
			response.UnprocessableEntity(writer, errs)
			return
		}
		
		context.Set(request, "rawResource", rawResource)
		
		next.ServeHTTP(writer, request)
	})
}


func resourceMiddleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		
		rawResource := context.Get(request, "rawResource")
		
		session, collection := database.GetCollection(models.TableKanji)
		defer session.Close()

		kanjiResource := models.Kanji{}

		err := collection.Find(bson.M{"writing": rawResource.Writing}).One(&kanjiResource)
		if err != mgo.ErrNotFound {
			response.UnprocessableEntity(writer, url.Values{"illegal_operation": []string{"Resource already exists."}})
			return
		}
		
		next.ServeHTTP(writer, request)
	})
}

func createKanji(w http.ResponseWriter, request *http.Request) {
	kanji.Post(writer, request, rawResource)
}

func deleteKanji(w http.ResponseWriter, request *http.Request) {
	kanji.Delete(writer, request)
}

func getKanji(w http.ResponseWriter, request *http.Request) {
	kanji.Get(writer, request)
}

func getAllKanji(writer http.ResponseWriter, request *http.Request) {
	kanji.GetAll(writer, request)
}

func updateKanji(writer http.ResponseWriter, request *http.Request) {

	k, errs := kanji.RequestHandler(request)
	if errs != nil {
		response.UnprocessableEntity(writer, errs)
		return
	}

	kanji.Put(writer, request, k)
}
