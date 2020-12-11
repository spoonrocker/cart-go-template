package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cabogabo/cart-api/cmd/controllers"

	"github.com/cabogabo/cart-api/cmd/commons"

	"github.com/gorilla/mux"
)

func main() {
	commons.Connect()
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	controllers.CartHandler(router)
	controllers.CartItemHandler(router)
	port := os.Getenv("PORT")
	log.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
