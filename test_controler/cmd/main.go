package main

import (
	"log"
	"test/config"
	"test/controller"
	"test/storage/postgres"
)

func main() {
	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err:", err.Error())
		return
	}
	defer store.DB.Close()

	con := controller.New(store)

	//con.CreateDriver()
	//con.GetByIdDriver()
	//con.GetDriverList()
	//con.UpdateDriver()
	//con.DeleteDriver()

	//con.CreateCar()
	//con.GetCarByID()
	//con.GetCarList()
	//con.UpdateCar()
	//con.DeleteCar()

}