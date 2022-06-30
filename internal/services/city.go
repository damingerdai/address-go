package service

import (
	"context"
	"github.com/damingerdai/address-go/internal/dao"
	"github.com/damingerdai/address-go/internal/models"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func ListCities(ctx context.Context) (*[]models.City, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	cityDao := dao.CityDao{Trx: trx}
	cities, err := cityDao.ListCities()
	if err != nil {
		return nil, errors.Wrap(err, "fail to get cities")
	}
	return cities, nil
}

func GetCity(ctx context.Context, id string) (*models.City, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	trx := ctx.Value("trx").(*sqlx.Tx)
	cityDao := dao.CityDao{Trx: trx}
	city, err := cityDao.GetCity(n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get city which id is %s", id)
	}
	return city, nil
}
