package dao

import (
	jdbc "database/sql"
	"github.com/damingerdai/address-go/internal/models"
	"github.com/damingerdai/address-go/internal/utils"

	"github.com/jmoiron/sqlx"
)

type CityDao struct {
	Trx *sqlx.Tx
}

func (cityDao *CityDao) ListCities() (*[]models.City, error) {
	result := make([]models.City, 0, 343)
	schema := "SELECT _id, name, city_id FROM city"
	err := cityDao.Trx.Select(&result, schema)
	if err != nil {
		return &result, utils.If(err != jdbc.ErrNoRows, err, nil).(error)
	}

	return &result, nil
}

func (cityDao *CityDao) GetCity(id int) (*models.City, error) {
	var result models.City
	schema := "SELECT _id, name, city_id FROM city WHERE _id = ?"
	err := cityDao.Trx.Get(&result, schema, id)
	if err != nil {
		if err == jdbc.ErrNoRows {
			return &result, nil
		}
		return nil, err
	}

	return &result, nil
}
