package service

import (
	"damingerdai/address/dao"
	"damingerdai/address/models"
	"damingerdai/address/utils"
	"fmt"
)

func CreateUser(user *models.User) (id int64, err error) {
	id, err = utils.CreateID()
	if err != nil {
		return
	}
	user.ID = id
	err = dao.CreateUser(user.ID, user.Username, user.Password)
	fmt.Println(err)
	return
}
