package kanji

import (
	"net/http"

	"app/database"
	"app/models"
	"app/utils/pagination"
	"app/utils/response"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// response-code: 204, response-body: empty
func Delete(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id, _ := vars["id"]

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	err := collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {

		response.NotFound(writer)
		return
	}

	response.NoContent(writer)
}

func Get(writer http.ResponseWriter, request *http.Request) {

	kanji := (context.Get(request, "kanji")).(*models.Kanji)

	response.Ok(writer, kanji)
}

func List(writer http.ResponseWriter, request *http.Request) {

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	paginator := pagination.New(request.URL.Query())
	paginator.Build(collection, &[]models.Kanji{})

	response.Ok(writer, paginator)
}

func Post(writer http.ResponseWriter, request *http.Request) {

	resource := (context.Get(request, "resource")).(*models.Kanji)

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	if err := collection.Insert(resource); err != nil {
		response.InternalServerError(writer)
		return
	}

	response.Ok(writer, resource)
}

func Put(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	id, _ := vars["id"]

	resource := (context.Get(request, "resource")).(*models.Kanji)

	session, collection := database.GetCollection(models.TableKanji)
	defer session.Close()

	if err := collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, resource); err != nil {
		response.NotFound(writer)
		return
	}

	response.Ok(writer, resource)
}
