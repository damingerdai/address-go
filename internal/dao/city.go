package dao

import "damingerdai/address/internal/models"

import "errors"

func ListCities() []*models.City {
	rows, err := GetConnection().Query("SELECT _id, name, city_id FROM city")
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
	sql := "SELECT _id as Id, name as Name, city_id as CityId FROM city WHERE _id = ?"
	stmt, err := GetConnection().Prepare(sql)
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
		return nil, errors.New("no city")
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
