package service

import (
	"damingerdai/address/internal/dao"
	"damingerdai/address/pkg/utils"
)

func CreateUser(username, password string) (int64, error) {
	id, _ := utils.CreateID()
	err := dao.CreateUser(id, username, password)
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}
