package models

import "github.com/google/uuid"

type Products struct {
	Id    uuid.UUID
	Name  string
	Price int
}
