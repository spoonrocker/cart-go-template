package postgres

import (
	"cartapi/cartapi"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCartItemStore_Create(t *testing.T) {
	db, mock := mockDB()
	store := NewCartItemStore(*db)

	testCases := []struct {
		name          string
		givenItem     *cartapi.CartItem
		expectedItem  *cartapi.CartItem
		expectedError error
	}{
		{
			name: "create with existing cart",
			givenItem: &cartapi.CartItem{
				CartId:   1,
				Product:  "Bananas",
				Quantity: 2,
			},
			expectedItem: &cartapi.CartItem{
				Id:       1,
				CartId:   1,
				Product:  "Bananas",
				Quantity: 2,
			},
			expectedError: nil,
		},
	}

	mock.ExpectQuery(createCartItemQuery).
		WithArgs(1, "Bananas", 2).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	for _, tc := range testCases {
		cart, err := store.Create(tc.givenItem)
		if !assert.EqualValues(t, tc.expectedItem, cart) {
			t.Error("Failed with different cart item")
		}
		if err != tc.expectedError {
			t.Error("Failed with different error")
		}
	}
}

func TestCartItemStore_Delete(t *testing.T) {
	db, mock := mockDB()
	store := NewCartItemStore(*db)

	existingCartId := 1
	existingItemId := 1
	missingCartId := 3
	missingItemId := 4

	testCases := []struct {
		name          string
		givenCartId   int
		givenItemId   int
		expectedError error
	}{
		{
			name:          "delete with existing cart",
			givenCartId:   existingCartId,
			givenItemId:   existingItemId,
			expectedError: nil,
		},
		{
			name:          "delete with missing cart",
			givenCartId:   missingCartId,
			givenItemId:   existingItemId,
			expectedError: cartapi.ErrMissingCartOrItem,
		},
		{
			name:          "delete with missing cart item",
			givenCartId:   existingItemId,
			givenItemId:   missingItemId,
			expectedError: cartapi.ErrMissingCartOrItem,
		},
	}

	mock.ExpectExec(deleteCartItemQuery).
		WithArgs(existingCartId, existingItemId).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(deleteCartItemQuery).
		WithArgs(missingCartId, existingItemId).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec(deleteCartItemQuery).
		WithArgs(existingCartId, missingItemId).
		WillReturnResult(sqlmock.NewResult(0, 0))

	for _, tc := range testCases {
		err := store.Delete(tc.givenCartId, tc.givenItemId)
		if err != tc.expectedError {
			t.Error("Failed with different error")
		}
	}
}
