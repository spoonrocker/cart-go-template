// Package doc Cart API
//
// Documentation for Cart API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package doc

import "github.com/fedo3nik/cart-go-api/internal/domain/model"

type addItemRequest struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

// swagger:parameters addItemParams addItem
type addItemParams struct {
	// in: path
	// example: 1
	CartID int `json:"cartID"`
	// in: body
	Body addItemRequest
}

// swagger:parameters removeItemParams removeItem
type removeItemParams struct {
	// in: path
	// example: 3
	CartID int `json:"cartID"`
	// in: path
	// example: 5
	ItemID int `json:"itemID"`
}

// swagger:parameters getCartParams getCart
type getCartParams struct {
	// in: path
	// example: 1
	CartID int `json:"cartID"`
}

// New cart created successfully
// swagger:response cartResponse
type cartResponse struct {
	// ID of the new cart
	CartID int `json:"id"`
	// Array of cartItems
	Items []model.CartItem `json:"items"`
}

// CartItem added into the cart successfully
// swagger:response addItemResponse
type addItemResponse struct {
	// ID of the new cartItem
	ID int `json:"id"`
	// CartID in which item was placed
	CartID int `json:"cart_id"`
	// Product title
	Product string `json:"product"`
	// Quantity of the products in the cartItem
	Quantity int `json:"quantity"`
}

// CartItem removed from the cart successfully
// swagger:response removeItemResponse
type removeItemResponse struct {
}

// Error caused
// swagger:response errorResponse
type errorResponse struct {
	// Error message
	Message string `json:"message"`
}
