package response

import (
	"encoding/json"
	"net/http"
)

func Created(entity interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

func Ok(entity interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entity)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
