package postgres

import (
	"crudHomeWork/models"
	"database/sql"
	"github.com/google/uuid"
)

type OrderProductsRepo struct {
	Db *sql.DB
}

func NewOrderProductsRepo(db *sql.DB) OrderProductsRepo {
	return OrderProductsRepo{Db: db}
}

func (o OrderProductsRepo) Insert(op models.OrderProducts) (string, error) {
	id := uuid.New()
	if _, err := o.Db.Exec(`insert into order_products values ($1,$2,$3,$4,$5)`,
		id, op.OrderID, op.ProductID, op.Quantity, op.Price); err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o OrderProductsRepo) GetById(id uuid.UUID) (models.OrderProducts, error) {
	op := models.OrderProducts{}

	err := o.Db.QueryRow(`SELECT id ,order_id,product_id,quantity,price from order_products where id = $1`, id).Scan(
		&op.Id,
		&op.OrderID,
		&op.ProductID,
		&op.Quantity,
		&op.Price,
	)
	if err != nil {
		return models.OrderProducts{}, err
	}

	return op, nil
}

func (o OrderProductsRepo) GetList() ([]models.OrderProducts, error) {
	ops := []models.OrderProducts{}

	rows, err := o.Db.Query(`select *from order_products`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		op := models.OrderProducts{}

		err := rows.Scan(
			&op.Id,
			&op.OrderID,
			&op.ProductID,
			&op.Quantity,
			&op.Price,
		)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func (o OrderProductsRepo) Update(op models.OrderProducts) error {
	if _, err := o.Db.Exec(`update order_products set order_id = $1,product_id = $2,quantity = $3,price = $4 where id = $5`,
		op.OrderID, op.ProductID, op.Quantity, op.Price, op.Id); err != nil {
		return err
	}
	return nil
}

func (o OrderProductsRepo) Delete(id uuid.UUID) error {
	if _, err := o.Db.Exec(`delete from order_products where id = $1`, id); err != nil {
		return err
	}
	return nil

}
