package service

import (
	"context"
	"damingerdai/address/internal/dao"
	"damingerdai/address/internal/models"
	"github.com/jmoiron/sqlx"
	"strconv"

	"github.com/pkg/errors"
)

func ListProvince(ctx context.Context) ([]*models.Province, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	provinces, err := dao.ListProvinces(trx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get provinces")
	}
	return provinces, nil
}

func GetProvince(ctx context.Context, id string) (*models.Province, error) {
	n, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	trx := ctx.Value("trx").(*sqlx.Tx)
	province, err := dao.GetProvince(trx, n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get the province which id is %s", id)
	}
	return province, nil
}
