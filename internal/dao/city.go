package dao

import (
	"damingerdai/address/internal/models"
	"damingerdai/address/internal/utils"
	jdbc "database/sql"

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

func (cityDao *CityDao) GetCity(id int64) (*models.City, error) {
	var result models.City
	schema := "SELECT _id, name, city_id FROM city WHERE _id = ?"
	err := cityDao.Trx.Get(&result, schema, id)
	if err != nil {
		return &result, utils.If(err != jdbc.ErrNoRows, err, nil).(error)
	}

	return &result, nil
}

func ListCities(trx *sqlx.Tx) ([]*models.City, error) {
	result := make([]*models.City, 0, 343)
	schema := "SELECT _id, name, city_id FROM city"
	rows, err := trx.Query(schema)
	if err != nil {
		return nil, utils.If(err != jdbc.ErrNoRows, err, nil).(error)
	}
	defer rows.Close()
	for rows.Next() {
		var id, cityId int
		var name string
		err = rows.Scan(&id, &name, &cityId)
		if err != nil {
			panic(err.Error())
		}
		city := models.City{
			Id:     id,
			Name:   name,
			CityId: cityId,
		}
		result = append(result, &city)
	}
	return result, nil
}

func GetCity(trx *sqlx.Tx, id int) (*models.City, error) {
	schema := "SELECT _id as Id, name as Name, city_id as CityId FROM city WHERE _id = ?"
	stmt, err := trx.Prepare(schema)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cityId int
	var name string
	b := rows.Next()
	if b == false {
		return nil, jdbc.ErrNoRows
	}
	err = rows.Scan(&id, &name, &cityId)
	if err != nil {
		panic(err.Error())
	}
	city := models.City{
		Id:     id,
		Name:   name,
		CityId: cityId,
	}
	return &city, nil
}
