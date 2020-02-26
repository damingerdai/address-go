package service

import (
	"damingerdai/address/internal/dao"
	"damingerdai/address/pkg/utils"

	"github.com/pkg/errors"
)

func CreateUser(username, password string) (int64, error) {
	id, err := utils.CreateID()
	if err != nil {
		return 0, errors.Wrapf("fail to create user id: %s", err.Error())
	}
	err := dao.CreateUser(id, username, password)
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}
