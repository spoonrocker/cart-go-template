package postgres

import (
	"cartapi/cartapi"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCartStore_Cart(t *testing.T) {
	db, mock := mockDB()
	store := NewCartStore(*db)

	testCases := []struct {
		name          string
		id            int
		expectedCart  *cartapi.Cart
		expectedError error
	}{
		{
			name: "existing cart with items",
			id:   1,
			expectedCart: &cartapi.Cart{
				Id: 1,
				Items: []cartapi.CartItem{
					{
						Id:       1,
						CartId:   1,
						Product:  "Bananas",
						Quantity: 2,
					},
					{
						Id:       2,
						CartId:   1,
						Product:  "Apples",
						Quantity: 3,
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "missing cart",
			id:            2,
			expectedCart:  nil,
			expectedError: cartapi.ErrCartNotFound,
		},
	}

	mock.ExpectQuery(cartByIdQuery).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectQuery(cartItemsByCartIdQuery).
		WithArgs(1).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "cart_id", "product", "quantity"}).
				AddRow(1, 1, "Bananas", 2).
				AddRow(2, 1, "Apples", 3))
	mock.ExpectQuery(cartByIdQuery).WithArgs(2).WillReturnError(sql.ErrNoRows)

	for _, tc := range testCases {
		cart, err := store.Cart(tc.id)
		if !assert.EqualValues(t, tc.expectedCart, cart) {
			t.Error("Failed with different cart")
		}
		if err != tc.expectedError {
			t.Error("Failed with different error")
		}
	}
}
