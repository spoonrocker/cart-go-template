package http

import (
	"cartapi/cartapi"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CartServiceMock struct {
	mock.Mock
}

func (csm *CartServiceMock) Cart(id int) (*cartapi.Cart, error) {
	args := csm.Called(id)
	return args.Get(0).(*cartapi.Cart), args.Error(1)
}

func (csm *CartServiceMock) CreateCart() (*cartapi.Cart, error) {
	args := csm.Called()
	return args.Get(0).(*cartapi.Cart), args.Error(1)
}

var persistedCart = cartapi.Cart{
	Id:    1,
	Items: []cartapi.CartItem{},
}

func TestCartLookupRequest(t *testing.T) {
	existingCartId := 1
	missingCartId := 2

	testCases := []struct {
		name           string
		id             int
		expectedCart   *cartapi.Cart
		expectedStatus int
	}{
		{
			name:           "existing cart lookup request",
			id:             existingCartId,
			expectedCart:   &persistedCart,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing cart lookup request",
			id:             missingCartId,
			expectedCart:   &cartapi.Cart{},
			expectedStatus: http.StatusNotFound,
		},
	}

	serviceMock := CartServiceMock{}
	serviceMock.On("Cart", existingCartId).Return(&persistedCart, nil)
	serviceMock.On("Cart", missingCartId).Return((*cartapi.Cart)(nil), cartapi.ErrCartNotFound)

	router := NewRouter(NewCartHandler(&serviceMock), CartItemHandler{})

	for _, tc := range testCases {
		req := httptest.NewRequest("GET", fmt.Sprintf("/cart/%d", tc.id), nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		if tc.expectedStatus != rr.Code {
			t.Errorf("Expected status %d, but got %d", tc.expectedStatus, rr.Code)
		}

		var returnedBody cartapi.Cart
		err := json.Unmarshal(rr.Body.Bytes(), &returnedBody)
		if err != nil || !assert.EqualValues(t, tc.expectedCart, &returnedBody) {
			t.Error("Failed with different body")
		}
	}
}

func TestCartCreateRequest(t *testing.T) {
	serviceMock := CartServiceMock{}
	router := NewRouter(NewCartHandler(&serviceMock), CartItemHandler{})

	req := httptest.NewRequest("POST", "/cart", nil)

	serviceMock.On("CreateCart").Return(&persistedCart, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if http.StatusCreated != rr.Code {
		t.Errorf("Expected status %d, but got %d", http.StatusCreated, rr.Code)
	}

	var returnedBody cartapi.Cart
	err := json.Unmarshal(rr.Body.Bytes(), &returnedBody)
	if err != nil || !assert.EqualValues(t, persistedCart, returnedBody) {
		t.Error("Failed with different body")
	}
}
