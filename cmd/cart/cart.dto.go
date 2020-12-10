package cart

type CartDTO struct {
	ID int
}

type CartItemDTO struct {
	ID       int
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
	CartID   int    `json:"cart_id"`
}

type CartRO struct {
	ID    int          `json:"id"`
	Items []CartItemRO `json:"items"`
}

type CartItemRO struct {
	ID       int    `json:"id"`
	CartID   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
