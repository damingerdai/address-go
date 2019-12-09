package dao

import (
	"damingerdai/address/models"
)

func ListCities() []*models.City {
	rows, err := conn.Query("SELECT _id, name, city_id FROM city")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	result := make([]*models.City, 0, 343)
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

	return result
}
