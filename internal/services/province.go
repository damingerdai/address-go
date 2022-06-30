package service

import (
	"context"
	"github.com/damingerdai/address-go/internal/dao"
	"github.com/damingerdai/address-go/internal/models"
	"github.com/jmoiron/sqlx"
	"strconv"

	"github.com/pkg/errors"
)

func ListProvince(ctx context.Context) (*[]models.Province, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	provinceDao := dao.ProvinceDao{Trx: trx}
	provinces, err := provinceDao.ListProvinces()
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
	provinceDao := dao.ProvinceDao{Trx: trx}
	province, err := provinceDao.GetProvince(n)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get the province which id is %s", id)
	}
	return province, nil
}
