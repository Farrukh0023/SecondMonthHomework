package main

import (
	"crudHomeWork/Storege/postgres"
	"crudHomeWork/config"
	"crudHomeWork/controller"
	"log"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}
	defer store.Db.Close()

	con := controller.New(store)

	con.CreateUsers()
	//con.GetByIdUsers()
	//con.GetListUsers()
	//con.UpdateUsers()
	//con.DeleteUsers()

	//con.CreateProduct()
	//con.GetProductByID()
	//con.GetProductList()
	//con.UpdateProduct()
	//con.DeleteProduct()

	//con.CreateOrders()
	//con.GetByIdOrders()
	//con.GetListOrders()
	//con.UpdateOrders()
	//con.DeleteOrders()

	//con.CreateOrderProducts()
	//con.GetByIdOrderProducts()
	//con.GetByIdOrderProducts()
	//con.UpdateOrderProducts()
	//con.DeleteOrderProduct()
}
