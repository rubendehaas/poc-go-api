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

		if request.Header["Token"] == nil {
			response.Forbidden(writer, "Not Authorized")
			return
		}

		tokenString := request.Header["Token"][0]

		jwtToken, tokenError := parseToken(tokenString)

		if tokenError != nil {
			response.Forbidden(writer, "Not Authorized: token invalid")
			return
		}

		if expiredError := isTokenExpired(jwtToken); expiredError != nil {
			response.Forbidden(writer, expiredError.Error())
			return
		}

		if existError := tokenExists(tokenString); existError != nil {
			response.NotFound(writer)
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func tokenExists(tokenString string) error {

	session, collection := database.GetCollection(models.TableToken)
	defer session.Close()

	token := models.Token{}

	err := collection.Find(bson.M{"token": tokenString}).One(&token)
	if err == mgo.ErrNotFound {
		return fmt.Errorf("Not Authorized: token does not exist")
	}

	return nil
}

func isTokenExpired(jwtToken *jwt.Token) error {

	claims := jwtToken.Claims.(jwt.MapClaims)

	if time.Now().Unix() > int64(claims["expired_at"].(float64)) {
		return fmt.Errorf("Not Authorized: token expired")
	}

	return nil
}

func parseToken(tokenString string) (*jwt.Token, error) {

	var jwtToken *jwt.Token

	rawToken, tokenError := jwt.Parse(tokenString, func(rawToken *jwt.Token) (interface{}, error) {

		if _, ok := rawToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}

		jwtToken = rawToken

		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if tokenError != nil {
		return nil, fmt.Errorf(tokenError.Error())
	}

	if !rawToken.Valid {
		return nil, fmt.Errorf("There was an error")
	}

	return jwtToken, nil
}
