package router

import (
	"app/api/kanji"
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
	kanji.Post(w, r)
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
	kanji.Put(w, r)
}
