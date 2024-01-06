package models

import "github.com/google/uuid"

type Driver struct {
	ID       uuid.UUID
	FullName string
	Phone    string
	CarID    uuid.UUID
}
