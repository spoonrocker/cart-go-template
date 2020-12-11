package services

import (
	"encoding/json"

	"github.com/cabogabo/cart-api/cmd/models"
	"github.com/cabogabo/cart-api/cmd/validators"

	"github.com/cabogabo/cart-api/cmd/commons"
)

func InsertCartItem(cartID int, body []byte) (validators.CartItemRO, commons.ErrorMessage) {
	var err commons.ErrorMessage

	_, err = GetCart(cartID)

	if err != (commons.ErrorMessage{}) {
		err.ErrorType = "not_found"
		err.Message = "Cart not found"
	}

	cartItemDTO := validators.CartItemDTO{}
	json.Unmarshal(body, &cartItemDTO)

	if cartItemDTO.Product == "" {
		err.ErrorType = "invalid_field"
		err.Message = "Product cannot be empty"
	}

	if cartItemDTO.Quantity <= 0 {
		err.ErrorType = "invalid_field"
		err.Message = "Quantity cannot be negative"
	}

	cartItem := models.CartItem{
		Tablename: "cart_items",
		CartID:    cartID,
		Product:   cartItemDTO.Product,
		Quantity:  cartItemDTO.Quantity,
	}

	if err == (commons.ErrorMessage{}) {
		cartItem, _ = cartItem.Insert()
	}

	return cartItem.ToResponseObject(), err
}

func GetCartItem(itemID int) (validators.CartItemRO, commons.ErrorMessage) {
	cartItem := models.CartItem{
		ID:        itemID,
		Tablename: "cart_items",
	}

	var err commons.ErrorMessage

	cartItem, _ = cartItem.GetOne()

	if cartItem.ID == 0 {
		err.ErrorType = "not_found"
		err.Message = "Cart item not found"
	}

	return cartItem.ToResponseObject(), err
}

func DeleteCartItem(cartID int, itemID int) (validators.CartItemRO, commons.ErrorMessage) {
	cartItem := models.CartItem{
		Tablename: "cart_items",
		CartID:    cartID,
		ID:        itemID,
	}

	_, err := GetCart(cartID)

	if err != (commons.ErrorMessage{}) {
		err.ErrorType = "not_found"
		err.Message = "Cart not found"

		return cartItem.ToResponseObject(), err
	}

	_, err = GetCartItem(itemID)

	if err != (commons.ErrorMessage{}) {
		err.ErrorType = "not_found"
		err.Message = "Cart item not found"

		return cartItem.ToResponseObject(), err
	}

	cartItem.Delete()

	return cartItem.ToResponseObject(), err
}
