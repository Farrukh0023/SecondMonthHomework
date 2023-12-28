package ticket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	Id   uuid.UUID
	From string
	To   string
	Date time.Time
}

func NewTicket(id, from, to string, date time.Time) Ticket {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		// Geçersiz UUID durumu için hata kontrolü
		fmt.Println("Geçersiz UUID.")
		return Ticket{}
	}

	return Ticket{
		Id:   parsedID,
		From: from,
		To:   to,
		Date: date,
	}
}
