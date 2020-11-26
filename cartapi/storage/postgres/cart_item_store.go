package postgres

import (
	"cartapi/cartapi"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const createCartItemQuery = `INSERT INTO cart_item (cart_id, product, quantity) VALUES ($1, $2, $3) RETURNING id`
const deleteCartItemQuery = `DELETE FROM cart_item WHERE cart_id=$1 AND id=$2`

type CartItemStore struct {
	*sqlx.DB
}

func NewCartItemStore(db sqlx.DB) CartItemStore {
	return CartItemStore{&db}
}

func (cis CartItemStore) Create(item *cartapi.CartItem) (*cartapi.CartItem, error) {
	lastId := 0
	err := cis.QueryRow(createCartItemQuery, item.CartId, item.Product, item.Quantity).Scan(&lastId)

	if err == nil {
		item.Id = lastId
		return &cartapi.CartItem{
			Id:       lastId,
			CartId:   item.CartId,
			Product:  item.Product,
			Quantity: item.Quantity,
		}, nil
	}

	if "foreign_key_violation" == err.(*pq.Error).Code.Name() {
		return nil, cartapi.ErrCartNotFound
	}
	return nil, err
}
func (cis CartItemStore) Delete(cartId, id int) error {
	res, err := cis.Exec(deleteCartItemQuery, cartId, id)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if affected != 1 {
		return cartapi.ErrMissingCartOrItem
	}

	return nil
}
