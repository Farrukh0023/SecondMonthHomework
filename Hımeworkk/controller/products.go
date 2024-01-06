package controller

import (
	"crudHomeWork/models"
	"fmt"
	"github.com/google/uuid"
)

func (c Controller) CreateProduct() {
	p := getProductInfo()
	id, err := c.Store.ProductStorege.Insert(p)
	if err != nil {
		fmt.Println("error while creating data inside controller err: ", err.Error())
		return
	}

	fmt.Println("id : ", id)

}

func (c Controller) GetProductByID() {
	idStr := ""
	fmt.Print("id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("Wrong Id !")
		return
	}

	product, err := c.Store.ProductStorege.GetById(id)
	if err != nil {
		fmt.Println("error while getting by id inside controller err : ", err.Error())
		return
	}

	fmt.Println("your product is ", product)

}

func (c Controller) GetProductList() {

	products, err := c.Store.ProductStorege.GetList()
	if err != nil {
		fmt.Println("error while getting product list inside controller err : ", err.Error())
	}
	fmt.Println(products)
}

func (c Controller) UpdateProduct() {
	p := getProductInfo()
	if err := c.Store.ProductStorege.Update(p); err != nil {

		fmt.Println("error while updating db inside controller err : ", err.Error())
		return

	}
	fmt.Println("Data successfully updated ! ")
}

func (c Controller) DeleteProduct() {
	idStr := ""

	fmt.Println("Enter id : ", idStr)
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("Wrong id ! ")
		return
	}

	if err := c.Store.ProductStorege.Delete(id); err != nil {
		fmt.Println("error while deleting date from db inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully deleted !")
}

func getProductInfo() models.Products {
	var (
		name  string
		price int
		idStr string
		cmd   int
		id    string
	)
a:
	fmt.Println(`		1 => create
		2 => update
		3 => exit`)
	fmt.Scan(&cmd)

	if cmd == 1 {
		fmt.Print("Enter Product name : ")
		fmt.Scan(&name)
		fmt.Print("Enter product price : ")
		fmt.Scan(&price)
		fmt.Print("Enter Category id : ")
		fmt.Scan(&idStr)

		return models.Products{
			Name:  name,
			Price: price,
		}
	} else if cmd == 2 {
		fmt.Print("Enter id : ", id)
		fmt.Scan(&id)
		fmt.Print("Enter Product name : ")
		fmt.Scan(&name)
		fmt.Print("Enter product price : ")
		fmt.Scan(&price)
		fmt.Print("Enter Category id : ")
		fmt.Scan(&idStr)

		return models.Products{
			Id:    uuid.MustParse(id),
			Name:  name,
			Price: price,
		}
	} else if cmd == 3 {
		return models.Products{}

	} else {
		fmt.Println("Command not found !")
		goto a
	}

}
