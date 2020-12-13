package tests

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/cabogabo/cart-api/cmd/validators"

	"github.com/cabogabo/cart-api/cmd/services"

	"github.com/cabogabo/cart-api/cmd/commons"
)

// For testing, just run the command  -> go test
// all the service functions are going to be tested

func TestCreateCart(t *testing.T) {
	os.Setenv("DATABASE", "tests.db")
	commons.Connect()
	cart := services.InsertCart()
	log.Println(cart)
}

func TestCreateCartItem(t *testing.T) {
	body := validators.CartItemDTO{
		Product:  "example",
		Quantity: 3,
	}

	// body := validators.CartItemDTO{
	// 	Product:  "",
	// 	Quantity: 3,
	// } // Force an error

	cartItemDTO, _ := json.Marshal(body)

	cartItem, err := services.InsertCartItem(1, []byte(cartItemDTO))

	log.Println(cartItem)
	log.Println(err)

	cartItem, err = services.InsertCartItem(1, []byte(cartItemDTO)) // Add two items for testing delete cart items
}

func TestDeleteCartItem(t *testing.T) {
	cartItem, err := services.DeleteCartItem(1, 2)
	// cartItem, err := services.DeleteCartItem(5, 1) // Force an error
	// cartItem, err := services.DeleteCartItem(1, 5) // Force an error

	log.Println(cartItem)
	log.Println(err)
}

func TestGetCart(t *testing.T) {
	cart, err := services.GetCart(1)
	// cart, err := services.GetCart(9999) // Force an error
	log.Println(cart)
	log.Println(err)
}
