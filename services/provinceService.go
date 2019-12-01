package service

import "damingerdai/address/models"

import "damingerdai/address/dao"

import "strconv"

func ListProvinces() []*models.Province {
	return dao.ListProvinces()
}

func GetProvince(id string) (*models.Province, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return dao.GetProvince(n)
}
