package cartapi

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type CartStoreMock struct {
	mock.Mock
}

func (csm *CartStoreMock) Cart(id int) (*Cart, error) {
	args := csm.Called(id)
	return args.Get(0).(*Cart), args.Error(1)
}

func (csm *CartStoreMock) Create() (*Cart, error) {
	args := csm.Called()
	return args.Get(0).(*Cart), args.Error(1)
}

var persistedCart = Cart{
	Id:    1,
	Items: []CartItem{},
}

func TestCartService_Cart(t *testing.T) {
	testCases := []struct {
		name          string
		givenId       int
		expectedCart  *Cart
		expectedError error
	}{
		{
			name:          "lookup of an existing cart",
			givenId:       1,
			expectedCart:  &persistedCart,
			expectedError: nil,
		},
		{
			name:          "lookup of a non existing cart",
			givenId:       2,
			expectedCart:  nil,
			expectedError: ErrCartNotFound,
		},
	}

	storeMock := CartStoreMock{}
	storeMock.On("Cart", 1).Return(&persistedCart, nil)
	storeMock.On("Cart", 2).Return((*Cart)(nil), ErrCartNotFound)

	service := NewCartService(&storeMock)

	for _, tc := range testCases {
		cart, err := service.Cart(tc.givenId)
		if cart != tc.expectedCart {
			t.Error("failed with different cart expectation")
		}
		if err != tc.expectedError {
			t.Error("failed with different error expectation")
		}
	}
}

func TestCartService_CreateCart(t *testing.T) {
	storeMock := CartStoreMock{}
	storeMock.On("Create").Return(&persistedCart, nil)

	service := NewCartService(&storeMock)

	cart, err := service.CreateCart()
	if cart != &persistedCart {
		t.Error("failed with different cart expectation")
	}
	if err != nil {
		t.Error("failed with different error expectation")
	}
}
