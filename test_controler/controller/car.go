package controller

import (
	"fmt"
	"test/models"
	"time"

	"github.com/google/uuid"
)

func (c Controller) CreateCar() {
	car := getCarInfo()

	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year intput is not correct")
		return
	}

	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error while creating data inside controller err: ", err.Error())
		return
	}
	fmt.Println("id: ", id)
}

func (c Controller) GetCarByID() {
	idStr := ""
	fmt.Print("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}

	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Print("error while getting car by id err: ", err.Error())
		return
	}
	fmt.Println("your car is: ", car)
}

func (c Controller) GetCarList() {
	cars, err := c.Store.CarStorage.GetList()
	if err != nil {
		fmt.Println("Error while getting list err : ", err.Error())
		return
	}

	fmt.Println(cars)

}

func (c Controller) UpdateCar() {
	car := getCarInfo()

	if !checkCarInfo(car) {
		return
	}

	err := c.Store.CarStorage.Update(car)
	if err != nil {
		fmt.Println("error while updating car!", err)
		return
	}
	if car.ID.String() != "" {
		fmt.Println("Successfully updated!")
	} else {
		fmt.Println("Successfully created!")
	}

}

func (c Controller) DeleteCar() {
	idStr := ""
	fmt.Print("Enter id : ")
	fmt.Scan(&idStr)

	id := uuid.MustParse(idStr)

	if err := c.Store.CarStorage.Delete(id); err != nil {
		fmt.Println("error while deleting db err : ", err.Error())
	}

	fmt.Println("Successfully deleted !")

}

func checkCarInfo(car models.Car) bool {
	if car.Year <= 0 || car.Year > time.Now().Year()+1 {
		fmt.Println("year intput is not correct")
		return false
	}

	return true
}

func getCarInfo() models.Car {
	var (
		model, brand, idStr string
		cmd, year           int
	)

a:
	fmt.Print(`enter command: 
			1 - create
			2 - update
	`)
	fmt.Scan(&cmd)

	if cmd == 2 {
		fmt.Print("enter id: ")
		fmt.Scan(&idStr)

		fmt.Print("enter model and brand: ")
		fmt.Scan(&model, &brand)

		fmt.Print("enter year: ")
		fmt.Scan(&year)
	} else if cmd == 1 {

		fmt.Print("enter model and brand: ")
		fmt.Scan(&model, &brand)

		fmt.Print("enter year: ")
		fmt.Scan(&year)

	} else {
		fmt.Println("not found")
		goto a
	}

	if idStr != "" {
		return models.Car{
			ID:    uuid.MustParse(idStr),
			Model: model,
			Brand: brand,
			Year:  year,
		}
	}

	return models.Car{
		Model: model,
		Brand: brand,
		Year:  year,
	}
}
