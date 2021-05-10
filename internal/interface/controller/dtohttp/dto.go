package controller

// CartResponse represents json response for the CreateCart and GetCart handlers.
type CartResponse struct {
	ID    int            `json:"id"`    // Cart ID
	Items []ItemResponse `json:"items"` // Items in the cart
}

// AddItemRequest represents json request for the AddItem handler.
type AddItemRequest struct {
	Product  string `json:"product"`  // Product title
	Quantity int    `json:"quantity"` // Quantity of the products in the item
}

// ItemResponse represents json request for the AddItem handler.
type ItemResponse struct {
	ID       int    `json:"id"`       // Item ID
	CartID   int    `json:"cart_id"`  // ID of the cart in which item was placed
	Product  string `json:"product"`  // Product title
	Quantity int    `json:"quantity"` // Quantity of the products in the item
}

// RemoveItemResponse represents json response for the RemoveItem handler.
type RemoveItemResponse struct {
}

// ErrorResponse represents json response for the cases when the error is occurred.
type ErrorResponse struct {
	Message string `json:"message"` // Message of the error
}
