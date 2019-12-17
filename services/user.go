package service

import (
	"damingerdai/address/dao"
	"damingerdai/address/models"
	"damingerdai/address/utils"
	"strconv"

	"errors"
)

func CreateUser(user *models.User) (id int64, err error) {
	id, err = utils.CreateID()
	if err != nil {
		return
	}
	user.ID = id
	err = dao.CreateUser(user.ID, user.Username, user.Password)
	return
}

func GetUser(id string) (*models.User, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.New("bad request for user id")
	}
	return dao.GetUserById(n)
}

func HasUser(user *models.User) (bool, error) {
	return dao.HasUser(user)
}
