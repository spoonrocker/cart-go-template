//go:generate mockgen -destination=mocks/service_mock.go -package=mocks github.com/Kleiber/cart-go-template/src/model Model
package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Model interface {
	InsertCart(cart Cart) (*Cart, error)
	SelectCart(cartId int) (*Cart, error)
	SelectItemCart(cartId, itemId int) (*Item, error)
	ListItemsCart(cartId int) ([]Item, error)
	InsertItemCart(cartId int, item Item) (*Item, error)
	DeleteItemCart(cartId, itemId int) error
}

type ModelCart struct {
	DB *sql.DB
}

func NewModel(host string, port int, user string, password string, dbname string) (Model, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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

func (m *ModelCart) InsertCart(cart Cart) (*Cart, error) {
	id := 0
	err := m.DB.QueryRow("insert into cart(date) values($1) returning id", cart.Date).Scan(&id)
	if err != nil {
		return nil, err
	}

	cart.Id = id
	return &cart, nil
}

func (m *ModelCart) SelectCart(cartId int) (*Cart, error) {
	cart := Cart{
		Items: []Item{},
	}

	err := m.DB.QueryRow("select id from cart where id=$1", cartId).Scan(&cart.Id)
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

	err := m.DB.QueryRow("select id, cart_id, product, quantity from items where cart_id=$1 and id=$2", cartId, itemId).Scan(&item.Id, &item.CartId, &item.Product, &item.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ItemNotFoundError{cartId, itemId}
		} else {
			return nil, err
		}
	}

	return &item, nil
}

func (m *ModelCart) ListItemsCart(cartId int) ([]Item, error) {
	rows, err := m.DB.Query("select * from items where cart_id=$1", cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var i Item
		err = rows.Scan(&i.Id, &i.CartId, &i.Product, &i.Quantity)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	return items, nil
}

func (m *ModelCart) InsertItemCart(cartId int, item Item) (*Item, error) {
	if item.Quantity < 0 {
		return nil, &InvaidQuantityError{}
	}

	if item.Product == "" {
		return nil, &InvalidProductError{}
	}

	id := 0
	err := m.DB.QueryRow("insert into items(cart_id, product, quantity) values ($1, $2, $3) returning id", cartId, item.Product, item.Quantity).Scan(&id)
	if err != nil {
		return nil, err
	}

	item.Id = id
	item.CartId = cartId
	return &item, nil
}

func (m *ModelCart) DeleteItemCart(cartId, itemId int) error {
	_, err := m.DB.Exec("delete from items where cart_id=$1 and id=$2", cartId, itemId)
	if err != nil {
		return err

	}
	return nil
}
