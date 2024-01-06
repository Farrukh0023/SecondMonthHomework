package postgres

import (
	"crudHomeWork/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	Db                  *sql.DB
	UsersStorege        UsersRepo
	ProductStorege      ProductsRepo
	OrderStorege        OrdersRepo
	OrderProductStoerge OrderProductsRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}
	usersRepo := NewUsersRepo(db)
	productsRepo := NewProductsRepo(db)
	ordersRepo := NewOrderRepo(db)
	orderProductsRepo := NewOrderProductsRepo(db)

	return Store{
		Db:                  db,
		UsersStorege:        usersRepo,
		ProductStorege:      productsRepo,
		OrderStorege:        ordersRepo,
		OrderProductStoerge: orderProductsRepo,
	}, nil
}
