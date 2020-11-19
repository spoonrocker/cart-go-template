package model

type Item struct {
	Id       int    `json:"id"`
	CartId   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	Id    int    `json:"id"`
	Items []Item `json:"items"`
}
