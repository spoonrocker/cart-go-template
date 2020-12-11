package models

import (
	"encoding/json"

	"github.com/cabogabo/cart-api/cmd/validators"

	"github.com/boltdb/bolt"
	"github.com/cabogabo/cart-api/cmd/commons"
)

type CartItem struct {
	Tablename string `json:"-"`
	ID        int    `json:"id"`
	CartID    int    `json:"cart_id"`
	Product   string `json:"product"`
	Quantity  int    `json:"quantity"`
}

func (cartItem CartItem) Insert() (CartItem, error) {
	commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cartItem.Tablename))
		id, _ := b.NextSequence()
		cartItem.ID = int(id)
		buf, err := json.Marshal(cartItem)
		if err != nil {
			return err
		}
		return b.Put(commons.Itob(cartItem.ID), buf)
	})

	return cartItem.GetOne()
}

func (cartItem CartItem) GetOne() (CartItem, error) {
	auxCartItem := CartItem{}
	err := commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cartItem.Tablename))

		v := b.Get([]byte(commons.Itob(cartItem.ID)))

		err := json.Unmarshal([]byte(v), &auxCartItem)

		if err != nil {
			return err
		}
		return nil
	})

	return auxCartItem, err
}

func (cartItem CartItem) Delete() error {
	return commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cartItem.Tablename))
		err := b.Delete([]byte(commons.Itob(cartItem.ID)))
		return err
	})

}

func (cartItem CartItem) ToResponseObject() validators.CartItemRO {
	return validators.CartItemRO{
		ID:       cartItem.ID,
		CartID:   cartItem.CartID,
		Quantity: cartItem.Quantity,
		Product:  cartItem.Product,
	}
}
