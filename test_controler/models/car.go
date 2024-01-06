package models

import "github.com/google/uuid"

type Car struct {
	ID    uuid.UUID
	Model string
	Brand string
	Year  int
}
