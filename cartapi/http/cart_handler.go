package http

import (
	"cartapi/cartapi"
	"cartapi/cartapi/logger"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"net/http"
)

type CartHandler struct {
	service cartapi.CartService
}

func NewCartHandler(cartService cartapi.CartService) CartHandler {
	return CartHandler{service: cartService}
}

func (ch *CartHandler) addRoutes(r chi.Router) {
	r.Get("/{id}", ch.handleLookup)
	r.Post("/", ch.handleCreate)
}

func (ch *CartHandler) handleLookup(w http.ResponseWriter, r *http.Request) {
	id, err := intPathParam("id", w, r)
	if err != nil {
		writeJsonResponse(ErrorResponse{Code: InvalidParameter, Message: "invalid cart id"},
			http.StatusBadRequest, w, r, "lookupCart")
		logger.WithReqIdAndAction(log.Debug(), r, "lookupCart").
			Str("id", chi.URLParam(r, "id")).
			Msg("invalid cart id")
		return
	}

	cart, err := ch.service.Cart(id)
	if err != nil {
		if err == cartapi.ErrCartNotFound {
			writeJsonResponse(ErrorResponse{Code: EntityNotFound, Message: "cart not found"},
				http.StatusNotFound, w, r, "lookupCart")
			logger.WithReqIdAndAction(log.Debug(), r, "lookupCart").
				Int("id", id).
				Msg("cart not found")
			return
		}

		writeJsonResponse(ErrorResponse{Code: UnknownError, Message: "cart lookup failed"},
			http.StatusInternalServerError, w, r, "lookupCart")
		logger.WithReqIdAndAction(log.Error().Stack().Err(err), r, "lookupCart").
			Int("id", id).
			Msg("cart lookup failed")
		return
	}

	writeJsonResponse(cart, http.StatusOK, w, r, "lookupCart")
}

func (ch *CartHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	cart, err := ch.service.CreateCart()

	if err != nil {
		writeJsonResponse(ErrorResponse{Code: UnknownError, Message: "cart creation failed"},
			http.StatusInternalServerError, w, r, "createCart")
		logger.WithReqIdAndAction(log.Error().Stack().Err(err), r, "createCart").
			Msg("cart creation failed")
		return
	}

	writeJsonResponse(cart, http.StatusCreated, w, r, "createCart")
}
