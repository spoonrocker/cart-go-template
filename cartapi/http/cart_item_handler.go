package http

import (
	"cartapi/cartapi"
	"cartapi/cartapi/logger"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CartItemHandler struct {
	service cartapi.CartItemService
}

func NewCartItemHandler(cartItemService cartapi.CartItemService) CartItemHandler {
	return CartItemHandler{service: cartItemService}
}

func (cih *CartItemHandler) addRoutes(r chi.Router) {
	r.Post("/", cih.handleCreateItem)
	r.Delete("/{id}", cih.handleDeleteItem)
}

func (cih *CartItemHandler) handleCreateItem(w http.ResponseWriter, r *http.Request) {
	cartId, err := intPathParam("cartId", w, r)
	if err != nil {
		writeJsonResponse(ErrorResponse{Code: InvalidParameter, Message: "invalid cart id"},
			http.StatusBadRequest, w, r, "createCartItem")
		logger.WithReqIdAndAction(log.Debug(), r, "createCartItem").
			Str("cartId", chi.URLParam(r, "cartId")).
			Msg("attempt to add cart item to missing cart")
		return
	}

	var cartItem cartapi.CartItem
	if err = json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
		writeJsonResponse(ErrorResponse{Code: UnprocessableBody, Message: err.Error()},
			http.StatusUnprocessableEntity, w, r, "createCartItem")
		logger.WithReqIdAndAction(log.Debug().Err(err), r, "createCartItem").
			Msg("failed to parse request")
		return
	}
	cartItem.CartId = cartId

	createdItem, err := cih.service.CreateCartItem(&cartItem)
	if err != nil {
		if err == cartapi.ErrCartNotFound {
			writeJsonResponse(ErrorResponse{Code: UnprocessableBody, Message: "attempt to add item to missing cart"},
				http.StatusNotFound, w, r, "createCartItem")
			logger.WithReqIdAndAction(log.Debug(), r, "createCartItem").
				Int("cartId", cartId).
				Msg("cart not found")
			return
		}
		if err == cartapi.ErrInvalidItemProduct || err == cartapi.ErrInvalidItemQuantity {
			writeJsonResponse(ErrorResponse{Code: InvalidCartItem, Message: err.Error()},
				http.StatusBadRequest, w, r, "createCartItem")
			logger.WithReqIdAndAction(log.Debug(), r, "createCartItem").
				Int("cartId", cartId).
				Msg(err.Error())
			return
		}
		writeJsonResponse(ErrorResponse{Code: UnknownError, Message: "cart item creation failed"},
			http.StatusInternalServerError, w, r, "createCartItem")
		logger.WithReqIdAndAction(log.Error().Stack().Err(err), r, "createCartItem").
			Msg("cart item creation failed")
		return
	}

	writeJsonResponse(createdItem, http.StatusCreated, w, r, "createCartItem")
}

func (cih *CartItemHandler) handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	cartId, err := intPathParam("cartId", w, r)
	if err != nil {
		writeJsonResponse(ErrorResponse{Code: InvalidParameter, Message: "invalid cart id"},
			http.StatusBadRequest, w, r, "deleteCartItem")
		logger.WithReqIdAndAction(log.Debug(), r, "deleteCartItem").
			Str("cartId", chi.URLParam(r, "cartId")).
			Msg("attempt to delete cart item with invalid cartId")
		return
	}

	id, err := intPathParam("id", w, r)
	if err != nil {
		writeJsonResponse(ErrorResponse{Code: InvalidParameter, Message: "attempt to delete cart item with invalid id"},
			http.StatusBadRequest, w, r, "deleteCartItem")
		logger.WithReqIdAndAction(log.Debug(), r, "deleteCartItem").
			Str("id", chi.URLParam(r, "id")).
			Msg("attempt to delete cart item with invalid id")
		return
	}

	if err = cih.service.DeleteCartItem(cartId, id); err != nil {
		if err == cartapi.ErrMissingCartOrItem {
			writeJsonResponse(ErrorResponse{Code: EntityNotFound, Message: "cart or cart item not found"},
				http.StatusNotFound, w, r, "deleteCartItem")
			logger.WithReqIdAndAction(log.Debug(), r, "deleteCartItem").
				Int("cartId", cartId).
				Int("id", id).
				Msg("cart or cart item not found")
			return
		}

		writeJsonResponse(ErrorResponse{Code: UnknownError, Message: "unknown error while deleting cart item"},
			http.StatusInternalServerError, w, r, "deleteCartItem")
		logger.WithReqIdAndAction(log.Error().Stack().Err(err), r, "deleteCartItem").
			Int("cartId", cartId).
			Int("id", id).
			Msg("failed to delete cart item")
	}

	w.WriteHeader(http.StatusOK)
}
