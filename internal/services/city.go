package service

import (
	"damingerdai/address/internal/dao"
	"damingerdai/address/internal/models"
	"github.com/pkg/errors"
	"strconv"
)

func ListCities() ([]*models.City, error) {
	cities, err := dao.ListCities()
	if err != nil {
		return nil, errors.Wrap(err, "fail to get cities")
	}
	return cities, nil
}

func GetCity(id string) (*models.City, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	city, err := dao.GetCity(n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get city which id is %s", id)
	}
	return city, nil
}
