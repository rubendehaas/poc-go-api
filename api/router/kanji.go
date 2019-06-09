package router

import (
	"app/api/kanji"
	"app/database"
	"app/models"
	"app/utils/response"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func (p *Provider) RegisterKanji() {

	createHandler := http.HandlerFunc(createKanji)
	deleteHandler := http.HandlerFunc(deleteKanji)
	getHandler := http.HandlerFunc(getKanji)
	getAllHandler := http.HandlerFunc(getAllKanji)
	updateHandler := http.HandlerFunc(updateKanji)

	p.router.Handle("/kanji", isAuthorized(requestMiddleware(createHandler))).Methods("POST")
	p.router.Handle("/kanji/{id}", isAuthorized(modelBindingMiddleware(deleteHandler))).Methods("DELETE")
	p.router.Handle("/kanji", isAuthorized(getAllHandler)).Methods("GET")
	p.router.Handle("/kanji/{id}", isAuthorized(modelBindingMiddleware(getHandler))).Methods("GET")
	p.router.Handle("/kanji/{id}", isAuthorized(requestMiddleware(modelBindingMiddleware(updateHandler)))).Methods("PUT")
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

func isAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] != nil {

			var jwtToken *jwt.Token

			tokenString := request.Header["Token"][0]

			rawToken, errToken := jwt.Parse(tokenString, func(rawToken *jwt.Token) (interface{}, error) {

				_, ok := rawToken.Method.(*jwt.SigningMethodHMAC)

				if !ok {
					return nil, fmt.Errorf("There was an error")
				}

				jwtToken = rawToken

				return mySigningKey, nil
			})

			if errToken != nil {
				fmt.Fprintf(writer, errToken.Error())
				return
			}

			claims := jwtToken.Claims.(jwt.MapClaims)

			fmt.Printf("%v %v /n", time.Now().Unix(), int64(claims["expired_at"].(float64)))

			if time.Now().Unix() > int64(claims["expired_at"].(float64)) {
				response.Forbidden(writer, "Not Authorized: token expired")
			}

			session, collection := database.GetCollection(models.TableToken)
			defer session.Close()

			token := models.Token{}

			err := collection.Find(bson.M{"token": tokenString}).One(&token)
			if err == mgo.ErrNotFound {
				response.NotFound(writer)
				return
			}

			if rawToken.Valid {
				next.ServeHTTP(writer, request)
			}
		}

		response.Forbidden(writer, "Not Authorized")
	})
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
