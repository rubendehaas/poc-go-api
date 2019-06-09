package router

import (
	"app/api/token"
	"net/http"
)

func (p *Provider) RegisterToken() {

	getHandler := http.HandlerFunc(getToken)

	p.router.Handle("/token", getHandler).Methods("GET")
}

func getToken(writer http.ResponseWriter, request *http.Request) {
	token.Get(writer, request)
}
