package dao

import (
	"damingerdai/address/models"

	"github.com/pkg/errors"
)

func ListProvinces() []*models.Province {
	rows, err := conn.Query("SELECT _id, name, province_id FROM province")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	result := make([]*models.Province, 0, 32)
	for rows.Next() {
		//将行数据保存到record字典
		var id, provinceId int
		var name string
		err = rows.Scan(&id, &name, &provinceId)
		if err != nil {
			panic(err.Error())
		}
		province := models.Province{
			Id:         id,
			Name:       name,
			ProvinceId: provinceId,
		}

		result = append(result, &province)

	}

	return result

}

func GetProvince(id int) (*models.Province, error) {
	rows := conn.QueryRow("SELECT _id, name, province_id FROM province where _id = ?", id)

	var provinceID int
	var name string
	err := rows.Scan(&id, &name, &provinceID)
	if err != nil {
		return nil, errors.Wrap(err, "db error")
	}
	province := models.Province{
		Id:         id,
		Name:       name,
		ProvinceId: provinceID,
	}
	return &province, nil

}
