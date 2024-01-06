package controller

import (
	"crudHomeWork/models"
	"fmt"
	"github.com/google/uuid"
)

func (c Controller) CreateUsers() {
	var (
		firstName, lastName, email, phone string
	)
	fmt.Print("Enter user's first name : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter user's last name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter user's email : ")
	fmt.Scan(&email)
	fmt.Print("Enter user's phone : ")
	fmt.Scan(&phone)

	user := models.Users{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	id, err := c.Store.UsersStorege.Insert(user)
	if err != nil {
		fmt.Println("error while inserting user inside controller err : ", err.Error())
		return
	}

	fmt.Println("Id : ", id)
}

func (c Controller) GetByIdUsers() {
	idStr := ""
	fmt.Print("Enter user id : ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("It is not uuid type !")
		return
	}
	user, err := c.Store.UsersStorege.GetById(id)
	if err != nil {
		fmt.Println("error while getting by id user inside controller err: ", err.Error())
		return
	}
	fmt.Println(user)
}

func (c Controller) GetListUsers() {
	users, err := c.Store.UsersStorege.GetList()
	if err != nil {
		fmt.Println("Error while getting users list inside controller err: ", err.Error())
		return
	}
	fmt.Println(users)
}

func (c Controller) UpdateUsers() {
	var (
		idStr, firstName, lastName, email, phone string
	)
	fmt.Print("Enter user's id : ")
	fmt.Scan(&idStr)
	fmt.Print("Enter user's first name : ")
	fmt.Scan(&firstName)
	fmt.Print("Enter user's last name : ")
	fmt.Scan(&lastName)
	fmt.Print("Enter user's email : ")
	fmt.Scan(&email)
	fmt.Print("Enter user's phone : ")
	fmt.Scan(&phone)

	user := models.Users{
		Id:        uuid.MustParse(idStr),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := c.Store.UsersStorege.Update(user); err != nil {
		fmt.Println("Error while updating user inside controller err: ")
		return
	}

	fmt.Println("Data successfully updated !")

}

func (c Controller) DeleteUsers() {
	var (
		idstr string
	)
	fmt.Print("id : ", idstr)
	fmt.Scan(&idstr)

	id, err := uuid.Parse(idstr)
	if err != nil {
		fmt.Println("Id type is not uuid !")
		return
	}
	if err := c.Store.UsersStorege.Delete(id); err != nil {
		fmt.Println("error while deleting user inside controller err : ", err.Error())
		return
	}

	fmt.Println("Data successfully deleted !")

}
