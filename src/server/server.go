package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Kleiber/cart-go-template/src/model"
	"github.com/Kleiber/cart-go-template/src/service"

	"github.com/gorilla/mux"
)

const (
	varCart = "cartId"
	varItem = "itemId"
)

type Server struct {
	Handler     http.Handler
	CartService service.Service
}

func NewServer(cartService service.Service) *Server {
	s := &Server{
		CartService: cartService,
	}
	s.Handler = s.buildHandler()

	return s
}

func (s *Server) buildHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/carts", s.createCart).Methods("POST")
	router.HandleFunc("/carts/{cartId}", s.viewCart).Methods("GET")
	router.HandleFunc("/carts/{cartId}/items", s.AddToCart).Methods("POST")
	router.HandleFunc("/carts/{cartId}/items/{itemId}", s.RemoveFromCart).Methods("DELETE")

	return router
}

func (s *Server) createCart(w http.ResponseWriter, r *http.Request) {
	newCart, err := s.CartService.CreateNewEmptyCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(newCart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) viewCart(w http.ResponseWriter, r *http.Request) {
	cartId, _ := strconv.Atoi(mux.Vars(r)[varCart])

	cart, err := s.CartService.GetCart(cartId)
	if err != nil {
		if _, ok := err.(*model.CartNotFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) AddToCart(w http.ResponseWriter, r *http.Request) {
	cartId, _ := strconv.Atoi(mux.Vars(r)[varCart])

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var item model.Item
	err = json.Unmarshal(buf, &item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cartItem, err := s.CartService.AddNewItemToCart(cartId, item)
	if err != nil {
		if _, ok := err.(*model.InvaidQuantityError); ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, ok := err.(*model.InvalidProductError); ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if _, ok := err.(*model.CartNotFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cartItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *Server) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	cartId, _ := strconv.Atoi(mux.Vars(r)[varCart])
	itemId, _ := strconv.Atoi(mux.Vars(r)[varItem])

	cartItems, err := s.CartService.RemoveItemFromCart(cartId, itemId)
	if err != nil {
		if _, ok := err.(*model.ItemNotFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cartItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
