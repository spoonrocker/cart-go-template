package cart

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/cabogabo/cart-api/cmd/commons"
)

func getCart(ID int) (CartRO, error) {
	cart := CartRO{}

	err := commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("carts"))

		v := b.Get([]byte(commons.Itob(ID)))

		err := json.Unmarshal([]byte(v), &cart)

		if err != nil {
			return err
		}

		if cart.ID == 0 {
			return fmt.Errorf("Not found")
		}

		return nil
	})

	if err == nil {
		cart = getCartItemsByCart(ID, cart)
	}

	return cart, err
}

func getCartItem(ID int) (CartItemRO, error) {
	cartItem := CartItemRO{}

	err := commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cart_items"))

		v := b.Get([]byte(commons.Itob(ID)))

		err := json.Unmarshal([]byte(v), &cartItem)

		if err != nil {
			return err
		}

		if cartItem.ID == 0 {
			return fmt.Errorf("Not found")
		}

		return nil
	})

	return cartItem, err
}

func getCartItemsByCart(ID int, cart CartRO) CartRO {
	cartItems := []CartItemRO{}

	commons.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cart_items"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			cartItem := CartItemRO{}

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
