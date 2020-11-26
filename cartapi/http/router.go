package http

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type Router struct {
	chi.Router
}

func NewRouter(cartHandler CartHandler, cartItemHandler CartItemHandler) Router {
	router := chi.NewRouter()

	router.Use(middleware.Heartbeat("/ping"))
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Route("/cart", cartHandler.addRoutes)
	router.Route("/cart/{cartId}/item", cartItemHandler.addRoutes)

	return Router{router}
}

func intPathParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, param))
	if err == nil {
		return id, nil
	}

	http.Error(w, fmt.Sprintf("%s must be an integer", param), http.StatusBadRequest)
	return 0, err
}
