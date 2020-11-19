package model

import "fmt"

type CartNotFoundError struct {
	CartId int
}

func (e *CartNotFoundError) Error() string {
	return fmt.Sprintf("cart woth id %d not found", e.CartId)
}

type ItemNotFoundError struct {
	CartId int
	ItemId int
}

func (e *ItemNotFoundError) Error() string {
	return fmt.Sprintf("item with id %d was not found in cart with id %d", e.ItemId, e.CartId)
}

type InvaidQuantityError struct {
}

func (e *InvaidQuantityError) Error() string {
	return fmt.Sprintf("product value cannot be negative")
}

type InvalidProductError struct {
}

func (e *InvalidProductError) Error() string {
	return fmt.Sprintf("product value cannot be empty")
}
