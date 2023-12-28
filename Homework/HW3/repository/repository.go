package repository

import (
	"database/sql"
	"errors"
	"tiket/ticket"
	"tiket/user"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	Db *sql.DB
}

func (r Repository) New(db *sql.DB) Repository {
	return Repository{
		Db: db,
	}
}

func (r Repository) AddTicket(ticket ticket.Ticket) error {
	_, err := r.Db.Exec(`insert into ticket values ($1, $2, $3, $4)`, ticket.Id, ticket.From, ticket.To, ticket.Date)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetTicketById(id uuid.UUID) (ticket.Ticket, error) {
	t := ticket.Ticket{}
	row := r.Db.QueryRow(`select * from ticket where id = $1`, id)

	if err := row.Scan(&t.Id, &t.From, &t.To, &t.Date); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return ticket.Ticket{}, err
		}
	}
	return t, nil
}

func (r Repository) GetAllTickets() ([]ticket.Ticket, error) {
	ts := []ticket.Ticket{}
	rows, err := r.Db.Query(`select * from ticket`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		t := ticket.Ticket{}

		if err = rows.Scan(&t.Id, &t.From, &t.To, &t.Date); err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func (r Repository) UpdateTicket(t ticket.Ticket) error {
	_, err := r.Db.Exec(`update ticket set where_from = $1, where_to = $2, date = $3 where id = $4`,
		t.From, t.To, t.Date, t.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) DeleteTicket(id uuid.UUID) error {
	_, err := r.Db.Exec(`delete from ticket where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) AddUser(user user.User) error {
	_, err := r.Db.Exec(`insert into users values ($1, $2, $3, $4, $5)`,
		user.Id, user.FirstName, user.LastName, user.Email, user.Ticket_id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetUserById(id uuid.UUID) (user.User, error) {
	u := user.User{}
	row := r.Db.QueryRow(`select * from users where id = $1`, id)

	if err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Ticket_id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return user.User{}, err
		}
	}
	return u, nil
}

func (r Repository) GetAllUsers() ([]user.User, error) {
	us := []user.User{}
	rows, err := r.Db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := user.User{}

		if err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Ticket_id); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

func (r Repository) UpdateUsers(u user.User) error {
	_, err := r.Db.Exec(`update users set first_name = $1, last_name = $2, email = $3 where id = $4`,
		u.FirstName, u.LastName, u.Email, u.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) DeleteUser(id uuid.UUID) error {
	_, err := r.Db.Exec(`delete from users where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) ReportAll(from, to string) ([]Report, error) {
	reps := []Report{}
	rows, err := r.Db.Query(`select u.first_name, u.last_name, u.email, t.where_from, t.where_to, t.date from users as u join ticket as t on u.ticket_id = t.Id
	where t.where_from = $1 and t.where_to = $2`, from, to)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rep := Report{}
		if err := rows.Scan(&rep.FirstName, &rep.LastName, &rep.Email, &rep.From, &rep.To, &rep.Date); err != nil {
			return nil, err
		}
		reps = append(reps, rep)
	}

	return reps, nil
}

type Report struct {
	FirstName string
	LastName  string
	Email     string
	From      string
	To        string
	Date      time.Time
}
