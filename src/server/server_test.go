package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Kleiber/cart-go-template/src/model"
	cartMock "github.com/Kleiber/cart-go-template/src/service/mocks"
)

type ServerTest struct {
	server      *httptest.Server
	cartService *cartMock.MockService
}

func setupServerTest(t *testing.T) (*require.Assertions, *gomock.Controller, ServerTest, func()) {
	r := require.New(t)
	ctrl := gomock.NewController(t)

	cartService := cartMock.NewMockService(ctrl)
	handler := NewServer(cartService).Handler
	server := httptest.NewServer(handler)
	serverTest := ServerTest{
		server:      server,
		cartService: cartService,
	}

	finally := func() {
		serverTest.server.Close()
	}

	return r, ctrl, serverTest, finally
}

func TestServer(t *testing.T) {
	r, _, serverTest, finally := setupServerTest(t)
	defer finally()

	cart := model.Cart{
		Id: 1,
		Items: []model.Item{
			{Id: 1, CartId: 1, Product: "Shoes", Quantity: 10},
			{Id: 2, CartId: 1, Product: "Socks", Quantity: 5},
		},
	}

	testCase := []struct {
		apiPath            string
		apiName            string
		httpMethod         string
		expectedStatusCode int
	}{
		{
			apiPath:            "/carts",
			apiName:            "create-cart",
			httpMethod:         http.MethodPost,
			expectedStatusCode: http.StatusOK,
		},
		{
			apiPath:            "/carts/1",
			apiName:            "view-cart",
			httpMethod:         http.MethodGet,
			expectedStatusCode: http.StatusOK,
		},
		{
			apiPath:            "/carts/1/items",
			apiName:            "add-item",
			httpMethod:         http.MethodPost,
			expectedStatusCode: http.StatusOK,
		},
		{
			apiPath:            "/carts/1/items/1",
			apiName:            "remove-item",
			httpMethod:         http.MethodDelete,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCase {
		body := &bytes.Buffer{}

		switch tc.apiName {
		case "create-cart":
			serverTest.cartService.EXPECT().CreateNewEmptyCart().Return(&cart, nil)
		case "view-cart":
			serverTest.cartService.EXPECT().GetCart(cart.Id).Return(&cart, nil)
		case "add-item":
			serverTest.cartService.EXPECT().AddNewItemToCart(cart.Id, cart.Items[0]).Return(&cart.Items[0], nil)
			buf, err := json.Marshal(cart.Items[0])
			r.NoError(err)
			body = bytes.NewBuffer(buf)
		case "remove-item":
			serverTest.cartService.EXPECT().RemoveItemFromCart(cart.Id, cart.Items[0].Id).Return(cart.Items, nil)
		}

		url := fmt.Sprintf("%s%s", serverTest.server.URL, tc.apiPath)
		req, err := http.NewRequest(tc.httpMethod, url, body)
		r.NoError(err)

		res, err := serverTest.server.Client().Do(req)
		r.NoError(err)
		r.Equal(tc.expectedStatusCode, res.StatusCode)
	}
}
