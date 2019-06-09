package token

import (
	"fmt"
	"net/http"
	"time"

	"app/database"
	"app/models"
	"app/utils/response"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func Get(writer http.ResponseWriter, request *http.Request) {

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["expired_at"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := jwtToken.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return
	}

	token := &models.Token{
		Token: tokenString,
	}

	session, collection := database.GetCollection(models.TableToken)
	defer session.Close()

	if err := collection.Insert(token); err != nil {
		response.InternalServerError(writer)
		return
	}

	response.Ok(writer, token)
}
