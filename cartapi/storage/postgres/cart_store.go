package postgres

import (
	"cartapi/cartapi"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

const cartByIdQuery = `SELECT * FROM cart WHERE id=$1`
const cartItemsByCartIdQuery = `SELECT * FROM cart_item WHERE cart_id=$1`
const createCartQuery = `INSERT INTO cart DEFAULT VALUES RETURNING id`

type CartStore struct {
	*sqlx.DB
}

func NewCartStore(db sqlx.DB) CartStore {
	return CartStore{&db}
}

func (cs CartStore) Cart(id int) (*cartapi.Cart, error) {
	var cart cartapi.Cart

	if err := cs.Get(&cart, cartByIdQuery, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, cartapi.ErrCartNotFound
		}
		return nil, err
	}

	items := []cartapi.CartItem{}
	if err := cs.Select(&items, cartItemsByCartIdQuery, id); err != nil {
		return nil, err
	}

	cart.Items = items
	return &cart, nil
}

func (cs CartStore) Create() (*cartapi.Cart, error) {
	id := 0

	if err := cs.QueryRow(createCartQuery).Scan(&id); err != nil {
		return nil, err
	}

	return &cartapi.Cart{
		Id:    id,
		Items: []cartapi.CartItem{},
	}, nil
}
