package models

import "github.com/google/uuid"

type Users struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
