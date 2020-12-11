package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/cabogabo/cart-api/cmd/commons"
	"github.com/cabogabo/cart-api/cmd/commons/response"
	"github.com/cabogabo/cart-api/cmd/services"
	"github.com/gorilla/mux"
)

func CartItemHandler(router *mux.Router) {
	router.HandleFunc("/carts/{cartID}/items", createCartItem).Methods("POST")
	router.HandleFunc("/carts/{cartID}/items/{itemID}", deleteCartItem).Methods("DELETE")
}

func createCartItem(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	cartItem, errorMessage := services.InsertCartItem(cartID, reqBody)

	if errorMessage != (commons.ErrorMessage{}) {
		response.ResponseError(errorMessage, w)
		return
	}

	response.Created(cartItem, w)
}

func deleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])
	cartItemID, _ := strconv.Atoi(mux.Vars(r)["itemID"])

	_, errorMessage := services.DeleteCartItem(cartID, cartItemID)

	if errorMessage != (commons.ErrorMessage{}) {
		response.ResponseError(errorMessage, w)
		return
	}

	response.NoContent(w)
}
