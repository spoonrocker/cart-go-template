package services

import (
	"github.com/cabogabo/cart-api/cmd/commons"
	"github.com/cabogabo/cart-api/cmd/models"
	"github.com/cabogabo/cart-api/cmd/validators"
)

func InsertCart() validators.CartRO {
	cart := models.Cart{
		Tablename: "carts",
	}
	cart, _ = cart.Insert()
	return cart.ToResponseObject()
}

func GetCart(cartID int) (validators.CartRO, commons.ErrorMessage) {
	cart := models.Cart{
		ID:        cartID,
		Tablename: "carts",
	}

	var err commons.ErrorMessage

	cart, _ = cart.GetOne()

	if cart.ID == 0 {
		err.ErrorType = "not_found"
		err.Message = "Cart not found"
	}

	return cart.ToResponseObject(), err
}
