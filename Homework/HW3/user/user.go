package user

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Ticket_id uuid.UUID
}

func NewUser(id, firstName, lastName, email, ticketId string) (User, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return User{}, fmt.Errorf("Geçersiz UUID: %v", err)
	}

	parsedTicketID, err := uuid.Parse(ticketId)
	if err != nil {
		return User{}, fmt.Errorf("Geçersiz Ticket UUID: %v", err)
	}

	return User{
		Id:        parsedID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Ticket_id: parsedTicketID,
	}, nil
}
