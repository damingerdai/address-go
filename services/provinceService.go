package service

import "damingerdai/address/models"

import "damingerdai/address/dao"

func ListProvince() []*models.Province {
	return dao.ListProvinces()
}
