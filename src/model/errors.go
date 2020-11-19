package model

import "fmt"

type CartNotFoundError struct {
	CartId int
}

func (e *CartNotFoundError) Error() string {
	return fmt.Sprintf("cart %q not found", e.CartId)
}

type ItemNotFoundError struct {
	ItemId int
}

func (e *ItemNotFoundError) Error() string {
	return fmt.Sprintf("item %q not found", e.ItemId)
}
