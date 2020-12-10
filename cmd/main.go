package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cabogabo/cart-api/cmd/commons"

	"github.com/cabogabo/cart-api/cmd/cart"
	"github.com/gorilla/mux"
)

func main() {
	commons.Connect()
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	cart.HandleRequests(router)
	port := os.Getenv("PORT")
	log.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
