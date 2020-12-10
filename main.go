package main

import (
	"log"
	"net/http"

	"github.com/cabogabo/cart-api/commons"

	"github.com/cabogabo/cart-api/cart"
	"github.com/gorilla/mux"
)

func main() {
	commons.Connect()
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	cart.HandleRequests(router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
