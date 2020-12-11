package models

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/cabogabo/cart-api/cmd/commons"
	"github.com/cabogabo/cart-api/cmd/validators"
)

type Cart struct {
	Tablename string     `json:"-"`
	ID        int        `json:"id"`
	Items     []CartItem `json:"items"`
}

func (cart Cart) Insert() (Cart, error) {
	commons.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cart.Tablename))
		id, _ := b.NextSequence()
		cart.ID = int(id)
		buf, err := json.Marshal(cart)
		if err != nil {
			return err
		}

		return b.Put(commons.Itob(cart.ID), buf)
	})

	return cart.GetOne()
}

func (cart Cart) GetOne() (Cart, error) {
	auxCart := Cart{
		Tablename: "carts",
	}
	err := commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(cart.Tablename))

		v := b.Get([]byte(commons.Itob(cart.ID)))

		err := json.Unmarshal([]byte(v), &auxCart)

		if err != nil {
			return err
		}

		return nil
	})

	if auxCart.ID > 0 {
		auxCart = auxCart.GetItemsByCart()
	}

	return auxCart, err
}

func (cart Cart) GetItemsByCart() Cart {
	cartItems := []CartItem{}
	commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cart_items"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			cartItem := CartItem{}

			err := json.Unmarshal([]byte(v), &cartItem)

			if err != nil {
				return err
			}

			if cartItem.CartID == cart.ID {
				cartItems = append(cartItems, cartItem)
			}
		}

		return nil
	})

	cart.Items = cartItems

	return cart
}

func (cart Cart) ToResponseObject() validators.CartRO {
	cartRO := validators.CartRO{
		ID: cart.ID,
	}

	cartItems := []validators.CartItemRO{}
	for _, item := range cart.Items {
		cartItem := item.ToResponseObject()
		cartItems = append(cartItems, cartItem)
	}

	cartRO.Items = cartItems
	return cartRO
}
