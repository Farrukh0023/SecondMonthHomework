package postgres

import (
	"crudHomeWork/models"
	"database/sql"
	"github.com/google/uuid"
)

type OrdersRepo struct {
	Db *sql.DB
}

func NewOrderRepo(db *sql.DB) OrdersRepo {
	return OrdersRepo{Db: db}
}

func (o OrdersRepo) Insert(order models.Orders) (string, error) {
	id := uuid.New()
	if _, err := o.Db.Exec(`insert into orders values ($1,$2,$3,$4)`, id, order.Amount, order.UserID, order.CreatedAT); err != nil {
		return "", err
	}

	return id.String(), nil

}

func (o OrdersRepo) GetById(id uuid.UUID) (models.Orders, error) {
	order := models.Orders{}
	if err := o.Db.QueryRow(`SELECT id,amount,user_id,created_at from orders where id = $1`, id).Scan(
		&order.Id,
		&order.Amount,
		&order.UserID,
		&order.CreatedAT,
	); err != nil {
		return models.Orders{}, err
	}

	return order, nil
}

func (o OrdersRepo) GetList() ([]models.Orders, error) {
	orders := []models.Orders{}

	rows, err := o.Db.Query(`SELECT *from orders`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := models.Orders{}

		if err := rows.Scan(
			&order.Id,
			&order.Amount,
			&order.UserID,
			&order.CreatedAT,
		); err != nil {
			return nil, err
		}

		orders = append(orders, order)

	}
	return orders, nil
}

func (o OrdersRepo) Update(order models.Orders) error {
	if _, err := o.Db.Exec(`update orders set amount = $1,user_id = $2 where id = 3`, order.Amount, order.UserID, order.Id); err != nil {
		return err
	}
	return nil
}

func (o OrdersRepo) Delete(id uuid.UUID) error {
	if _, err := o.Db.Exec(`delete from orders where id = $1`, id); err != nil {
		return err
	}
	return nil
}
