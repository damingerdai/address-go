package service

import (
	"context"
	"damingerdai/address/internal/dao"
	"damingerdai/address/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strconv"
)

func ListCities(ctx context.Context) ([]*models.City, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	cities, err := dao.ListCities(trx)
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
	city, err := dao.GetCity(trx, n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get city which id is %s", id)
	}
	return city, nil
}
