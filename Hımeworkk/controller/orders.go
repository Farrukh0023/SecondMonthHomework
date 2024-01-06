package controller

import (
	"crudHomeWork/models"
	"fmt"
	"github.com/google/uuid"
)

func (c Controller) CreateOrders() {
	var (
		userIdStr string
		amount    int
	)
	fmt.Print("Enter order user id : ")
	fmt.Scan(&userIdStr)
	fmt.Print("Enter order amount : ")
	fmt.Scan(&amount)
	order := models.Orders{
		Amount: amount,
		UserID: uuid.MustParse(userIdStr),
	}
	id, err := c.Store.OrderStorege.Insert(order)
	if err != nil {
		fmt.Println("error while creating order inside controller err: ", err.Error())
		return
	}

	fmt.Println("Id : ", id)
}

func (c Controller) GetByIdOrders() {
	idStr := ""
	fmt.Print("Enter order id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	order, err := c.Store.OrderStorege.GetById(id)
	if err != nil {
		fmt.Println("error while getting by id order inside controller err : ", err.Error())
		return
	}

	fmt.Println(order)
}

func (c Controller) GetListOrders() {
	orders, err := c.Store.OrderStorege.GetList()
	if err != nil {
		fmt.Println("error while getting orders list inside controller err: ", err.Error())
		return
	}
	fmt.Println(orders)
}

func (c Controller) UpdateOrders() {
	var (
		idStr, userIdStr string
		amount           int
	)
	fmt.Print("Enter order id : ")
	fmt.Scan(&idStr)
	fmt.Print("Enter order user id : ")
	fmt.Scan(&userIdStr)
	fmt.Print("Enter order amount : ")
	fmt.Scan(&amount)
	order := models.Orders{
		Id:     uuid.MustParse(idStr),
		Amount: amount,
		UserID: uuid.MustParse(userIdStr),
	}
	if err := c.Store.OrderStorege.Update(order); err != nil {
		fmt.Println("error while updating order inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully updated !")
}

func (c Controller) DeleteOrders() {
	idStr := ""
	fmt.Print("Enter order id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	if err := c.Store.OrderStorege.Delete(id); err != nil {
		fmt.Println("error while deleting order inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully deleted !")
}
