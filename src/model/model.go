package model

import "time"

type Item struct {
	Id       int    `json:"id"`
	CartId   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

type Cart struct {
	Id    int       `json:"id"`
	Date  time.Time `json:"-"`
	Items []Item    `json:"items"`
}
