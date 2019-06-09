package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"app/database"
	"app/models"
	"app/utils/response"

	jwt "github.com/dgrijalva/jwt-go"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Authorize(next http.Handler) http.Handler {

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

				return []byte(os.Getenv("JWT_TOKEN")), nil
			})

			if errToken != nil {
				fmt.Fprintf(writer, errToken.Error())
				return
			}

			claims := jwtToken.Claims.(jwt.MapClaims)

			if time.Now().Unix() > int64(claims["expired_at"].(float64)) {
				response.Forbidden(writer, "Not Authorized: token expired")
				return
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
				return
			}
		}

		response.Forbidden(writer, "Not Authorized")
	})
}
