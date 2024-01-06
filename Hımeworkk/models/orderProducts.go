package models

import "github.com/google/uuid"

type OrderProducts struct {
	Id        uuid.UUID
	OrderID   uuid.UUID
	ProductID uuid.UUID
	Quantity  int
	Price     int
}
