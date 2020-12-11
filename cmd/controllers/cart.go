package controllers

import (
	"net/http"
	"strconv"

	"github.com/cabogabo/cart-api/cmd/commons"
	"github.com/cabogabo/cart-api/cmd/commons/response"
	"github.com/cabogabo/cart-api/cmd/services"
	"github.com/gorilla/mux"
)

func CartHandler(router *mux.Router) {
	router.HandleFunc("/carts", createCart).Methods("POST")
	router.HandleFunc("/carts/{cartID}", getOneCart).Methods("GET")
}

func createCart(w http.ResponseWriter, r *http.Request) {
	cart := services.InsertCart()
	response.Created(cart, w)
}

func getOneCart(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])

	cart, err := services.GetCart(cartID)

	if err != (commons.ErrorMessage{}) {
		response.ResponseError(err, w)
		return
	}

	response.Ok(cart, w)
}
