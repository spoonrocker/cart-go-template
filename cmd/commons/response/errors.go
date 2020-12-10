package response

import "net/http"

func NotFound(entity string, err error, w http.ResponseWriter) {
	http.Error(w, entity+" not found", http.StatusNotFound)
}

func FieldNotValid(entity string, w http.ResponseWriter) {
	http.Error(w, entity+" is invalid", http.StatusBadRequest)
}
