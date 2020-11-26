package cartapi

import "errors"

var ErrCartNotFound = errors.New("cart not found")

type Cart struct {
	Id    int        `db:"id" json:"id"`
	Items []CartItem `json:"items"`
}

type CartService interface {
	Cart(id int) (*Cart, error)
	CreateCart() (*Cart, error)
}

type CartServiceImpl struct {
	store CartStore
}

func NewCartService(store CartStore) CartService {
	return CartServiceImpl{store: store}
}

type CartStore interface {
	Cart(id int) (*Cart, error)
	Create() (*Cart, error)
}

func (cs CartServiceImpl) Cart(id int) (*Cart, error) {
	return cs.store.Cart(id)
}

func (cs CartServiceImpl) CreateCart() (*Cart, error) {
	return cs.store.Create()
}
