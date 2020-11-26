package postgres

import (
	"cartapi/cartapi"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func NewDB() sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
		cartapi.Config.Database.Host, cartapi.Config.Database.Port,
		cartapi.Config.Database.Username, cartapi.Config.Database.Password,
		cartapi.Config.Database.Database, cartapi.Config.Database.Schema,
		cartapi.Config.Database.SslMode)

	db, err := sqlx.Connect("postgres", psqlconn)
	if err != nil {
		panic("failed to connect to the database")
	}
	return *db
}
