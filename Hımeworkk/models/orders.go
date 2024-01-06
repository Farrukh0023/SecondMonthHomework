package models

import "github.com/google/uuid"

type Orders struct {
	Id        uuid.UUID
	Amount    int
	UserID    uuid.UUID
	CreatedAT string
}
