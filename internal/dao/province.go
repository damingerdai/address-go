package dao

import (
	"damingerdai/address/internal/models"
	"github.com/jmoiron/sqlx"
)

func ListProvinces(trx *sqlx.Tx) ([]*models.Province, error) {
	sql := "SELECT _id, name, province_id FROM province"
	stmt, err := trx.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Province, 0, 32)
	for rows.Next() {
		var id, provinceID int
		var name string
		err = rows.Scan(&id, &name, &provinceID)
		if err != nil {
			return nil, err
		}
		province := models.Province{
			Id:         id,
			Name:       name,
			ProvinceId: provinceID,
		}

		result = append(result, &province)
	}

	return result, nil
}

func GetProvince(trx *sqlx.Tx, id int) (*models.Province, error) {
	sql := "SELECT _id, name, province_id FROM province where _id = ?"
	stmt, err := trx.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if ok := rows.Next(); ok {
		var provinceID int
		var name string
		err := rows.Scan(&id, &name, &provinceID)
		if err != nil {
			return nil, err
		}
		province := models.Province{
			Id:         id,
			Name:       name,
			ProvinceId: provinceID,
		}
		return &province, nil
	}

	return nil, nil
}
