package controller

import (
	"fmt"
	"github.com/google/uuid"
	"test/models"
	"unicode"
)

func (c Controller) CreateDriver() {
	driver := getDriverInfo()

	if !checkPhoneNumber(driver.Phone) {
		fmt.Println("the phone number format is not correct!")
		return
	}

	id, err := c.Store.DriverStorage.Insert(driver)
	if err != nil {
		fmt.Println("error while inserting driver inside controller err: ", err.Error())
		return
	}

	fmt.Println("your new driver's id is: ", id)
}

func (c Controller) GetByIdDriver() {
	idStr := ""
	fmt.Print("Enter Driver id : ")
	fmt.Scan(&idStr)
	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !", err.Error())
		return
	}
	driver, err := c.Store.DriverStorage.GetById(id)
	if err != nil {
		fmt.Println("error whilw getting by id driver inside controller err : ", err.Error())
		return
	}

	fmt.Println(driver)
}

func (c Controller) GetDriverList() {
	drivers, err := c.Store.DriverStorage.GetList()
	if err != nil {
		fmt.Println("error while getting driver list inside controller err : ", err.Error())
		return
	}
	fmt.Println(drivers)
}

func (c Controller) UpdateDriver() {
	driver := getDriverInfo()

	if err := c.Store.DriverStorage.Update(driver); err != nil {
		fmt.Println("error while updating driver date inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully updated !")
}

func (c Controller) DeleteDriver() {
	idStr := ""
	fmt.Println("enter id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}

	if err := c.Store.DriverStorage.Delete(id); err != nil {
		fmt.Println("error while deleting driver row inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully deleted !")
}

func checkPhoneNumber(phone string) bool {
	for _, r := range phone {
		if !(unicode.IsNumber(r) || r == '+') {
			return false
		}
	}

	return true
}

func getDriverInfo() models.Driver {
	var (
		fullName, phone, idStr, car_id string
		cmd                            int
	)
a:
	fmt.Println(`		1 => create
		2 => Update`)
	fmt.Scan(&cmd)

	if cmd == 1 {

		fmt.Print("enter driver's full name: ")
		fmt.Scan(&fullName)

		fmt.Print("enter phone: ")
		fmt.Scan(&phone)

		fmt.Print("enter car id : ")
		fmt.Scan(&car_id)

		return models.Driver{
			FullName: fullName,
			Phone:    phone,
			CarID:    uuid.MustParse(car_id),
		}
	} else if cmd == 2 {
		fmt.Print("enter id : ")
		fmt.Scan(&idStr)
		fmt.Print("enter driver's full name: ")
		fmt.Scan(&fullName)

		fmt.Println("enter phone: ")
		fmt.Scan(&phone)

		return models.Driver{
			ID:       uuid.MustParse(idStr),
			FullName: fullName,
			Phone:    phone,
		}
	} else {
		fmt.Println("wrong input !")
		goto a
	}
}
