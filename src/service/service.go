//go:generate mockgen -destination=mocks/service_mock.go -package=mocks github.com/Kleiber/cart-go-template/src/service Service
package service

import (
	"github.com/Kleiber/cart-go-template/src/model"
)

type Service interface {
	CreateNewEmptyCart() (*model.Cart, error)
	GetCart(cartId int) (*model.Cart, error)
	AddNewItemToCart(cartId int, item model.Item) (*model.Item, error)
	RemoveItemFromCart(cartId, itemId int) ([]model.Item, error)
}

type CartService struct {
	CartModel model.Model
}

func NewCartService(cartModel model.Model) Service {
	return &CartService{
		CartModel: cartModel,
	}
}

func (c *CartService) CreateNewEmptyCart() (*model.Cart, error) {
	return nil, nil
}

func (c *CartService) GetCart(cartId int) (*model.Cart, error) {
	return nil, nil
}

func (c *CartService) AddNewItemToCart(cartId int, item model.Item) (*model.Item, error) {
	return nil, nil
}

func (c *CartService) RemoveItemFromCart(cartId, itemId int) ([]model.Item, error) {
	return nil, nil
}
