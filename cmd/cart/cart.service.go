package cart

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/cabogabo/cart-api/cmd/commons"
	"github.com/cabogabo/cart-api/cmd/commons/response"
	"github.com/gorilla/mux"
)

func createCart(w http.ResponseWriter, r *http.Request) {
	cartDTO := CartDTO{}
	commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("carts"))
		id, _ := b.NextSequence()
		cartDTO.ID = int(id)
		buf, err := json.Marshal(cartDTO)
		if err != nil {
			return err
		}

		return b.Put(commons.Itob(cartDTO.ID), buf)
	})

	cart, err := getCart(cartDTO.ID)

	if err != nil {
		response.NotFound("Cart", err, w)
		return
	}

	response.Created(cart, w)
}

func fetchCart(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])

	cart, err := getCart(cartID)

	if err != nil {
		response.NotFound("Cart", err, w)
		return
	}

	response.Ok(cart, w)
}

func createCartItem(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = getCart(cartID)

	if err != nil {
		response.NotFound("Cart", err, w)
		return
	}

	cartItemDTO := CartItemDTO{}
	json.Unmarshal(reqBody, &cartItemDTO)

	if cartItemDTO.Product == "" {
		err = fmt.Errorf("Invalid field")
		response.FieldNotValid("Product", w)
		return
	}

	if cartItemDTO.Quantity <= 0 {
		err = fmt.Errorf("Invalid field")
		response.FieldNotValid("Quantity", w)
		return
	}

	commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cart_items"))
		id, _ := b.NextSequence()
		cartItemDTO.ID = int(id)
		cartItemDTO.CartID = cartID
		buf, err := json.Marshal(cartItemDTO)
		if err != nil {
			return err
		}

		return b.Put(commons.Itob(cartItemDTO.ID), buf)
	})

	cartItem, err := getCartItem(cartItemDTO.ID)

	if err != nil {
		response.NotFound("Cart item", err, w)
		return
	}

	response.Created(cartItem, w)
}

func deleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartID, _ := strconv.Atoi(mux.Vars(r)["cartID"])
	cartItemID, _ := strconv.Atoi(mux.Vars(r)["itemID"])

	_, err := getCart(cartID)

	if err != nil {
		response.NotFound("Cart", err, w)
		return
	}

	_, err = getCartItem(cartItemID)

	if err != nil {
		response.NotFound("Cart item", err, w)
		return
	}

	commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cart_items"))
		err := b.Delete([]byte(commons.Itob(cartItemID)))
		return err
	})

	if err != nil {
		log.Fatalln(err)
	}

	response.NoContent(w)
}
