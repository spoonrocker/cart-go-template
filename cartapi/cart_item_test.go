package cartapi

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type CartItemStoreMock struct {
	mock.Mock
}

func (cism *CartItemStoreMock) Create(item *CartItem) (*CartItem, error) {
	args := cism.Called(item)
	return args.Get(0).(*CartItem), args.Error(1)
}

func (cism *CartItemStoreMock) Delete(cartId, id int) error {
	args := cism.Called(cartId, id)
	return args.Error(0)
}

var persistedCartItem = CartItem{
	Id:       7,
	CartId:   1,
	Product:  "Bananas",
	Quantity: 2,
}

func TestCartItemService_CreateCartItem(t *testing.T) {
	itemWithExistingCart := CartItem{
		CartId:   1,
		Product:  "Bananas",
		Quantity: 2,
	}
	itemWithMissingCart := CartItem{
		CartId:   999,
		Product:  "Bananas",
		Quantity: 2,
	}
	itemWithMissingProduct := CartItem{
		CartId:   1,
		Product:  "",
		Quantity: 2,
	}
	itemWithNonPositiveQuantity := CartItem{
		CartId:   1,
		Product:  "Bananas",
		Quantity: 0,
	}

	testCases := []struct {
		name          string
		givenItem     *CartItem
		expectedItem  *CartItem
		expectedError error
	}{
		{
			name:          "create valid cart",
			givenItem:     &itemWithExistingCart,
			expectedItem:  &persistedCartItem,
			expectedError: nil,
		},
		{
			name:          "item with missing cart",
			givenItem:     &itemWithMissingCart,
			expectedItem:  nil,
			expectedError: ErrCartNotFound,
		},
		{
			name:          "item with missing product",
			givenItem:     &itemWithMissingProduct,
			expectedItem:  nil,
			expectedError: ErrInvalidItemProduct,
		},
		{
			name:          "item with missing cart",
			givenItem:     &itemWithNonPositiveQuantity,
			expectedItem:  nil,
			expectedError: ErrInvalidItemQuantity,
		},
	}

	storeMock := CartItemStoreMock{}
	storeMock.On("Create", &itemWithExistingCart).Return(&persistedCartItem, nil)
	storeMock.On("Create", &itemWithMissingCart).Return((*CartItem)(nil), ErrCartNotFound)

	service := NewCartItemService(&storeMock)

	for _, tc := range testCases {
		item, err := service.CreateCartItem(tc.givenItem)
		if item != tc.expectedItem {
			t.Error("failed with different cart item expectation")
		}
		if err != tc.expectedError {
			t.Error("failed with different error expectation")
		}
	}
}

func TestCartItemService_DeleteCartItem(t *testing.T) {
	existingItemId := 10
	existingCartId := 10
	missingId := 999
	missingCartId := 9999

	testCases := []struct {
		name          string
		givenId       int
		givenCartId   int
		expectedError error
	}{
		{
			name:          "delete existing cart item",
			givenId:       existingItemId,
			givenCartId:   existingCartId,
			expectedError: nil,
		},
		{
			name:          "non existing item",
			givenId:       existingItemId,
			givenCartId:   missingCartId,
			expectedError: ErrMissingCartOrItem,
		},
		{
			name:          "non existing cart",
			givenId:       missingId,
			givenCartId:   existingCartId,
			expectedError: ErrMissingCartOrItem,
		},
	}

	storeMock := CartItemStoreMock{}
	storeMock.On("Delete", existingCartId, existingItemId).Return(nil)
	storeMock.On("Delete", existingCartId, missingId).Return(ErrMissingCartOrItem)
	storeMock.On("Delete", missingCartId, existingItemId).Return(ErrMissingCartOrItem)

	service := NewCartItemService(&storeMock)

	for _, tc := range testCases {
		err := service.DeleteCartItem(tc.givenCartId, tc.givenId)
		if err != tc.expectedError {
			t.Error("failed with different error expectation")
		}
	}
}
