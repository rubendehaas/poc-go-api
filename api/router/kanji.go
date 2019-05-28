package router

import (
	"app/api/kanji"
	"app/utils/response"
	"net/http"
)

func (p *Provider) RegisterKanji() {

	p.router.HandleFunc("/kanji", createKanji).Methods("POST")
	p.router.HandleFunc("/kanji/{id}", deleteKanji).Methods("DELETE")
	p.router.HandleFunc("/kanji", getAllKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", getKanji).Methods("GET")
	p.router.HandleFunc("/kanji/{id}", updateKanji).Methods("PUT")
}

func createKanji(w http.ResponseWriter, r *http.Request) {

	k, errs := kanji.RequestHandler(r)
	if errs != nil {
		response.UnprocessableEntity(w, errs)
	}

	kanji.Post(w, r, k)
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
	}

	kanji.Put(w, r, k)
}
