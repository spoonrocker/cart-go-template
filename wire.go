//+build wireinject

package cart_go_template

import (
	"cartapi/cartapi"
	"cartapi/cartapi/http"
	"cartapi/cartapi/storage/postgres"
	"github.com/google/wire"
)

func InitServer() http.Server {
	wire.Build(
		http.NewServer,
		http.NewRouter,
		http.NewCartHandler,
		http.NewCartItemHandler,
		cartapi.NewCartService,
		cartapi.NewCartItemService,
		postgres.NewDB,
		postgres.NewCartStore,
		postgres.NewCartItemStore,
		wire.Bind(new(cartapi.CartStore), new(postgres.CartStore)),
		wire.Bind(new(cartapi.CartItemStore), new(postgres.CartItemStore)))

	return http.Server{}
}
