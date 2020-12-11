package validators

type CartDTO struct {
	ID int
}

type CartRO struct {
	ID    int          `json:"id"`
	Items []CartItemRO `json:"items"`
}
