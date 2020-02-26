package service

import (
	"damingerdai/address/internal/dao"
	"damingerdai/address/internal/models"
	"damingerdai/address/internal/utils"
	"database/sql"

	"github.com/pkg/errors"
)

func CreateUser(username, password string) (int64, error) {
	id, err := utils.CreateID()
	if err != nil {
		return 0, errors.Wrapf(err, "fail to create user id: %s", err.Error())
	}
	err = dao.CreateUser(id, username, password)
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func GetUser(id int64) (*models.User, error) {
	user, err := dao.GetUserById(id)
	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrapf(err, "fail to get user id(%s): %s", id, err.Error())
	} else {
		return user, nil
	}
}
