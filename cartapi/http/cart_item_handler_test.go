package http

import (
	"bytes"
	"cartapi/cartapi"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type CartItemServiceMock struct {
	mock.Mock
}

func (cism *CartItemServiceMock) CreateCartItem(cartItem *cartapi.CartItem) (*cartapi.CartItem, error) {
	args := cism.Called(cartItem)
	return args.Get(0).(*cartapi.CartItem), args.Error(1)
}

func (cism *CartItemServiceMock) DeleteCartItem(cartId, id int) error {
	args := cism.Called(cartId, id)
	return args.Error(0)
}

func TestCartItemCreateRequest(t *testing.T) {
	persistedItem := cartapi.CartItem{
		Id:       1,
		CartId:   1,
		Product:  "Bananas",
		Quantity: 1,
	}
	persistedItemBytes, _ := json.Marshal(persistedItem)
	persistedItemBytes = append(persistedItemBytes, '\n')

	validItem := cartapi.CartItem{
		CartId:   1,
		Product:  "Bananas",
		Quantity: 1,
	}
	validItemBytes, _ := json.Marshal(validItem)

	missingCart := cartapi.CartItem{CartId: 2}
	missingCartBytes, _ := json.Marshal(missingCart)

	invalidItem := cartapi.CartItem{CartId: 1, Product: ""}
	invalidItemBytes, _ := json.Marshal(invalidItem)

	testCases := []struct {
		name           string
		cartId         string
		givenItem      []byte
		expectedItem   []byte
		expectedStatus int
	}{
		{
			name:           "create request with valid item and params",
			cartId:         "1",
			givenItem:      validItemBytes,
			expectedItem:   persistedItemBytes,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "create request with invalid cartId",
			cartId:         "notInteger",
			givenItem:      nil,
			expectedItem:   nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "create request with malformed body",
			cartId:         "1",
			givenItem:      []byte(`{"product": 1122}`),
			expectedItem:   nil,
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "create request with missing cart",
			cartId:         "2",
			givenItem:      missingCartBytes,
			expectedItem:   nil,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "create request with invalid cart item",
			cartId:         "1",
			givenItem:      invalidItemBytes,
			expectedItem:   nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	serviceMock := CartItemServiceMock{}
	serviceMock.On("CreateCartItem", &validItem).Return(&persistedItem, nil)
	serviceMock.On("CreateCartItem", &missingCart).Return((*cartapi.CartItem)(nil), cartapi.ErrCartNotFound)
	serviceMock.On("CreateCartItem", &invalidItem).Return((*cartapi.CartItem)(nil), cartapi.ErrInvalidItemProduct)

	router := NewRouter(CartHandler{}, NewCartItemHandler(&serviceMock))

	for _, tc := range testCases {
		req := httptest.NewRequest("POST", fmt.Sprintf("/cart/%s/item", tc.cartId), bytes.NewBuffer(tc.givenItem))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if tc.expectedStatus != rr.Code {
			t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rr.Code)
		}

		retBody := rr.Body.Bytes()
		if tc.expectedItem != nil && !bytes.Equal(tc.expectedItem, retBody) {
			t.Error("Failed with different body")
		}
	}
}

func TestCartItemDeleteRequest(t *testing.T) {
	existingCartId := 1
	existingItemId := 1
	missingCartId := 2
	missingItemId := 2

	testCases := []struct {
		name           string
		cartId         string
		itemId         string
		expectedStatus int
	}{
		{
			name:           "delete request with valid params",
			cartId:         strconv.Itoa(existingCartId),
			itemId:         strconv.Itoa(existingItemId),
			expectedStatus: http.StatusOK,
		},
		{
			name:           "delete request with invalid cartId",
			cartId:         "notInteger",
			itemId:         "1",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "delete request with invalid id",
			cartId:         "1",
			itemId:         "notInteger",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "delete request with non existing cart",
			cartId:         strconv.Itoa(missingCartId),
			itemId:         strconv.Itoa(existingItemId),
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "delete request with non existing item",
			cartId:         strconv.Itoa(existingCartId),
			itemId:         strconv.Itoa(missingItemId),
			expectedStatus: http.StatusNotFound,
		},
	}

	serviceMock := CartItemServiceMock{}
	serviceMock.On("DeleteCartItem", existingCartId, existingItemId).Return(nil)
	serviceMock.On("DeleteCartItem", existingCartId, missingItemId).Return(cartapi.ErrMissingCartOrItem)
	serviceMock.On("DeleteCartItem", missingCartId, existingItemId).Return(cartapi.ErrMissingCartOrItem)

	router := NewRouter(CartHandler{}, NewCartItemHandler(&serviceMock))

	for _, tc := range testCases {
		req := httptest.NewRequest("DELETE", fmt.Sprintf("/cart/%s/item/%s", tc.cartId, tc.itemId), nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if tc.expectedStatus != rr.Code {
			t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rr.Code)
		}
	}
}
