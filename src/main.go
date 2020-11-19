package main

import (
	"fmt"
	"net/http"

	"github.com/Kleiber/cart-go-template/src/model"
	"github.com/Kleiber/cart-go-template/src/server"
	"github.com/Kleiber/cart-go-template/src/service"
)

const (
	dbHost     = "localhost"
	dbPort     = 5434
	dbUser     = "cart"
	dbPassword = "cart"
	dbName     = "dbcart"
)

func main() {
	cartModel, err := model.NewModel(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		panic(fmt.Sprintf("Cannot initialize database connection: %s", err))
	}
	cartService := service.NewCartService(cartModel)
	cartServer := server.NewServer(cartService)

	address := fmt.Sprintf("localhost:3000")
	fmt.Printf("Server listening on %s", address)

	err = http.ListenAndServe(address, cartServer.Handler)
	if err != nil {
		panic(fmt.Sprintf("Cannot initialize server: %s", err))
	}
}
