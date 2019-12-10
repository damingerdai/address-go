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

func GetCity(id int) (*models.City, error) {
	rows := conn.QueryRow("SELECT _id as Id, name as Name, city_id as CityId FROM city WHERE _id = ?", id)
	var cityId int
	var name string
	err := rows.Scan(&id, &name, &cityId)
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
