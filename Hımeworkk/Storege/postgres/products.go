package postgres

import (
	"crudHomeWork/models"
	"database/sql"
	"github.com/google/uuid"
)

type ProductsRepo struct {
	Db *sql.DB
}

func NewProductsRepo(db *sql.DB) ProductsRepo {
	return ProductsRepo{Db: db}
}

func (p ProductsRepo) Insert(products models.Products) (string, error) {
	id := uuid.New()

	if _, err := p.Db.Exec(`insert into products values ($1,$2,$3)`, products.Id, products.Name, products.Price); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (p ProductsRepo) GetById(id uuid.UUID) (models.Products, error) {
	pro := models.Products{}

	if err := p.Db.QueryRow(`select id, name, price  from products where id = $1`, id).Scan(
		&pro.Id,
		&pro.Name,
		&pro.Price,
	); err != nil {
		return models.Products{}, err
	}

	return pro, nil
}

func (p ProductsRepo) GetList() ([]models.Products, error) {

	products := []models.Products{}

	rows, err := p.Db.Query(`Select *from products`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := models.Products{}

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil

}

func (p ProductsRepo) Update(product models.Products) error {

	if _, err := p.Db.Exec(`UPDATE products set name = $1, price = $2 where Id = $3`,
		product.Name, product.Price, product.Id); err != nil {
		return err
	}

	return nil

}

func (p ProductsRepo) Delete(id uuid.UUID) error {

	if _, err := p.Db.Exec(`DELETE from products where id = $1`, id); err != nil {
		return err
	}
	return nil
}
