//go:generate mockgen -destination=mocks/service_mock.go -package=mocks github.com/Kleiber/cart-go-template/src/model Model
package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Model interface {
	InsertCart(cart Cart) error
	SelectCart(cartId int) (*Cart, error)
	SelectItemCart(cartId, itemId int) (*Item, error)
	ListItemsCart(cartId int) ([]Item, error)
	InsertItemCart(cartId int, item Item) error
	DeleteItemCart(cartId, itemId int) error
}

type ModelCart struct {
	DB *sql.DB
}

func NewModel(user, password, dbname string) (Model, error) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &ModelCart{DB: db}, nil
}

func (m *ModelCart) InsertCart(cart Cart) error {
	_, err := m.DB.Exec("insert into dbcart.cart(id) values (?)", cart.Id)
	if err != nil {
		return err
	}
	return nil
}

func (m *ModelCart) SelectCart(cartId int) (*Cart, error) {
	cart := Cart{}

	err := m.DB.QueryRow("select id from dbcart.cart where id=?", cartId).Scan(cart.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &CartNotFoundError{cartId}
		} else {
			return nil, err
		}
	}

	return &cart, nil
}

func (m *ModelCart) SelectItemCart(cartId, itemId int) (*Item, error) {
	item := Item{}

	err := m.DB.QueryRow("select id,id_cart,product,quantity from dbcart.items where cart_id=? and id=?", cartId, itemId).Scan(item.Id, item.CartId, item.Product, item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ItemNotFoundError{itemId}
		} else {
			return nil, err
		}
	}

	return &item, nil
}

func (m *ModelCart) ListItemsCart(cartId int) ([]Item, error) {
	rows, err := m.DB.Query("select * from dbcart.items where id_cart=?", cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var i Item
		err = rows.Scan(&i.Id, &i.CartId, &i.Product, i.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}

func (m *ModelCart) InsertItemCart(cartId int, item Item) error {
	_, err := m.DB.Exec("insert into dbcart.items(id, id_cart, product, quantity) values (?, ?, ?, ?)", item.Id, cartId, item.Product, item.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (m *ModelCart) DeleteItemCart(cartId, itemId int) error {
	_, err := m.DB.Exec("delete from dbcart.items where cart_id=$1 and id=$2", cartId, itemId)
	if err != nil {
		return err

	}
	return nil
}
