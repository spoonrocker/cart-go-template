package cart

import (
	"github.com/gorilla/mux"
)

func HandleRequests(router *mux.Router) {
	router.HandleFunc("/carts", createCart).Methods("POST")
	router.HandleFunc("/carts/{cartID}", fetchCart).Methods("GET")
	router.HandleFunc("/carts/{cartID}/items", createCartItem).Methods("POST")
	router.HandleFunc("/carts/{cartID}/items/{itemID}", deleteCartItem).Methods("DELETE")
}
