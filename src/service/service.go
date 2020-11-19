//go:generate mockgen -destination=mocks/service_mock.go -package=mocks github.com/Kleiber/cart-go-template/src/service Service
package service

import (
	"time"

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
	cart := model.Cart{
		Items: []model.Item{},
		Date:  time.Now(),
	}

	newCart, err := c.CartModel.InsertCart(cart)
	if err != nil {
		return nil, err
	}

	return newCart, nil
}

func (c *CartService) GetCart(cartId int) (*model.Cart, error) {
	newCart, err := c.CartModel.SelectCart(cartId)
	if err != nil {
		return nil, err
	}

	items, err := c.CartModel.ListItemsCart(cartId)
	if err != nil {
		return nil, err
	}

	newCart.Items = items
	return newCart, nil
}

func (c *CartService) AddNewItemToCart(cartId int, item model.Item) (*model.Item, error) {
	_, err := c.CartModel.SelectCart(cartId)
	if err != nil {
		return nil, err
	}

	newItem, err := c.CartModel.InsertItemCart(cartId, item)
	if err != nil {
		return nil, err
	}

	return newItem, nil
}

func (c *CartService) RemoveItemFromCart(cartId, itemId int) ([]model.Item, error) {
	_, err := c.CartModel.SelectItemCart(cartId, itemId)
	if err != nil {
		return nil, err
	}

	err = c.CartModel.DeleteItemCart(cartId, itemId)
	if err != nil {
		return nil, err
	}

	items, err := c.CartModel.ListItemsCart(cartId)
	if err != nil {
		return nil, err
	}

	return items, nil
}
