package router

import (
	"net/http"
	"os"

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

	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), provider.router)
}
