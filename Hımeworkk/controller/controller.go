package controller

import "crudHomeWork/Storege/postgres"

type Controller struct {
	Store postgres.Store
}

func New(store postgres.Store) Controller {
	return Controller{Store: store}
}
