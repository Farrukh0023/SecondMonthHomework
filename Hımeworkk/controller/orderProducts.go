package controller

import (
	"crudHomeWork/models"
	"fmt"
	"github.com/google/uuid"
)

func (c Controller) CreateOrderProducts() {
	var (
		orderIdStr, productIdStr string
		quantity, price          int
	)
	fmt.Print("Order id : ")
	fmt.Scan(&orderIdStr)
	fmt.Print("Product id : ")
	fmt.Scan(&productIdStr)
	fmt.Print("Quantity : ")
	fmt.Scan(&quantity)
	fmt.Print("Price : ")
	fmt.Scan(&price)

	op := models.OrderProducts{
		OrderID:   uuid.MustParse(orderIdStr),
		ProductID: uuid.MustParse(productIdStr),
		Quantity:  quantity,
		Price:     price,
	}

	id, err := c.Store.OrderProductStoerge.Insert(op)
	if err != nil {
		fmt.Println("error while creating order product inside controller err: ", err.Error())
		return
	}

	fmt.Println("Id : ", id)

}

func (c Controller) GetByIdOrderProducts() {
	idStr := ""

	fmt.Print("Id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	op, err := c.Store.OrderProductStoerge.GetById(id)
	if err != nil {
		fmt.Println("error while getting by id order product inside controller err : ", err.Error())
		return
	}

	fmt.Println(op)
}

func (c Controller) GetListOrderProducts() {
	ops, err := c.Store.OrderProductStoerge.GetList()
	if err != nil {
		fmt.Println("error while getting list of order product inside controller err : ", err.Error())
		return
	}
	fmt.Println(ops)
}

func (c Controller) UpdateOrderProducts() {
	var (
		idStr, orderIdStr, productIdStr string
		quantity, price                 int
	)
	fmt.Print("id : ")
	fmt.Scan(&idStr)
	fmt.Print("Order id : ")
	fmt.Scan(&orderIdStr)
	fmt.Print("Product id : ")
	fmt.Scan(&productIdStr)
	fmt.Print("Quantity : ")
	fmt.Scan(&quantity)
	fmt.Print("Price : ")
	fmt.Scan(&price)

	op := models.OrderProducts{
		Id:        uuid.MustParse(idStr),
		OrderID:   uuid.MustParse(orderIdStr),
		ProductID: uuid.MustParse(productIdStr),
		Quantity:  quantity,
		Price:     price,
	}

	if err := c.Store.OrderProductStoerge.Update(op); err != nil {
		fmt.Println("error while updating of order product inside controller err : ", err.Error())
		return
	}
	fmt.Println("Data successfully updated !")
}

func (c Controller) DeleteOrderProduct() {
	idStr := ""
	fmt.Print("id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	if err := c.Store.OrderProductStoerge.Delete(id); err != nil {
		fmt.Println("error while deleting of order product inside controller err : ", err.Error())
		return
	}
	fmt.Println("Data successfully deleted !")
}
