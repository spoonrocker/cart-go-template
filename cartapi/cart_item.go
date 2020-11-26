package cartapi

import "errors"

var ErrMissingCartOrItem = errors.New("missing cart or cart item")
var ErrInvalidItemProduct = errors.New("invalid cart item product")
var ErrInvalidItemQuantity = errors.New("invalid cart item quantity")

type CartItem struct {
	Id       int    `db:"id" json:"id"`
	CartId   int    `db:"cart_id" json:"cart_id"`
	Product  string `db:"product" json:"product"`
	Quantity int    `db:"quantity" json:"quantity"`
}

type CartItemService interface {
	CreateCartItem(cartItem *CartItem) (*CartItem, error)
	DeleteCartItem(cartId, id int) error
}

type CartItemServiceImpl struct {
	store CartItemStore
}

func NewCartItemService(store CartItemStore) CartItemService {
	return CartItemServiceImpl{store: store}
}

type CartItemStore interface {
	Create(item *CartItem) (*CartItem, error)
	Delete(cartId, id int) error
}

func (cis CartItemServiceImpl) CreateCartItem(cartItem *CartItem) (*CartItem, error) {
	if cartItem.Product == "" {
		return nil, ErrInvalidItemProduct
	}
	if cartItem.Quantity <= 0 {
		return nil, ErrInvalidItemQuantity
	}

	return cis.store.Create(cartItem)
}

func (cis CartItemServiceImpl) DeleteCartItem(cartId, id int) error {
	return cis.store.Delete(cartId, id)
}
