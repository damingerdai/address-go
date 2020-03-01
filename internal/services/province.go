package service

import (
	"damingerdai/address/internal/dao"
	"damingerdai/address/internal/models"
	"strconv"

	"github.com/pkg/errors"
)

func ListProvince() ([]*models.Province, error) {
	provinces, err := dao.ListProvinces()
	if err != nil {
		return nil, errors.Wrap(err, "fail to get provinces")
	}
	return provinces, nil
}

func GetProvince(id string) (*models.Province, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	province, err := dao.GetProvince(n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get the province which id is %s", id)
	}
	return province, nil
}
