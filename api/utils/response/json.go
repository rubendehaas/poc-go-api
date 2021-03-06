package response

import (
	"encoding/json"
	"net/http"
)

// 200
func Ok(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// 204
func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// 403
func Forbidden(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(m)
}

// 404
func NotFound(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusNotFound)
}

// 422
func UnprocessableEntity(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(m)
}

// 500
func InternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}
