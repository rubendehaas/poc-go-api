package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Provider struct {
	router *mux.Router
}

var (
	provider Provider
)

func Load() {

	provider = Provider{mux.NewRouter()}

	provider.RegisterKanji()
	provider.RegisterToken()

	http.ListenAndServe(":8080", provider.router)
}
