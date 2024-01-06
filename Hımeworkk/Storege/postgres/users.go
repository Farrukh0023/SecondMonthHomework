package postgres

import (
	"crudHomeWork/models"
	"database/sql"
	"github.com/google/uuid"
)

type UsersRepo struct {
	Db *sql.DB
}

func NewUsersRepo(db *sql.DB) UsersRepo {
	return UsersRepo{Db: db}
}

func (u UsersRepo) Insert(user models.Users) (string, error) {
	id := uuid.New()

	if _, err := u.Db.Exec(`insert into users values ($1,$2,$3,$4,$5)`, id, user.FirstName, user.LastName, user.Email, user.Phone); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (u UsersRepo) GetById(id uuid.UUID) (models.Users, error) {
	user := models.Users{}
	if err := u.Db.QueryRow(`SELECT *FROM users`).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
	); err != nil {
		return models.Users{}, err
	}

	return user, nil

}

func (u UsersRepo) GetList() ([]models.Users, error) {
	users := []models.Users{}

	rows, err := u.Db.Query(`SELECT *FROM users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.Users{}

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u UsersRepo) Update(user models.Users) error {
	if _, err := u.Db.Exec(`UPDATE users set first_name = $1,last_name = $2,email = $3,phone = $4 where id = $5`,
		user.FirstName, user.LastName, user.Email, user.Phone, user.Id); err != nil {
		return err
	}
	return nil
}

func (u UsersRepo) Delete(id uuid.UUID) error {
	if _, err := u.Db.Exec(`delete from users where id = $1`, id); err != nil {
		return err
	}
	return nil
}
