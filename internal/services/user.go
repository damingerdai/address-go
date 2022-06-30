package service

import (
	"context"
	"database/sql"
	"github.com/damingerdai/address-go/internal/dao"
	"github.com/damingerdai/address-go/internal/models"
	"github.com/damingerdai/address-go/internal/utils"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func CreateUser(ctx context.Context, username, password string) (int64, error) {
	id, err := utils.CreateID()
	if err != nil {
		return 0, errors.Wrapf(err, "fail to create user id: %s", err.Error())
	}
	trx := ctx.Value("trx").(*sqlx.Tx)
	err = dao.CreateUser(trx, id, username, password)
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func GetUser(ctx context.Context, id int64) (*models.User, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	user, err := dao.GetUserById(trx, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrapf(err, "fail to get user id(%d): %s", id, err.Error())
	} else {
		return user, nil
	}
}

func ListUsers(ctx context.Context) (*[]models.User, error) {
	trx := ctx.Value("trx").(*sqlx.Tx)
	users, err := dao.ListUsers(trx)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "fail to list user")
	} else {
		return users, nil
	}
}
