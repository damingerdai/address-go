package service

import (
	"damingerdai/address/dao"
	"damingerdai/address/models"
	"strconv"
)

func ListCities() []*models.City {
	return dao.ListCities()
}

func GetCity(id string) (*models.City, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return dao.GetCity(n)
}
