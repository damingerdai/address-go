package service

import "damingerdai/address/models"

import "damingerdai/address/dao"

func ListCities() []*models.City {
	return dao.ListCities()
}
