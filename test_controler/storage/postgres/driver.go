package postgres

import (
	"database/sql"
	"test/models"
	"github.com/google/uuid"
)

type DriverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) DriverRepo {
	return DriverRepo{
		DB: db,
	}
}

func (d DriverRepo) Insert(driver models.Driver) (string, error) {
	id := uuid.New()

	if _, err := d.DB.Exec(`insert into driver values ($1, $2, $3,$4)`, id, driver.FullName, driver.Phone, driver.CarID); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (d DriverRepo) GetById(id uuid.UUID) (models.Driver, error) {
	driver := models.Driver{}
	d.DB.QueryRow(`select *from driver`).Scan(
		&driver.ID,
		&driver.FullName,
		&driver.Phone,
		&driver.CarID,
	)
	return driver, nil
}

func (d DriverRepo) GetList() ([]models.Driver, error) {
	drivers := []models.Driver{}

	rows, err := d.DB.Query(`select *from driver`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := models.Driver{}

		err := rows.Scan(&driver.ID, &driver.FullName, &driver.Phone, &driver.CarID)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	return drivers, nil
}

func (d DriverRepo) Update(driver models.Driver) error {
	_, err := d.DB.Exec(`update driver set full_name = $1, phone = $2, car_id = $3 where id = $4`,
		driver.FullName, driver.Phone, driver.CarID, driver.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d DriverRepo) Delete(id uuid.UUID) error {
	_, err := d.DB.Exec(`delete from driver where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
