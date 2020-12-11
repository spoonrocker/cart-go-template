package response

import (
	"encoding/json"
	"net/http"

	"github.com/cabogabo/cart-api/cmd/commons"
)

func ResponseError(errorMessage commons.ErrorMessage, w http.ResponseWriter) {
	switch errorMessage.ErrorType {
	case "invalid_field":
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
	case "not_found":
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorMessage)
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage)
	}
}
