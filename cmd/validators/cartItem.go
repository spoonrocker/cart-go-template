package validators

type CartItemDTO struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type CartItemRO struct {
	ID       int    `json:"id"`
	CartID   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
